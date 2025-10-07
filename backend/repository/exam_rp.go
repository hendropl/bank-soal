package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type ExamRepository interface {
	Create(ctx context.Context, exam model.Exam) (*model.Exam, error)
	GetById(ctx context.Context, id int) (*model.Exam, error)
	Update(ctx context.Context, exam model.Exam) (*model.Exam, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Exam, error)
}

type examRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepository{db: db}
}

func (r *examRepository) Create(ctx context.Context, exam model.Exam) (*model.Exam, error) {
	if err := r.db.WithContext(ctx).Create(&exam).Error; err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) GetById(ctx context.Context, id int) (*model.Exam, error) {
	exam := model.Exam{}

	if err := r.db.WithContext(ctx).First(&exam, id).Error; err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) Update(ctx context.Context, exam model.Exam) (*model.Exam, error) {
	if err := r.db.WithContext(ctx).Updates(exam).Error; err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

func (r *examRepository) GetAll(ctx context.Context) ([]model.Exam, error) {
	var exams []model.Exam
	if err := r.db.WithContext(ctx).Find(&exams).Error; err != nil {
		return nil, err
	}
	return exams, nil
}
