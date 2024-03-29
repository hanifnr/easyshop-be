package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FPurcDelete struct{}

func (f FPurcDelete) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	purc := m.(*model.Purc)

	var exist int
	db.Select("1").Table("whd").Joins("JOIN wh ON whd.wh_id = wh.id").Where("purc_id = ? AND is_delete = FALSE", purc.Id).Scan(&exist)
	if exist == 1 {
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: "purchase already imported to warehouse"}
	}

	listPurcd := make([]*model.Purcd, 0)
	db.Where("purc_id=?", purc.Id).Find(&listPurcd)
	for _, purcd := range listPurcd {
		orderd := &model.Orderd{}
		db.Where("order_id = ? AND dno =?", purcd.OrderId, purcd.OrderDno).Find(&orderd)

		qtyPurc := orderd.Qtypurc - purcd.Qty
		if err := db.Exec("UPDATE orderd SET qtypurc = ?, imported = FALSE WHERE order_id = ? AND dno = ?",
			qtyPurc, purcd.OrderId, purcd.OrderDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}
		var importedOrderPurc int
		db.Select("COUNT(*)").Table("orderd").Where("order_id = ? AND qtypurc = 0", purcd.OrderId).Scan(&importedOrderPurc)
		if importedOrderPurc == 0 {
			if err := db.Exec("UPDATE public.order SET status_code = 'IC' WHERE id = ?", purcd.OrderId).Error; err != nil {
				return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
			}
			order := &model.Order{}
			db.Where("id = ?", purcd.OrderId).Find(&order)
			foln := &FOrderLogNew{
				Note: "generated by system",
			}
			foln.Run(order, db)
		} else {
			var countOrderd int

			db.Select("COUNT(*)").Table("orderd").Where("order_id = ?", purcd.OrderId).Scan(&countOrderd)
			order := &model.Order{}
			db.Where("id = ?", purcd.OrderId).Find(&order)
			if importedOrderPurc == countOrderd {
				if err := db.Exec("UPDATE public.order SET status_code = 'PA' WHERE id = ?", purcd.OrderId).Error; err != nil {
					return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
				}
				order.StatusCode = "PA"
			} else {
				if order.StatusCode != "IP" {
					if err := db.Exec("UPDATE public.order SET status_code = 'IP' WHERE id = ?", purcd.OrderId).Error; err != nil {
						return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
					}
					order.StatusCode = "IP"
				}
			}
			foln := &FOrderLogNew{
				Note: "generated by system",
			}
			foln.Run(order, db)
		}
	}
	return utils.StatusReturnOK()
}
