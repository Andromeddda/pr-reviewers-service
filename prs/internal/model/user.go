package model

type Users struct {
	UserID		string	`gorm:"column:user_id;primaryKey"`
	UserName	string  `gorm:"column:username;not null;type:text"`
	TeamName	string  `gorm:"column:team_name;index;type:text"`
	IsActive 	bool 	`gorm:"column:is_active;not null"`
}

func (Users) TableName() string {
	return "users"
}
