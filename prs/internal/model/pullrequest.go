package model

import "time"

type PullRequest struct {
	PullRequestID			string				`gorm:"column:pull_request_id;primaryKey"`
	PullRequestName			string  			`gorm:"column:pull_request_name;not null;type:text"`
	AuthorID				string  			`gorm:"column:author_id;index"`
	Status 					PullRequestStatus 	`gorm:"column:status;not null;type:text"`
	CreatedAt				time.Time			`gorm:"column:created_at;autoCreateTime"`
	MergedAt				*time.Time			`gorm:"column:merged_at"`
	
	// reviewers are resolved by query
}

func NewPullRequest(pull_request_id, pull_request_name, author_id string) *PullRequest {
	now := time.Now().UTC()
	return &PullRequest{
		PullRequestID:   pull_request_id,
		PullRequestName: pull_request_name,
		AuthorID:        author_id,
		Status:          PullRequestStatusOpen,
		CreatedAt:       now,
		MergedAt:        nil,
	}
}

func (PullRequest) TableName() string {
	return "pull_requests"
}