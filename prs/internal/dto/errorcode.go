package dto

type ErrorCode string

const (
	ErrorTeamExist 		ErrorCode = "TEAM_EXIST"
	ErrorPRExist 		ErrorCode = "PR_EXIST"
	ErrorPRMerged 		ErrorCode = "PR_MERGED"
	ErrorNotAssigned 	ErrorCode = "NOT_ASSIGNED"
	ErrorNoCandidate 	ErrorCode = "NO_CANDIDATE"
	ErrorNotFound 		ErrorCode = "NOT_FOUND"
)