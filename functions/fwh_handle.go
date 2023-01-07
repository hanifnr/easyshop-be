package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FWhHandle struct {
	Status string
}

func (f FWhHandle) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	wh := m.(*model.Wh)

	details := make([]*model.Whd, 0)
	db.Where("wh_id = ?", wh.Id).Find(&details)

	for _, whd := range details {
		purcd := &model.Purcd{}
		db.Where("purc_id = ? AND dno = ?", whd.PurcId, whd.PurcDno).Find(&purcd)

		orderd := &model.Orderd{}
		db.Where("order_id = ? AND dno =?", purcd.OrderId, purcd.OrderDno).Find(&orderd)

		var qtyArrived float32
		db.Select("SUM(qty)").Table("whd").
			Joins("JOIN wh ON wh.id = whd.wh_id").
			Joins("JOIN (SELECT purc_id,dno FROM purcd WHERE order_id = ?) pd ON whd.purc_id = pd.purc_id AND whd.purc_dno = pd.dno", purcd.OrderId).
			Where("status_code = 'IA' AND is_delete = FALSE").Scan(&qtyArrived)

		var arrived bool
		if f.Status == "IA" {
			arrived = orderd.Qty == whd.Qty+qtyArrived
		} else {
			arrived = false
		}

		if err := db.Exec("UPDATE orderd SET arrived = ? WHERE order_id = ? AND dno = ?", arrived, orderd.OrderId, orderd.Dno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLUpdate, Message: err.Error()}
		}
	}
	return utils.StatusReturnOK()
}
