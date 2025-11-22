package dto

type TeamMember struct {
	UserId		string 		`json:"id"`
	UserName	string		`json:"username"`
	IsActive	bool		`json:"is_active"`
}