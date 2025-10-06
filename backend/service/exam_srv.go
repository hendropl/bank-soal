package service

import (
	"context"
	"fmt"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type ExamService interface {
	Create(ctx context.Context, data model.Exam) error
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, data model.Exam) (*model.Exam, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Exam, error)
}

type examService struct {
	repo repository.ExamRepository
}

func NewExamService(repo repository.ExamRepository) ExamService {
	return &examService{
		repo: repo,
	}
}

func (s *examService) Create(ctx context.Context, data model.Exam) error {
	if data.CreatorId == 0 {
		return fmt.Errorf("creatorId is required")
	}

	if data.EndedAt.Before(data.StartedAt) {
		return fmt.Errorf("ended_at must be after started_at")
	}

	err := s.repo.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create exam: %w", err)
	}

	return nil
}

func (s *examService) GetById(ctx context.Context, id int) (*model.Exam, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *examService) Update(ctx context.Context, data model.Exam) (*model.Exam, error) {
	updatedData, err := s.repo.Update(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *examService) GetAll(ctx context.Context) ([]model.Exam, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *examService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
