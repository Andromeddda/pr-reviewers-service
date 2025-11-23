package repository

import (
	"context"
	"log"
	"prs/internal/model"
)

func (r *Repository) UserExist(ctx context.Context, user_id string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.User{}).Where("user_id = ?", user_id).Count(&count).Error

	if err != nil {
		return false, err
	}

	log.Printf("UserExist: user_id=%s, count=%d", user_id, count)

	return count == 1, nil
}

func (r *Repository) GetUser(ctx context.Context, user_id string) (*model.User, error) {
	var res model.User
	err := r.DB.Model(&model.User{}).Where("user_id = ?", user_id).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) AddUser(ctx context.Context, user *model.User) error {
	log.Printf("AddUser: user_id=%s, username=%s, team_name=%s, is_active=%v", user.UserID, user.UserName, user.TeamName, user.IsActive)
	return  r.DB.WithContext(ctx).Create(user).Error
}

func (r *Repository) UserSetIsActive(ctx context.Context, user_id string, is_active bool) error {
	return r.DB.WithContext(ctx).
        Model(&model.User{}).
        Where("user_id = ?", user_id).
        Update("is_active", is_active).Error
}