package functions

import (
	"easyshop/model"
	"easyshop/utils"
	"time"

	"gorm.io/gorm"
)

type FOrderLogNew struct {
	Note string
}

func (f FOrderLogNew) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	order := m.(*model.Order)
	lso := &model.OrderLog{
		OrderId:    order.Id,
		StatusCode: order.StatusCode,
		Note:       f.Note,
		Date:       time.Now(),
	}
	if err := model.Create(lso, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	return utils.StatusReturnOK()
}
