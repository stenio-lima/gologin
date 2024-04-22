package main

import (
	"github.com/stenio-lima/gologin/internal/app/database/config"
	"github.com/stenio-lima/gologin/internal/app/http/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.InitDB()
	if err != nil {
		logger.Errorf("config database init failed: %v", err)
		return
	}

	// Initialize Router
	router.Init()
}
