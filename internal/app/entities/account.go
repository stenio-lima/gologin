package entities

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Id        uint           `json:"id_account" gorm:"type:uuid;column:id_account;primaryKey"`
	Nome      string         `json:"nome" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(50);unique;not null"`
	Password  string         `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Account) TableName() string {
	return "account"
}
