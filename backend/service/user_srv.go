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
	Login(ctx context.Context, cred model.LoginCredential) (*model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, user model.User) error
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByNim(ctx context.Context, nim string) (*model.User, error)
	GetUsersByName(ctx context.Context, name string) ([]model.User, error)
	GetUsersByRole(ctx context.Context, role string) ([]model.User, error)
	ChangePassword(ctx context.Context, id int, password string) error
	ChangeRole(ctx context.Context, id int, role string) error
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
		return fmt.Errorf("failed to hash password: %w", err)
	}

	validate := validator.New()

	if err := validate.Struct(credential); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existingUser, _ := s.repo.GetUserByEmail(ctx, credential.Email)
	if existingUser != nil {
		return fmt.Errorf("email %s already used", credential.Email)
	}

	// 	if year < 1000 || year > 9999 { //soon
	//     return fmt.Errorf("Academic year must be a 4-digit number")
	// }

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

func (s *userService) Login(ctx context.Context, cred model.LoginCredential) (*model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, cred.Email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found", cred.Email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	return user, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (*model.User, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found: %w", email, err)
	}
	return user, nil
}

func (s *userService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	updatedUser, err := s.repo.Update(ctx, *user)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(ctx context.Context, user model.User) error {
	err := s.repo.DeleteUser(ctx, user)
	if err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}
	return nil
}

func (s *userService) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("users are not found: %w", err)
	}
	return users, nil
}

func (s *userService) GetUserByNim(ctx context.Context, nim string) (*model.User, error) {
	user, err := s.repo.GetUserByNim(ctx, nim)
	if err != nil {
		return nil, fmt.Errorf("user with nim %q are not found: %w", nim, err)
	}
	return user, nil
}

func (s *userService) GetUsersByName(ctx context.Context, name string) ([]model.User, error) {
	users, err := s.repo.GetUsersByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("user with name %q are not found: %w", name, err)
	}
	return users, nil
}

func (s *userService) GetUsersByRole(ctx context.Context, role string) ([]model.User, error) {
	users, err := s.repo.GetUsersByRole(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("user with role %q are not found: %w", role, err)
	}
	return users, nil
}

func (s *userService) ChangePassword(ctx context.Context, id int, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = s.repo.ChangePassword(ctx, id, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("cannot change password: %w", err)
	}
	return nil
}

func (s *userService) ChangeRole(ctx context.Context, id int, role string) error {
	err := s.repo.ChangeRole(ctx, id, role)
	if err != nil {
		return fmt.Errorf("cannot change role: %w", err)
	}
	return nil
}
