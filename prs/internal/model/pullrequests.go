package model

type PullRequest struct {
	ID			uint	`gorm:"column:id;primaryKey"`
	Name		string  `gorm:"column:name;not null;type:text"`
	AuthorID	uint  	`gorm:"column:author_id;index"`
	IsActive 	bool 	`gorm:"column:is_active;not null"`
	Status 		string 	`gorm:"type:text;index"` // OPEN|MERGED
}

func (PullRequest) TableName() string {
	return "pull_requests"
}