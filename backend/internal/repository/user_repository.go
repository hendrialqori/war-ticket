package repository

import (
	"context"
	"errors"

	"github.com/hendrialqori/war-ticket/backend/internal/entity"
	"github.com/hendrialqori/war-ticket/backend/internal/entity/mapper"
	"github.com/hendrialqori/war-ticket/backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) error
	FindByUsernameOrEmail(ctx context.Context, value string) (*entity.User, error)
	IsActive(ctx context.Context, email string) (bool, error)
	SetActive(ctx context.Context, user entity.UserActivation) error
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

// FindById implements [UserRepository].
func (u *userRepositoryImpl) FindById(ctx context.Context, id string) (*entity.User, error) {
	var user model.UserModel

	err := u.DB.WithContext(ctx).Preload("Activation").Where("id = ?", id).First(&user).Error

	return mapper.ToUserEntity(&user), err
}

func (u *userRepositoryImpl) SetActive(ctx context.Context, user entity.UserActivation) error {
	q := gorm.G[model.UserActivationModel](u.DB)
	// make query by email
	existing, err := q.Where("email = ?", user.Email).First(ctx)
	// handle it if not found record
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return q.Create(ctx, &model.UserActivationModel{
			ID:       user.ID,
			Email:    user.Email,
			IsActive: 1,
			UserID:   user.UserID,
		})
	}
	// handle it if there's another error
	if err != nil {
		return err
	}
	// if not error, update row
	_, err = q.Where("id = ?", existing.ID).
		Update(ctx, "is_active", 1)

	return err
}

func (u *userRepositoryImpl) IsActive(ctx context.Context, email string) (bool, error) {
	_, err := gorm.G[model.UserActivationModel](u.DB).Where("email = ? AND is_active = 1", email).First(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepositoryImpl) FindByUsernameOrEmail(ctx context.Context, value string) (*entity.User, error) {
	user, err := gorm.G[model.UserModel](u.DB).Where("email = ? OR username = ?", value, value).First(ctx)
	if err != nil {
		return nil, err
	}

	result := entity.User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		HashPassword: user.HashPassword,
	}

	return &result, nil
}

// create new user
func (u *userRepositoryImpl) Create(ctx context.Context, user entity.User) error {
	payload := model.UserModel{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		HashPassword: user.HashPassword,
	}

	result := gorm.WithResult()
	err := gorm.G[model.UserModel](u.DB, result).Create(ctx, &payload)

	return err
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB,
	}
}
