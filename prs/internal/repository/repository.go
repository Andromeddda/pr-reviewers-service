package repository

import (
	"prs/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Repository struct {
   DB *gorm.DB
}

func (r *Repository) WithTx(tx *gorm.DB) *Repository {
	return &Repository{
		DB: tx,
	}
}

func NewRepository(dsn string) (*Repository, error) {
   DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   if err != nil {
      return nil, err
   }

   err = DB.AutoMigrate(
	&model.User{},
	&model.Team{},
	&model.PullRequest{},
   )

   if err != nil {
	return nil, err
   }

   return &Repository{
      DB: DB,
   }, nil
}
