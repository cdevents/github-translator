package github

import (
	sdk "github.com/cdevents/sdk-go/pkg/api"
)

var SpecVersion = "0.4.1"

func (repoCreated *PushRepoCreated) RepositoryCreatedToCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepositoryCreatedEvent")
	cdEvent, err := sdk.NewRepositoryCreatedEventV0_2_0(SpecVersion)
	if err != nil {
		Log().Error("Error creating CDEvent RepositoryCreatedEvent %s\n", err)
	}

	cdEvent.SetSource(repoCreated.Url)
	cdEvent.SetSubjectName(repoCreated.Repository)
	cdEvent.SetSubjectId(repoCreated.HeadName)
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepositoryCreated CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (changeUpdated *PushChangeUpdated) RepoUpdatedToCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepoUpdatedToCDEvent")
	cdEvent, err := sdk.NewBranchCreatedEventV0_2_0(SpecVersion)
	if err != nil {
		Log().Error("Error creating CDEvent RepoUpdatedToCDEvent %s\n", err)
	}

	cdEvent.SetSource(changeUpdated.CommonFields.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: changeUpdated.NewHead})
	cdEvent.SetSubjectId(changeUpdated.NewHead)
	cdEvent.SetSubjectSource(changeUpdated.Repository)
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepoUpdatedToCDEvent CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}
