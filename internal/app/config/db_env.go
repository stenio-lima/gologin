package config

import (
	"errors"
	"os"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
)

type DbConfig struct {
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
}

func (e *DbConfig) GetEnvDatabase() error {
	e.DbUser = os.Getenv(DB_USER)
	e.DbPassword = os.Getenv(DB_PASSWORD)
	e.DbPort = os.Getenv(DB_PORT)
	e.DbName = os.Getenv(DB_NAME)

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
