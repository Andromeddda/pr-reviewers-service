package model

type PullRequestReviewer struct {
    PullRequestID 	string 		`gorm:"column:pull_request_id;primaryKey"`
    UserID        	string 		`gorm:"column:user_id;primaryKey"`
}

func (PullRequestReviewer) TableName() string {
	return "pull_request_reviewers"
}