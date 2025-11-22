package repository

import "errors"

var (
	ErrTeamExist = errors.New("team already exist")
	ErrTeamNotFound = errors.New("team not found")
)