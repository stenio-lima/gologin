package main

import (
	config2 "github.com/stenio-lima/gologin/internal/app/config"
	"github.com/stenio-lima/gologin/internal/app/http/router"
)

var (
	logger *config2.Logger
)

func main() {
	logger = config2.GetLogger("main")
	// Initialize Configs
	config2.InitializeConfigs()
	// Initialize Router
	router.Init()
}
