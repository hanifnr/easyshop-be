package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FPurcNew struct{}

func (f FPurcNew) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	purc := m.(*model.Purc)

	listPurcd := make([]*model.Purcd, 0)
	db.Where("purc_id=?", purc.Id).Find(&listPurcd)
	for _, purcd := range listPurcd {
		if err := db.Exec("UPDATE orderd SET qtypurc = ?, imported = TRUE WHERE order_id = ? AND dno = ?",
			purcd.Qty, purcd.OrderId, purcd.OrderDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}
	}

	return utils.StatusReturnOK()
}
