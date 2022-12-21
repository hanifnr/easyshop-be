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
	if lso.StatusCode == "IS" {
		imageWritter := utils.GetImageWritter()
		if err := imageWritter.DeleteFile(order.ProofLink); err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrIO, Message: err.Error()}
		}
		order.ProofLink = ""
		if err := model.Save(order, db); err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLUpdate, Message: err.Error()}
		}
	}
	if err := model.Create(lso, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	return utils.StatusReturnOK()
}
