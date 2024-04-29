package database

import (
	"fmt"
	"github.com/stenio-lima/gologin/internal/app/config"
	"github.com/stenio-lima/gologin/internal/app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbConfig := &config.DbConfig{}
	err := dbConfig.GetEnvDatabase()
	if err != nil {
		fmt.Println("Error connecting to configs to the database")
	}

	// Connect to PostgreSQL
	dbDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbName, dbConfig.DbPort)

	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("postgres opening error:", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&entities.Account{})
	if err != nil {
		fmt.Println("postgres migration error:", err)
		return nil, err
	}

	return db, nil

}
