package router

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes() *fiber.App {
	router := fiber.New()

	publicRoutes(router)

	return router
}
