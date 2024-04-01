package handlers

import "github.com/gofiber/fiber/v2"

type User interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUsers(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	LogoutUser(ctx *fiber.Ctx) error
	ActiveUserContract(ctx *fiber.Ctx) error
	AllUserContract(ctx *fiber.Ctx) error
}
type Client interface {
	CreateClient(ctx *fiber.Ctx) error
	GetClient(ctx *fiber.Ctx) error
	GetClientByID(ctx *fiber.Ctx) error
	UpdateClient(ctx *fiber.Ctx) error
	DeleteClient(ctx *fiber.Ctx) error
	LoginClient(ctx *fiber.Ctx) error
	ActiveClientContract(ctx *fiber.Ctx) error
	AllClientContract(ctx *fiber.Ctx) error
}
