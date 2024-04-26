package schemas

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement:false"`
	Email    string
	Password string
}
