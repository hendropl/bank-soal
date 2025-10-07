package service

import (
	"context"
	"fmt"

	"latih.in-be/model"
	"latih.in-be/repository"
)

type QuestionService interface {
	Create(ctx context.Context, data model.Question) (*model.Question, error)
	GetById(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, data model.Question) (*model.Question, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Question, error)
}

type questionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionService{
		repo: repo,
	}
}

func (s *questionService) Create(ctx context.Context, data model.Question) (*model.Question, error) {
	if data.ExamId == 0 {
		return nil, fmt.Errorf("examId is required")
	}

	createdData, err := s.repo.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create question: %w", err)
	}

	return createdData, nil
}

func (s *questionService) GetById(ctx context.Context, id int) (*model.Question, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *questionService) Update(ctx context.Context, data model.Question) (*model.Question, error) {
	updatedData, err := s.repo.Update(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedData, nil
}

func (s *questionService) GetAll(ctx context.Context) ([]model.Question, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *questionService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
