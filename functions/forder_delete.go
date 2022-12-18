package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FOrderDelete struct{}

func (f FOrderDelete) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	order := m.(*model.Order)

	var exist int
	db.Select("1").Table("purcd").Joins("JOIN purc ON purcd.purc_id = purc.id").Where("order_id = ? AND is_delete = FALSE", order.Id).Scan(&exist)
	if exist == 1 {
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: "order already imported to purchasing"}
	}
	return utils.StatusReturnOK()
}
