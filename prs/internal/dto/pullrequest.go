package dto

import (
	"prs/internal/model"
	"time"
)

type PullRequest struct {
	PullRequestId		string					`json:"pull_request_id"`
	PullRequestName		string					`json:"pull_request_name"`
	AuthorId			string					`json:"author_id"`	
	Status				model.PullRequestStatus `json:"status"`
	AssignedReviewers	[]string				`json:"assigned_reviewers"`
	CreatedAt			*time.Time				`json:"created_at,omitempty"`	
	MergedAt			*time.Time				`json:"merged_at,omitempty"`

}