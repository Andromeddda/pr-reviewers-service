package model

type Teams struct {
	TeamName		string  `gorm:"column:team_name;primaryKey;type:text"`

	// members are resolved by query
}

func (Teams) TableName() string {
	return "teams"
}