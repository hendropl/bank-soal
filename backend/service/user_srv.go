package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"latih.in-be/model"
	"latih.in-be/repository"
)

type UserService interface {
	Register(ctx context.Context, user model.RegisterCredential) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(ctx context.Context, credential model.RegisterCredential) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credential.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var validate = validator.New()

	if err := validate.Struct(credential); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingUser, _ := s.repo.GetUserByEmail(ctx, credential.Email)
	if existingUser != nil {
		return fmt.Errorf("email %s already used", credential.Email)
	}

	user := model.User{
		Name:     credential.Name,
		Email:    credential.Email,
		Password: string(hashedPassword),
		Major:    credential.Major,
		Nim:      credential.Nim,
		Faculty:  credential.Faculty,
	}

	_, err = s.repo.Register(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}
