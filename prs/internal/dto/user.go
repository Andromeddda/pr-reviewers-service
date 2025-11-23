package dto

type User struct {
	UserID		string 		`json:"user_id"`
	UserName	string		`json:"username"`
	TeamName	string		`json:"team_name"`
	IsActive	bool		`json:"is_active"`
}