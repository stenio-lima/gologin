package router

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func Init() {
	// Inicializa o Router com as configurações padrão do Fiber
	app := fiber.New()

	// Define uma rota e seu tipo através do path '/'
	app.Get("/login", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Inicia o server na porta de sua escolha
	log.Fatal(app.Listen(":3000"))
}
