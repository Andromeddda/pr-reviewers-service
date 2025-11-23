package service

import "errors"

var (
	ErrTeamExist = errors.New("team already exist")
	ErrTeamNotFound = errors.New("team not found")
	ErrUserNotFound = errors.New("user not found")
	ErrAuthorNotFound = errors.New("author not found")
	ErrPRNotFound = errors.New("pull-request not found")
	ErrPRExist = errors.New("pull-request already exist")
	ErrPRMerged = errors.New("pull-request already merged")
	ErrNotAssigned = errors.New("user wasn't assigned to pull-request")
	ErrNoCandidate = errors.New("no candidate to reassign to pull-request")
)