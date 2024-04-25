package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/stenio-lima/gologin/internal/app/helpers"
	"os"
)

type DbConfig struct {
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
}

func (e *DbConfig) GetEnvDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	e.DbUser = os.Getenv(helpers.DbUser)
	e.DbPassword = os.Getenv(helpers.DbPassword)
	e.DbPort = os.Getenv(helpers.DbPort)
	e.DbName = os.Getenv(helpers.DbName)

	switch {
	case e.DbUser == "":
		return errors.New("DbUser is empty")
	case e.DbPassword == "":
		return errors.New("DbPassword is empty")
	case e.DbPort == "":
		return errors.New("DbPort is empty")
	case e.DbName == "":
		return errors.New("DbName is empty")
	}

	return nil
}
