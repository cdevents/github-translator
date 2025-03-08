package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubEvent struct {
	Event   string
	repoUrl string
}

func NewGithubEvent(event string, repoURL string) (pEvent *GithubEvent) {
	pEvent = &GithubEvent{event, repoURL}
	return
}

func HandleTranslateGithubEvent(event string, header http.Header) (string, error) {
	Log().Debug("Handle translation into CDEvent from Github event %s\n", event)
	repoURL := ""
	if header.Get("X-Origin-Url") != "" {
		repoURL = header.Get("X-Origin-Url")
	}
	if header.Get("X-GitHub-Event") != "" {
		event = header.Get("X-GitHub-Event")
	}
	githubEvent := NewGithubEvent(event, repoURL)
	cdEvent, err := githubEvent.TranslateIntoCDEvent()
	if err != nil {
		Log().Error("Error translating Github event into CDEvent %s\n", err)
		return "", err
	}
	Log().Debug("Github Event translated into CDEvent %s\n", cdEvent)
	return cdEvent, nil
}

func (pEvent *GithubEvent) TranslateIntoCDEvent() (string, error) {
	eventMap := make(map[string]interface{})
	cdEvent := ""
	err := json.Unmarshal([]byte(pEvent.Event), &eventMap)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GithubEvent data into GithubEvent map", err)
		return "", err
	}
	eventType := eventMap["type"]
	oldBranch := eventMap["base_ref"]
	newBranch := eventMap["ref"]
	before := eventMap["before"]
	after := eventMap["after"]
	created := eventMap["created"]
	deleted := eventMap["deleted"]
	modified := eventMap["modified"]

	Log().Info("handling translating to CDEvent from Github Event type: %s ", eventType)

	switch eventType {
	case PushCreated:
		if before == ZeroedSha && created == true {
			if oldBranch != nil && newBranch != nil {
				cdEvent, err = pEvent.HandleBranchCreatedEvent()
				if err != nil {
					return "", err
				}
			}
		}
	case RepositoryCreatedEventType:
		if created != true {
			cdEvent, err = pEvent.HandleRepoCreatedEvent()
			if err != nil {
				return "", err
			}
		}
	case PushDeleted:
		if oldBranch == nil && newBranch != nil {
			if after == ZeroedSha && deleted == true {
				cdEvent, err = pEvent.HandleBranchDeletedEvent()
				if err != nil {
					return "", err
				}
			}
		}
	case PushModified:
		if oldBranch == nil && newBranch != nil {
			if modified == true {
				cdEvent, err = pEvent.HandleBranchModifiedEvent()
				if err != nil {
					return "", err
				}
			}
		}
	default:
		Log().Info("Not handling CDEvent translation for Github event type: %s\n", eventMap["type"])
		return "", fmt.Errorf("Github event type %s, not supported for translation", eventType)
	}
	return cdEvent, nil
}
