package model

type PullRequestReviewers struct {
    PullRequestID 	string 		`gorm:"column:pull_request_id;primaryKey"`
    UserID        	string 		`gorm:"column:user_id;primaryKey"`
}

func (PullRequestReviewers) TableName() string {
	return "pull_requests"
}