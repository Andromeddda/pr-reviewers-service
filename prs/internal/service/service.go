package service

import (
	"context"
	"log"
	"prs/internal/dto"
	"prs/internal/repository"
)

type PRService interface {
	AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, error)
 
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

func (prs *prservice) AddTeam(ctx context.Context, team *dto.Team) (*dto.Team, error) {
	err := prs.repo.AddTeam(ctx, team)

	if err != nil {
		return nil, err
	}

	log.Printf("Created new Team \"%s\" with %d members: ", team.TeamName, len(team.Members))
	for _, m := range(team.Members) {
		log.Printf("Team member: {\"user_id\": \"%s\", \"username\": \"%s\", \"is_active\": \"%v\"}", m.UserId, m.UserName, m.IsActive)
	}

	return team, nil
}


