package service

import (
	"context"
	"log"
	"math/rand/v2"
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
	CreatePullRequest(ctx context.Context, pull_request_id, pull_request_name, author_id string) (*dto.PullRequest, error)


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

func (prs *prservice) CreatePullRequest(ctx context.Context, pull_request_id, pull_request_name, author_id string) (*dto.PullRequest, error) {
	var res *dto.PullRequest
	
	err := prs.repo.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := prs.repo.WithTx(tx)

		// Check if already exist
		pr_exist, err := repo.PullRequestExist(ctx, pull_request_id)

		if err != nil {
			return err // Internal error
		}

		if pr_exist {
			return ErrPRExist // Pull Request already exist
		}

		// Create PR
		pr := model.NewPullRequest(pull_request_id, pull_request_name, author_id)
		err = repo.PullRequestCreate(ctx, pr)

		if err != nil {
			return err // Internal error
		}

		// Get author

		author_exist, err := repo.UserExist(ctx, author_id)

		if err != nil {
			return err // Internal error
		}

		if !author_exist {
			return ErrAuthorNotFound // Author not found
		}

		author, err := repo.GetUser(ctx, author_id)

		if err != nil {
			return err // Internal error
		}

		// Assign reviewers

		members, err := repo.GetTeamMembers(ctx, author.TeamName)

		if err != nil {
			return err // Internal error
		}

		reviewers := pickRandomReviewers(author_id, members)

		for _, r := range(reviewers) {
			err = repo.AddReviewer(ctx, &model.PullRequestReviewer{
				PullRequestID: pull_request_id,
				UserID: r,
			})

			if err != nil {
				return err // Internal error
			}
		}

		assigned_reviewers, err := repo.GetPullRequestReviewers(ctx, pull_request_id)

		if err != nil {
			return err // Internal error
		}

		res, err = mapper.PullRequestToDTO(pr, assigned_reviewers)

		if err != nil {
			return err // Internal error
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Printf("[CREATE PULL REQUEST] {\"%s\", \"%s\"} by  \"%s\" with %d reviewers: ", 
		res.PullRequestId, 
		res.PullRequestName,
		res.AuthorId, 
		len(res.AssignedReviewers))

	for _, m := range(res.AssignedReviewers) {
		log.Printf("Reviewer: \"user_id\": \"%s\"", m)
	}

	return res, nil
}

func pickRandomReviewers(author_id string, members []model.User) []string {
	// Exclude author and inactive users
	candidates := make([]model.User, 0, len(members))
	for _, m := range members {
		if m.UserID != author_id && m.IsActive {
			candidates = append(candidates, m)
		}
	}

	// No candidates
	if len(candidates) == 0 {
		return []string{}
	}

	// Shuffle
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	// Choose candidates
	n := min(2, len(candidates))
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = candidates[i].UserID
	}

	return result
}

