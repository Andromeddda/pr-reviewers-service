package repository

import (
	"context"
	"prs/internal/dto"
	"prs/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	AddTeam(ctx context.Context, team *dto.Team) error
	GetTeam(ctx context.Context, team_name string) (*dto.Team, error)
}

type repository struct {
   db *gorm.DB
}

func NewRepository(dsn string) (Repository, error) {
   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   if err != nil {
      return nil, err
   }

   err = db.AutoMigrate(
	&model.Users{},
	&model.Teams{},
	&model.PullRequest{},
   )

   if err != nil {
	return nil, err
   }

   return &repository{
      db: db,
   }, nil
}

func (r *repository) AddTeam(ctx context.Context, team *dto.Team) error {
	return r.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {

			// Check if team exist
			var count int64
			err := tx.Model(&model.Teams{}).Where("team_name = ?", team.TeamName).Count(&count).Error

			if err != nil {
				return err // Internal error
			}

			if count != 0 {
				return ErrTeamExist // User error
			}

			// Create new team
			t := model.Teams{
				TeamName: team.TeamName,
			}

			err = tx.WithContext(ctx).Create(&t).Error

			if err != nil {
				return err // Internal error
			}

			// Create users

			for _, m := range team.Members {
				u := model.Users{
					UserID: m.UserId,
					UserName: m.UserName,
					TeamName: t.TeamName,
					IsActive: m.IsActive,
				}

				err = tx.WithContext(ctx).Create(&u).Error

				if err != nil {
					return err
				}
			}

			return nil
		},
	)
}

func (r *repository) GetTeam(ctx context.Context, team_name string) (*dto.Team, error) {
	var team dto.Team

	tr_err := r.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			var teamList []model.Teams
			err := tx.Model(&model.Teams{}).Where("team_name = ?", team_name).Find(&teamList).Error

			if err != nil {
				return err
			}

			if len(teamList) != 1 {
				return ErrTeamNotFound
			}

			var userList []model.Users
			err = tx.Model(&model.Users{}).Where("team_name = ?", team_name).Find(&userList).Error

			if err != nil {
				return err
			}

			team.TeamName = teamList[0].TeamName

			for _, m := range(userList) {
				team.Members = append(team.Members, dto.TeamMember{
					UserId: m.UserID,
					UserName: m.UserName,
					IsActive: m.IsActive,
				})
			}
		
			return nil
		})

	if tr_err != nil {
		return nil, tr_err
	}

	return &team, nil

}