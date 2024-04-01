package handlers

import (
	"github.com/EraldCaka/rentio/internal/services"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ClientHandler struct {
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandler {
	return &ClientHandler{clientService: clientService}
}

func (h *ClientHandler) CreateClient(ctx *fiber.Ctx) error {
	var client *types.ClientRegisterRequest
	if err := ctx.BodyParser(&client); err != nil {
		return types.ErrBadRequest()
	}
	if validate := client.Validate(); len(validate) > 0 {
		return ctx.JSON(validate)
	}
	clientID, err := h.clientService.CreateClient(ctx.Context(), client)
	if err != nil {
		return err
	}
	return ctx.JSON(clientID)
}

func (h *ClientHandler) GetClients(ctx *fiber.Ctx) error {
	clients, err := h.clientService.GetAllClients(ctx.Context())
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(clients)
}

func (h *ClientHandler) GetClientByID(ctx *fiber.Ctx) error {
	var clientID = ctx.Params("id")
	client, err := h.clientService.GetClientByID(ctx.Context(), clientID)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(client)
}

func (h *ClientHandler) UpdateClient(ctx *fiber.Ctx) error {
	var (
		clientID  = ctx.Params("id")
		clientReq *types.ClientRequest
	)
	if err := ctx.BodyParser(&clientReq); err != nil {
		return types.ErrBadRequest()
	}
	if err := h.clientService.UpdateClient(ctx.Context(), clientID, clientReq); err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(clientID)
}

func (h *ClientHandler) DeleteClient(ctx *fiber.Ctx) error {
	var clientID = ctx.Params("id")
	err := h.clientService.DeleteClient(ctx.Context(), clientID)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(clientID)
}

func (h *ClientHandler) LoginClient(ctx *fiber.Ctx) error {
	var client *types.ClientRequest
	if err := ctx.BodyParser(&client); err != nil {
		return types.ErrBadRequest()
	}
	if err := h.clientService.LoginClient(ctx.Context(), client); err != nil {
		return ctx.JSON(err)
	}
	tokenData := types.JWTToken{
		Username: client.Username,
		Password: client.Password,
	}
	token, err := services.CreateJWTToken(tokenData, util.Secret, time.Hour)
	if err != nil {
		return err
	}
	ctx.Set("Authorization", token)
	return ctx.JSON(fiber.Map{"token": token})
}
