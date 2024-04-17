package router

import (
	"github.com/gofiber/fiber/v3"
)

func initRoutes(router *fiber.App) {
	v1 := router.Group("/api/v1")
	{
		v1.Post("/register", func(c fiber.Ctx) error {
			return c.JSON("Faça o seu Registro")
		})
		v1.Post("/login", func(c fiber.Ctx) error {
			return c.JSON("Faça o seu Login")
		})
		v1.Post("/forgot-password", func(c fiber.Ctx) error {
			return c.JSON("Redefina a sua Senha")
		})
		v1.Post("/users", func(c fiber.Ctx) error {
			return c.JSON("Receber os dados de usuário")
		})
		v1.Get("/users", func(c fiber.Ctx) error {
			return c.JSON("Retornar a listagem de todos os usuários no BD")
		})
		v1.Get("/users/:id", func(c fiber.Ctx) error {
			return c.JSON("Retornar os dados de um usuário especifico")
		})
		v1.Put("/users/:id", func(c fiber.Ctx) error {
			return c.JSON("Atualizar os dados de um usuário especifico")
		})
		v1.Delete("/users/:id", func(c fiber.Ctx) error {
			return c.JSON("Deletar os dados de um usuário especifico")
		})
	}
}
