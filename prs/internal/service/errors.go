package service

import "errors"

var (
	ErrTeamExist = errors.New("team already exist")
	ErrTeamNotFound = errors.New("team not found")
	ErrUserNotFound = errors.New("user not found")
	ErrAuthorNotFound = errors.New("author not found")
	ErrPRNotFound = errors.New("pull request not found")
	ErrPRExist = errors.New("pull-request already exist")
)