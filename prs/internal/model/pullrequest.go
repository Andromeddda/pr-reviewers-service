package model

import "time"

type PullRequest struct {
	PullRequestID			string				`gorm:"column:pull_request_id;primaryKey"`
	PullRequestName			string  			`gorm:"column:pull_request_name;not null;type:text"`
	AuthorID				string  			`gorm:"column:author_id;index"`
	IsActive 				bool 				`gorm:"column:is_active;not null"`
	Status 					PullRequestStatus 	`gorm:"column:status;not null;type:text"`
	CreatedAt				time.Time			`gorm:"column:created_at;autoCreateTime"`
	MergedAt				*time.Time			`gorm:"column:merged_at"`
	
}

func (PullRequest) TableName() string {
	return "pull_requests"
}