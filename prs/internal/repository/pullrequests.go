package repository

import (
	"context"
	"prs/internal/model"
)

func (r *Repository) PullRequestCreate(ctx context.Context, pr *model.PullRequest) error {
	return r.DB.WithContext(ctx).Create(pr).Error
}

func (r *Repository) PullRequestExist(ctx context.Context, pull_request_id string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.PullRequest{}).Where("pull_request_id = ?", pull_request_id).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *Repository) GetPullRequest(ctx context.Context, pull_request_id string) (*model.PullRequest, error) {
	var res model.PullRequest
	err := r.DB.Model(&model.PullRequest{}).Where("pull_request_id = ?", pull_request_id).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) AddReviewer(ctx context.Context, reviewer *model.PullRequestReviewer) error {
	return r.DB.WithContext(ctx).Create(reviewer).Error
}

func (r *Repository) GetPullRequestReviewers(ctx context.Context, pull_request_id string) ([]model.PullRequestReviewer, error) {
	var res []model.PullRequestReviewer
	err := r.DB.Model(&model.PullRequestReviewer{}).Where("pull_request_id = ?", pull_request_id).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) MergePullRequest(ctx context.Context, pull_request_id string) error {
	return r.DB.WithContext(ctx).
        Model(&model.PullRequest{}).
        Where("pull_request_id = ?", pull_request_id).
        Update("status", model.PullRequestStatusMerged).Error
}