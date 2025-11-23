package mapper

import (
	"errors"
	"prs/internal/dto"
	"prs/internal/model"
)

func TeamFromDTO(team *dto.Team) (*model.Team, []model.User) {
	var u []model.User

	for _, m := range(team.Members) {
		u = append(u, model.User{
			UserID: m.UserId,
			UserName: m.UserName,
			TeamName: team.TeamName,
			IsActive: m.IsActive,
		})
	}

	return &model.Team{TeamName: team.TeamName}, u
}

func TeamToDTO(team *model.Team, users []model.User) (*dto.Team, error) {
	res := dto.Team{
		TeamName: team.TeamName,
	}

	for _, m := range(users) {
		if m.TeamName != team.TeamName {
			return nil, errors.New("user's TeamName does not match with team's TeamName")
		}

		res.Members = append(res.Members, dto.TeamMember{
			UserId: m.UserID,
			UserName: m.UserName,
			IsActive: m.IsActive,
		})
	}

	return &res, nil
}

func UserToDTO(user *model.User) *dto.User {
	return &dto.User{
		UserID: user.UserID,
		UserName: user.UserName,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func PullRequestToDTO(pr *model.PullRequest, reviewers []model.PullRequestReviewer) (*dto.PullRequest, error) {
	var assigned_reviewers = make([]string, len(reviewers))

	for i, r := range(reviewers) {
		assigned_reviewers[i] = r.UserID

		if r.PullRequestID != pr.PullRequestID {
			return nil, errors.New("reviewer's pull_request_id does not match PR")
		}
	}
	
	return &dto.PullRequest{
		PullRequestId: pr.PullRequestID,
		PullRequestName: pr.PullRequestName,
		AuthorId: pr.AuthorID,
		CreatedAt: pr.CreatedAt,
		MergedAt: pr.MergedAt,
		AssignedReviewers: assigned_reviewers,
	}, nil
}
