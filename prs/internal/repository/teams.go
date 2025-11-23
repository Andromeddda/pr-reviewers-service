package repository

import (
	"context"
	"prs/internal/model"
)

func (r *Repository) TeamExist(ctx context.Context, team_name string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.Team{}).Where("team_name = ?", team_name).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
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

func (r *Repository) GetTeamMembers(ctx context.Context, team_name string) ([]model.User, error) {
	var userList []model.User
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("team_name = ?", team_name).Find(&userList).Error

	if err != nil {
		return nil, err
	}

	return userList, nil
}