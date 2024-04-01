package routes

import (
	"github.com/EraldCaka/rentio/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(userHandler *handlers.UserHandler, route fiber.Router) {

	route.Get("/users/:id", userHandler.GetUserByID)
	route.Put("/users/:id", userHandler.UpdateUser)
	route.Delete("/users/:id", userHandler.DeleteUser)
	route.Get("/users", userHandler.GetUsers)
	route.Post("/users/register", userHandler.CreateUser)
	route.Post("/users/login", userHandler.LoginUser)
	route.Post("/users/logout", userHandler.LogoutUser)
	route.Get("/users/contracts/:id", userHandler.AllUserContract)
	route.Get("/users/contracts/active/:id", userHandler.ActiveUserContract)
}
