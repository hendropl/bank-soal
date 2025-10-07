package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type ExamScoreRepository interface {
	Create(ctx context.Context, e model.ExamScore) (*model.ExamScore, error)
	GetById(ctx context.Context, id int) (*model.ExamScore, error)
	GetAll(ctx context.Context) ([]model.ExamScore, error)
	GetByExam(ctx context.Context, examId int) ([]model.ExamScore, error)
	Update(ctx context.Context, e model.ExamScore) (*model.ExamScore, error)
	Delete(ctx context.Context, id int) error
}

type examScoreRepository struct {
	db *gorm.DB
}

func NewExamScoreRepository(db *gorm.DB) ExamScoreRepository {
	return &examScoreRepository{db: db}
}

func (r *examScoreRepository) Create(ctx context.Context, e model.ExamScore) (*model.ExamScore, error) {
	if err := r.db.WithContext(ctx).Create(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examScoreRepository) GetById(ctx context.Context, id int) (*model.ExamScore, error) {
	e := model.ExamScore{}

	if err := r.db.WithContext(ctx).First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examScoreRepository) GetAll(ctx context.Context) ([]model.ExamScore, error) {
	var e []model.ExamScore
	if err := r.db.WithContext(ctx).Find(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *examScoreRepository) GetByExam(ctx context.Context, examId int) ([]model.ExamScore, error) {
	var e []model.ExamScore

	if err := r.db.WithContext(ctx).Where("exam_id = ?", examId).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *examScoreRepository) Update(ctx context.Context, e model.ExamScore) (*model.ExamScore, error) {
	if err := r.db.WithContext(ctx).Updates(e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *examScoreRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(id).Error; err != nil {
		return err
	}
	return nil
}
