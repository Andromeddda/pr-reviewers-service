package model

type Teams struct {
	Name		string  `gorm:"column:name;primaryKey;type:text"`

	// members are resolved by query
}

func (Teams) TableName() string {
	return "teams"
}