package github

import (
	"net/http"
	"os"
	"testing"
)

func TestHandleTranslateProjectCreatedEvent(t *testing.T) {
	event, err := os.ReadFile("testdata/repo-created.json")
	if err != nil {
		t.Fatalf("Failed to read project-created.json file: %v", err)
	}
	headers := http.Header{}
	headers.Set("X-Origin-Url", "http://github.cdevent.translate")
	headers.Set("X-GitHub-Event", "repository")

	cdEvent, err := HandleTranslateGithubEvent(string(event), headers)
	if err != nil {
		t.Errorf("Expected RepositoryCreated CDEvent to be successful.")
		return
	}
	Log().Info("Handle project-created github event into dev.cdevents.repository.created is successful ", cdEvent)
}
