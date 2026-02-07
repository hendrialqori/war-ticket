package usecase

import (
	"context"

	"github.com/hendrialqori/war-ticket/backend/internal/config"
	"github.com/hendrialqori/war-ticket/backend/internal/dto"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
	"github.com/hendrialqori/war-ticket/backend/internal/exception"
	"github.com/hendrialqori/war-ticket/backend/internal/repository"
	"github.com/hendrialqori/war-ticket/backend/internal/util"
)

type UserUsecase interface {
	Login(ctx context.Context, in dto.LoginRequest) (string, error)
	Register(ctx context.Context, in dto.RegisterRequest) error
	SetActive(ctx context.Context, in dto.SetActiveRequest) error
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
	config         *config.AppConfig
}

func (u *userUsecaseImpl) SetActive(ctx context.Context, in dto.SetActiveRequest) error {
	user, err := u.userRepository.FindByUsernameOrEmail(ctx, in.Email)
	if err != nil {
		return exception.New(404, "User not found")
	}

	// generate random string for id
	uuid, err := util.GenerateRandomString(20)
	if err != nil {
		return exception.New(400, "Error occure while generate uuid")
	}

	var userModel = entity.UserActivation{
		ID:       uuid,
		UserID:   user.ID,
		Email:    user.Email,
		IsActive: 1,
	}

	if err := u.userRepository.SetActive(ctx, userModel); err != nil {
		return err
	}

	return nil
}

func (u *userUsecaseImpl) Login(ctx context.Context, in dto.LoginRequest) (string, error) {
	user, err := u.userRepository.FindByUsernameOrEmail(ctx, in.Username)
	if err != nil {
		return "", exception.New(404, "User not found")
	}

	// check user is active or not
	isActive, _ := u.userRepository.IsActive(ctx, user.Email)
	if !isActive {
		return "", exception.New(400, "User not active yet, please do activation")
	}

	// check valid password
	isValid := util.CheckValidPassword(user.HashPassword, in.Password)
	if !isValid {
		return "", exception.New(400, "Wrong password")
	}
	// generate jwt token
	secretKey := []byte(u.config.Secret)
	jwtToken, err := util.CreateToken(secretKey, user)
	if err != nil {
		return "", exception.New(400, err.Error())
	}

	return jwtToken, nil
}

func (u *userUsecaseImpl) Register(ctx context.Context, in dto.RegisterRequest) error {
	_, err := u.userRepository.FindByUsernameOrEmail(ctx, in.Username)
	if err == nil {
		return exception.New(400, "Username already taken")
	}

	// hash password
	hashedPassword, err := util.HashPassword(in.Password)
	if err != nil {
		return exception.New(400, "Error occure while hashing password")
	}
	// generate random string for id
	uuid, err := util.GenerateRandomString(20)
	if err != nil {
		return exception.New(400, "Error occure while generate uuid")
	}
	// make object payload
	var userModel = entity.User{
		ID:           uuid,
		Email:        in.Email,
		Username:     in.Username,
		HashPassword: hashedPassword,
	}
	// passing into user repo
	err = u.userRepository.Create(ctx, userModel)
	if err != nil {
		return exception.New(556, err.Error())
	}

	return nil
}

func NewUserUsecase(userRepository repository.UserRepository, config *config.AppConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository,
		config,
	}
}
