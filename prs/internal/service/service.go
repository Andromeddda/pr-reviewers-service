package service

import (
	"context"
	"log"
	"prs/internal/dto"
	"prs/internal/mapper"
	"prs/internal/model"
	"prs/internal/repository"

	"gorm.io/gorm"
)

type PRService interface {
	AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, error)
	GetTeam(ctx context.Context, team_name string) (*dto.Team, error)
	UserSetIsActive(ctx context.Context, user_id string, is_active bool) (*dto.User, error)


	// TODO: UserSetIsActive
	// TODO: CreatePullRequest
	// TODO: MergePullRequest
	// TODO: ReassignPullRequest
	// TODO: UserGetReview

}

type prservice struct {
	repo 	repository.Repository
}

func NewPRService(repository repository.Repository) PRService {
	return &prservice{
		repo: repository,
	}	
}

func (prs *prservice) AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, error) {
	err := prs.repo.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := prs.repo.WithTx(tx)

		team_exist, err := repo.TeamExist(ctx, team.TeamName)

		if err != nil {
			return err // Internal error
		}

		if team_exist {
			return ErrTeamExist // Team already exist
		}

		// Convert
		t, u := mapper.TeamFromDTO(team)

		// Add team
		err = repo.AddTeam(ctx, t)

		if err != nil {
			return err // Internal error
		}

		// Add users
		for _, m := range(u) {
			err = repo.AddUser(ctx, &m)

			if err != nil {
				return err // Internal error
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Printf("[ADD NEW TEAM] \"%s\" with %d members: ", team.TeamName, len(team.Members))
	for _, m := range(team.Members) {
		log.Printf("Team member: {\"user_id\": \"%s\", \"username\": \"%s\", \"is_active\": \"%v\"}", m.UserId, m.UserName, m.IsActive)
	}

	return team, nil
}

func (prs *prservice) GetTeam(ctx context.Context, team_name string) (*dto.Team, error) {
	var res *dto.Team

	err := prs.repo.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := prs.repo.WithTx(tx)

		// Check if team exist
		team_exist, err := repo.TeamExist(ctx, team_name)

		if err != nil {
			return err // Internal error
		}

		if !team_exist {
			return ErrTeamNotFound // Team not found
		}

		// Get team
		t, err := repo.GetTeam(ctx, team_name)

		if err != nil {
			return err // Internal error
		}

		// Get members
		m, err := repo.GetTeamMembers(ctx, team_name)

		if err != nil {
			return err // Internal error
		}

		// Convert
		team, err := mapper.TeamToDTO(t, m)

		if err != nil {
			return err // Internal error
		}

		res = team

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Printf("[GET TEAM] \"%s\" with %d members: ", res.TeamName, len(res.Members))
	for _, m := range(res.Members) {
		log.Printf("Team member: {\"user_id\": \"%s\", \"username\": \"%s\", \"is_active\": \"%v\"}", m.UserId, m.UserName, m.IsActive)
	}

	return res, nil

}

func (prs *prservice) UserSetIsActive(ctx context.Context, user_id string, is_active bool) (*dto.User, error) {
	var user *model.User

	tr_err := prs.repo.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := prs.repo.WithTx(tx)

		// Check if user exist
		user_exist, err := repo.UserExist(ctx, user_id)

		if err != nil {
			return err // Internal error
		}

		if !user_exist {
			return ErrUserNotFound // User not found
		}

		// Update user
		err = repo.UserSetIsActive(ctx, user_id, is_active)

		if err != nil {
			return err // Internal error
		}

		// Return updated user
		u, err := repo.GetUser(ctx, user_id)

		if err != nil {
			return err
		}

		user = u

		return nil
	})

	if tr_err != nil {
		return nil, tr_err
	}

	res := mapper.UserToDTO(user)

	log.Printf("[USER SET IS ACTIVE] {%s, %s, %s, %v}", res.UserID, res.UserName, res.TeamName, res.IsActive)
	return res, nil

}



