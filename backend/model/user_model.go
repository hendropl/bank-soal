package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"name" validate:"required"`
	Nim       string    `json:"nim,omitempty" gorm:"unique" validate:"required"`
	ImgUrl    string    `json:"img_url,omitempty"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	Role      Role      `json:"role" validate:"oneof=admin user super_admin lecturer"`
	Major     string    `json:"major,omitempty"`
	Faculty   string    `json:"faculty,omitempty"`
	Status    Status    `json:"status" validate:"oneof=passed not_passed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
