package dto

type PullRequestReassigned struct {
	PR 				*PullRequest 	`json:"pr"`
	ReplacedBy 		string 			`json:"replaced_by"`
}


func NewPRReassigned(pr *PullRequest, replacedBy string) *PullRequestReassigned {
	return &PullRequestReassigned{
		PR: pr,
		ReplacedBy: replacedBy,
	}
}