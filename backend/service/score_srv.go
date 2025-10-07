package service

import (
	"context"
	"fmt"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type ScoreService interface {
	Create(ctx context.Context, data model.Score) (*model.Score, error)
	GetById(ctx context.Context, id int) (*model.Score, error)
	Update(ctx context.Context, data model.Score) (*model.Score, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, qId int) ([]model.Score, error)
}

type scoreService struct {
	repo repository.ScoreRepository
}

func NewScoreService(repo repository.ScoreRepository) ScoreService {
	return &scoreService{
		repo: repo,
	}
}

func (s *scoreService) Create(ctx context.Context, data model.Score) (*model.Score, error) {
	if data.ExamId == 0 {
		return nil, fmt.Errorf("examId is required")
	}

	createdData, err := s.repo.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create score: %w", err)
	}

	return createdData, nil
}

func (s *scoreService) GetById(ctx context.Context, id int) (*model.Score, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *scoreService) Update(ctx context.Context, data model.Score) (*model.Score, error) {
	updatedData, err := s.repo.Update(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *scoreService) GetAll(ctx context.Context, qId int) ([]model.Score, error) {
	data, err := s.repo.GetAll(ctx, qId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *scoreService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
