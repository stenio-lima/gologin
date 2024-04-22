package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitializeConfigs() {
	dbConfig := new(DbConfig)
	logger = NewLogger("gologin")

	err := dbConfig.GetEnvDatabase()
	if err != nil {
		logger.Error("Error when trying get environment data from database. Reason:", err.Error())
	}

}
