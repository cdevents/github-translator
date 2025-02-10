package github

import "encoding/json"

func (pEvent *GithubEvent) HandleRepoCreatedEvent() (string, error) {
	var RepoCreated PushRepoCreated
	err := json.Unmarshal([]byte(pEvent.Event), &RepoCreated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GithubEvent into RepoCreated struct", err)
		return "", err
	}
	Log().Info("RepoCreated Gith received : ", RepoCreated.Repository, RepoCreated.HeadName, RepoCreated.CommonFields.Type)
	RepoCreated.Url = pEvent.repoUrl
	cdEvent, err := RepoCreated.RepositoryCreatedToCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated repo-created github event into dev.cdevents.repository.created CDEvent: ", cdEvent)
	return cdEvent, nil
}

func (pEvent *GithubEvent) HandleBranchCreatedEvent() (string, error) {
	var changeUpdated PushChangeUpdated
	err := json.Unmarshal([]byte(pEvent.Event), &changeUpdated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GithubEvent into Branch-created struct", err)
		return "", err
	}
	Log().Info("BranchCreated Git received : ", changeUpdated.Repository, changeUpdated.OldHead, changeUpdated.CommonFields.Type)
	changeUpdated.Url = pEvent.repoUrl
	cdEvent, err := changeUpdated.RepoUpdatedToCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated branch-created github event into dev.cdevents.repository.created CDEvent: ", cdEvent)
	return cdEvent, nil
}
