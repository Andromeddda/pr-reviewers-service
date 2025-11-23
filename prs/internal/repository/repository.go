package repository

import (
	"context"
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

func (r *Repository) TeamExist(ctx context.Context, team_name string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.Team{}).Where("team_name = ?", team_name).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func (r *Repository) GetTeam(ctx context.Context, team_name string) (*model.Team, error) {
	var res model.Team
	err := r.DB.Model(&model.Team{}).Where("team_name = ?", team_name).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}


func (r *Repository) AddTeam(ctx context.Context, team *model.Team) error {
	return  r.DB.WithContext(ctx).Create(team).Error
}

func (r *Repository) AddUser(ctx context.Context, user *model.User) error {
	return  r.DB.WithContext(ctx).Create(user).Error
}

func (r *Repository) GetTeamMembers(ctx context.Context, team_name string) ([]model.User, error) {
	var userList []model.User
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("team_name = ?", team_name).Find(&userList).Error

	if err != nil {
		return nil, err
	}

	return userList, nil
}
