package model

type PullRequestStatus string

const (
	PullRequestStatusOpen 	PullRequestStatus = "OPEN"
	PullRequestStatusMerged PullRequestStatus = "MERGED"
)