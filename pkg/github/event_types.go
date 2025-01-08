package github

/**
* Github-webhook-Event types
**/

type CommonFields struct {
	Type      string  `json:"type"`
	CreatedAt float64 `json:"createdAt"`
	Url       string  `json:"repoURL,omitempty"`
}

/*
* push events for creation, updates and deletes
**/

type PushRepoCreated struct {
	Repository string `json:"repository"`
	HeadName   string `json:"headName"`
	CommonFields
}

type PushChangeUpdated struct {
}

type PushBranchDeleted struct {
}
