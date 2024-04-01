package router

import (
	"context"
	"fmt"
	"github.com/EraldCaka/rentio/db"
	"github.com/EraldCaka/rentio/internal/handlers"
	"github.com/EraldCaka/rentio/internal/middleware"
	"github.com/EraldCaka/rentio/internal/routes"
	"github.com/EraldCaka/rentio/internal/services"
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/EraldCaka/rentio/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var config = fiber.Config{
	ErrorHandler: types.ErrorHandler,
}

var api *fiber.App

func NewRouter() {
	api = fiber.New(config)
	database, err := db.NewPGInstance(context.Background())
	if err != nil {
		fmt.Println(types.NewError(500, fmt.Sprintf("Could not initialize database connection: %s", err)))
		return
	}

	api.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5555",
		AllowMethods:     "GET, POST, DELETE, PUT",
		AllowHeaders:     "Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	var (
		userService   = services.NewUserService(database)
		clientService = services.NewClientService(database)
	)
	var (
		userHandler   = handlers.NewUserHandler(userService)
		clientHandler = handlers.NewClientHandler(clientService)
	)

	var route = api.Group("/rentio/api/v1")
	route.Use(middleware.AuthMiddleware(util.Secret))
	routes.UserRoutes(userHandler, route)
	routes.ClientRoutes(clientHandler, route)
}
func Start(address string) error {
	return api.Listen(address)
}
