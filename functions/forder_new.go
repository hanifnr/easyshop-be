package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FOrderNew struct{}

func (f FOrderNew) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	order := m.(*model.Order)

	if !IsDataExist(order.CustId, &model.Cust{}, db) {
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: "Customer inactive!"}
	}

	flsoNew := &FOrderLogNew{}
	if retval := flsoNew.Run(order, db); retval.ErrCode != 0 {
		return retval
	}

	return utils.StatusReturnOK()
}
