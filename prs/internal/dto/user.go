package dto

type User struct {
	UserId		string 		`json:"id"`
	UserName	string		`json:"username"`
	TeamName	string		`json:"team_name"`
	IsActive	bool		`json:"is_active"`
}