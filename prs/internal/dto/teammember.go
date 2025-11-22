package dto

type TeamMember struct {
	UserId		string 		`json:"user_id"`
	UserName	string		`json:"username"`
	IsActive	bool		`json:"is_active"`
}