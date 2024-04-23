package main

import (
	"fmt"
	"github.com/stenio-lima/gologin/internal/app/config"
	"github.com/stenio-lima/gologin/internal/app/http/router"
	"log"
)

func main() {
	config.InitializeConfigs() // Initialize Configs

	server := router.InitRoutes() // Initialize Router

	log.Fatal(server.Listen(fmt.Sprint(config.PortServer))) // Start server
}
