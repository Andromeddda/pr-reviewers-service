package dto

type UsersGetReview struct {
	UserID 			string 					`json:"user_id"`
	PullRequests 	[]PullRequestShort		`json:"pull_requests"`
}

func NewUsersGetReview(user_id string, pull_requests []PullRequestShort) *UsersGetReview {
	return &UsersGetReview{
		UserID: user_id,
		PullRequests: pull_requests,
	}
}