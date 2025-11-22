package service

import (
	"context"
	"prs/internal/dto"
	"prs/internal/repository"
)

type PRService interface {
	AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, *dto.ErrorResponse)
 
	// TODO: GetTeam
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

func (prs *prservice) AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, *dto.ErrorResponse) {
	// TODO
	return team, nil
}


