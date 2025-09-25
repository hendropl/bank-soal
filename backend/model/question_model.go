package model

import "time"

type Question struct {
	Id           int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ExamId       int        `json:"exam_id" gorm:"not null"`
	CreatorId    int        `json:"creator_id" gorm:"not null"`
	QuestionText string     `json:"question_text" validate:"required"`
	Difficulty   Difficulty `json:"difficulty" validate:"oneof=easy medium hard"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Options      []Option   `json:"options,omitempty" gorm:"foreignKey:QuestionID"`
}
