package router

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func Init() {
	// Inicializa o Router com as configurações padrão do Fiber
	router := fiber.New()

	// Inicializar as Rotas
	initRoutes(router)

	// Inicia o server
	log.Fatal(router.Listen(":3000"))
}
