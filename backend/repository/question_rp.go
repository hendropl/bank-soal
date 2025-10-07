package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type QuestionRepository interface {
	Create(ctx context.Context, q model.Question) (*model.Question, error)
	GetById(ctx context.Context, id int) (*model.Question, error)
	GetAll(ctx context.Context) ([]model.Question, error)
	GetByExam(ctx context.Context, examId int) ([]model.Question, error)
	Update(ctx context.Context, q model.Question) (*model.Question, error)
	Delete(ctx context.Context, id int) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(ctx context.Context, q model.Question) (*model.Question, error) {
	if err := r.db.WithContext(ctx).Create(&q).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) GetById(ctx context.Context, id int) (*model.Question, error) {
	q := model.Question{}

	if err := r.db.WithContext(ctx).First(&q, id).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var q []model.Question
	if err := r.db.WithContext(ctx).Find(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questionRepository) GetByExam(ctx context.Context, examId int) ([]model.Question, error) {
	var q []model.Question

	if err := r.db.WithContext(ctx).Where("exam_id = ?", examId).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questionRepository) Update(ctx context.Context, q model.Question) (*model.Question, error) {
	if err := r.db.WithContext(ctx).Updates(q).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(id).Error; err != nil {
		return err
	}
	return nil
}
