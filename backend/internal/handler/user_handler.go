package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hendrialqori/war-ticket/backend/internal/dto"
	"github.com/hendrialqori/war-ticket/backend/internal/dto/mapper"
	"github.com/hendrialqori/war-ticket/backend/internal/exception"
	"github.com/hendrialqori/war-ticket/backend/internal/usecase"
	"github.com/hendrialqori/war-ticket/backend/internal/util"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	SetActive(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
}

type userHandlerImpl struct {
	userUsecase usecase.UserUsecase
	validate    *validator.Validate
}

// GetProfile implements [UserHandler].
func (u *userHandlerImpl) GetProfile(c *fiber.Ctx) error {
	ctx := c.Context()

	credential := util.GetCredential(c)

	result, err := u.userUsecase.GetProfile(ctx, credential.ID)
	if err != nil {
		return err
	}

	user := mapper.ToUserDTO(result)

	return util.MapToResponse(c, 200, user, "Success retrieve credential")
}

// Active implements [UserHandler].
func (u *userHandlerImpl) SetActive(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.SetActiveRequest

	if err := c.BodyParser(&req); err != nil {
		return exception.New(400, "Invalid request")
	}

	err := u.validate.Struct(req)
	if err != nil {
		return err
	}

	err = u.userUsecase.SetActive(ctx, req)
	if err != nil {
		return err
	}

	return util.MapToResponse(c, 200, nil, "User activated")
}

func (u *userHandlerImpl) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return exception.New(400, "Invalid request")
	}

	err := u.validate.Struct(req)
	if err != nil {
		return err
	}

	err = u.userUsecase.Register(ctx, req)
	if err != nil {
		return err
	}

	return util.MapToResponse(c, 201, nil, "User registered")
}

func (u *userHandlerImpl) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return exception.New(400, "Invalid request")
	}

	err := u.validate.Struct(req)
	if err != nil {
		return err
	}

	jwtToken, err := u.userUsecase.Login(ctx, req)
	if err != nil {
		return err
	}

	return util.MapToResponse(c, 201, jwtToken, "You're logged in")
}

func NewUserHandler(userUsecase usecase.UserUsecase, validate *validator.Validate) UserHandler {
	return &userHandlerImpl{
		userUsecase,
		validate,
	}
}
