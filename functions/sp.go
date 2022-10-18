package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type SP interface {
	Run(m model.Model, db *gorm.DB) utils.StatusReturn
}
