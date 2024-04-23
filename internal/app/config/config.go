package config

import (
	"github.com/stenio-lima/gologin/internal/app/helpers"
	"gorm.io/gorm"
	"os"
)

var (
	db         *gorm.DB
	logger     *Logger
	PortServer string
)

func InitializeConfigs() {
	dbConfig := new(DbConfig)
	logger = NewLogger("gologin")

	err := dbConfig.GetEnvDatabase()
	if err != nil {
		logger.Error("Error when trying get environment data from database. Reason:", err.Error())
	}

	PortServer = os.Getenv(helpers.PortServer)
}
