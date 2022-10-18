package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type SQLFunction interface {
	Run(m model.Model, db *gorm.DB) utils.StatusReturn
}
