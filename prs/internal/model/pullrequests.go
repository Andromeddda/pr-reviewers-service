package model

import "time"

type PullRequest struct {
	PullRequestID			string				`gorm:"column:pull_request_id;primaryKey"`
	PullRequestNAme			string  			`gorm:"column:pull_request_name;not null;type:text"`
	AuthorID				string  			`gorm:"column:author_id;index"`
	IsActive 				bool 				`gorm:"column:is_active;not null"`
	Status 					PullRequestStatus 
	CreatedAt				time.Time
	MergedAt				time.Time
	
}

func (PullRequest) TableName() string {
	return "pull_requests"
}