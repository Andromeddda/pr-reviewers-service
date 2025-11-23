package model

type Team struct {
	TeamName		string  `gorm:"column:team_name;primaryKey;type:text"`

	// members are resolved by query
}

func (Team) TableName() string {
	return "teams"
}