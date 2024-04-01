package routes

import (
	"github.com/EraldCaka/rentio/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(userHandler *handlers.ClientHandler, route fiber.Router) {

	route.Get("/clients/:id", userHandler.GetClientByID)
	route.Put("/clients/:id", userHandler.UpdateClient)
	route.Delete("/clients/:id", userHandler.DeleteClient)
	route.Get("/clients", userHandler.GetClients)
	route.Post("/clients/Register", userHandler.CreateClient)
	route.Post("/clients/login", userHandler.LoginClient)
}
