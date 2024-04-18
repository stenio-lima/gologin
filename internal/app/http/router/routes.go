package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stenio-lima/gologin/internal/app/http/handlers"
)

func initRoutes(router *fiber.App) {
	v1 := router.Group("/api/v1")
	{
		v1.Post("/login", handlers.CreateLoginHandler)
		v1.Post("/register", handlers.CreateRegisterHandler)
		v1.Post("/forgot-password", handlers.ForgotPasswordHandler)

	}
}
