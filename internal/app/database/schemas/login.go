package schemas

import (
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Login    string // E-mail também seria válido? Consultar com o mestre.
	Password string
}
