package github

/**
* Github-webhook-Event types
**/

type CommonFields struct {
	Type      string  `json:"type"`
	CreatedAt float64 `json:"createdAt"`
	Url       string  `json:"repoURL,omitempty"`
	Verb      Verb    `json:"verb,omitempty"`
}

type Verb struct {
	Created  bool `json:"created"`
	Modified bool `json:"modified"`
	Deleted  bool `json:"deleted"`
	Forced   bool `json:"forced"`
}
type Commiter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Change struct {
	Repository    string   `json:"repository"`
	Branch        string   `json:"branch"`
	Id            string   `json:"id"`
	Number        int      `json:"number"`
	Subject       string   `json:"subject"`
	Owner         Commiter `json:"commiter"`
	CommitMessage string   `json:"commitMessage"`
	CreatedOn     float64  `json:"createdOn"`
	Status        string   `json:"status"`
	Ref           string   `json:"ref"`
	BaseRef       string   `json:"base_ref"`
	Before        string   `json:"before"`
	After         string   `json:"after"`
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
	CommonFields
	Change
	Commiter
}
