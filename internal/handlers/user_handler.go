package handlers

import (
	"github.com/EraldCaka/rentio/internal/services"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user *types.UserCreateRequest
	if err := ctx.BodyParser(&user); err != nil {
		return types.ErrBadRequest()
	}
	if validate := user.Validate(); len(validate) > 0 {
		return ctx.JSON(validate)
	}
	userID, err := h.userService.Create(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) GetUsers(ctx *fiber.Ctx) error {
	users, err := h.userService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(users)
}

func (h *UserHandler) ActiveUserContract(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	user, err := h.userService.GetActiveUserContract(ctx.Context(), userID)
	if err != nil {
		ctx.JSON(err)
	}
	return ctx.JSON(user)
}

func (h *UserHandler) AllUserContract(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	user, err := h.userService.GetAllUserContract(ctx.Context(), userID)
	if err != nil {
		ctx.JSON(err)
	}
	return ctx.JSON(user)
}

func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	user, err := h.userService.GetByID(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(user)
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var (
		userID  = ctx.Params("id")
		userReq *types.UserRequest
	)
	if err := ctx.BodyParser(&userReq); err != nil {
		return types.ErrBadRequest()
	}
	if err := h.userService.Update(ctx.Context(), userID, userReq); err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	err := h.userService.Delete(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var user *types.UserRequest
	if err := ctx.BodyParser(&user); err != nil {
		return types.ErrBadRequest()
	}
	if err := h.userService.Login(ctx.Context(), user); err != nil {
		return ctx.JSON(err)
	}
	tokenData := types.JWTToken{
		Username: user.Username,
		Password: user.Password,
	}
	token, err := services.CreateJWTToken(tokenData, util.Secret, time.Hour)
	if err != nil {
		return err
	}
	ctx.Set("Authorization", token)
	return ctx.JSON(fiber.Map{"token": token})
}
