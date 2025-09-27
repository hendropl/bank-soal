package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/model"
)

type UserRepository interface {
	Register(ctx context.Context, user model.User) (*model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, user model.User) error
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByNim(ctx context.Context, nim string) (*model.User, error)
	GetUsersByName(ctx context.Context, name string) ([]model.User, error)
	GetUsersByRole(ctx context.Context, role string) ([]model.User, error)
	ChangePassword(ctx context.Context, id int, password string) error
	ChangeRole(ctx context.Context, id int, role string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	user := model.User{}

	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user model.User) (*model.User, error) {
	existUser := model.User{}
	if err := r.db.WithContext(ctx).First(&existUser, user.Id).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Updates(user).Error; err != nil {
		return nil, err
	}

	return &existUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user model.User) error {
	if err := r.db.WithContext(ctx).Delete(&user, user.Id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByNim(ctx context.Context, nim string) (*model.User, error) {
	user := model.User{}
	if err := r.db.WithContext(ctx).Where("nim = ?", nim).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUsersByName(ctx context.Context, name string) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUsersByRole(ctx context.Context, role string) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) ChangePassword(ctx context.Context, id int, password string) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) ChangeRole(ctx context.Context, id int, role string) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("role", role).Error; err != nil {
		return err
	}
	return nil
}
