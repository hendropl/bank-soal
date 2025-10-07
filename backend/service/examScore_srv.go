package service

import (
	"context"
	"fmt"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type ExamScoreService interface {
	Create(ctx context.Context, data model.ExamScore) (*model.ExamScore, error)
	GetById(ctx context.Context, id int) (*model.ExamScore, error)
	Update(ctx context.Context, data model.ExamScore) (*model.ExamScore, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.ExamScore, error)
}

type examScoreService struct {
	repo repository.ExamScoreRepository
}

func NewExamScoreService(repo repository.ExamScoreRepository) ExamScoreService {
	return &examScoreService{
		repo: repo,
	}
}

func (s *examScoreService) Create(ctx context.Context, data model.ExamScore) (*model.ExamScore, error) {
	if data.UserId == 0 {
		return nil, fmt.Errorf("userId is required")
	}

	createdData, err := s.repo.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create examScore: %w", err)
	}

	return createdData, nil
}

func (s *examScoreService) GetById(ctx context.Context, id int) (*model.ExamScore, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *examScoreService) Update(ctx context.Context, data model.ExamScore) (*model.ExamScore, error) {
	updatedData, err := s.repo.Update(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *examScoreService) GetAll(ctx context.Context) ([]model.ExamScore, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *examScoreService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
