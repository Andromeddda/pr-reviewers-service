package model

type User struct {
	UserID		string	`gorm:"column:user_id;primaryKey"`
	UserName	string  `gorm:"column:username;not null;type:text"`
	TeamName	string  `gorm:"column:team_name;index;type:text"`
	IsActive 	bool 	`gorm:"column:is_active;not null"`
}

func (User) TableName() string {
	return "users"
}
