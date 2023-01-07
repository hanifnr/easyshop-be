package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FWhDelete struct{}

func (f FWhDelete) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	wh := m.(*model.Wh)

	listWhd := make([]*model.Whd, 0)
	db.Where("wh_id=?", wh.Id).Find(&listWhd)
	for _, whd := range listWhd {
		if err := db.Exec("UPDATE purcd SET qtywh = 0, imported = FALSE WHERE purc_id = ? AND dno = ?",
			whd.PurcId, whd.PurcDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}

		purcd := &model.Purcd{}
		db.Where("purc_id = ? AND dno = ?", whd.PurcId, whd.PurcDno).Find(&purcd)

		orderd := &model.Orderd{}
		db.Where("order_id = ? AND dno =?", purcd.OrderId, purcd.OrderDno).Find(&orderd)

		qtyWh := orderd.Qtywh - whd.Qty
		if err := db.Exec("UPDATE orderd SET qtywh = ? WHERE order_id = ? AND dno = ? AND arrived = FALSE",
			qtyWh, purcd.OrderId, purcd.OrderDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}

		//cek apabila item order sudah diimport ke wh semua set status IR
		var importedOrderWh int
		db.Select("COUNT(*)").Table("orderd").Where("order_id = ? AND qty <> qtywh", purcd.OrderId).Scan(&importedOrderWh)

		order := &model.Order{}
		db.Where("id = ?", purcd.OrderId).Find(&order)
		if importedOrderWh == 0 {
			if err := db.Exec("UPDATE public.order SET status_code = 'IR' WHERE id = ?", purcd.OrderId).Error; err != nil {
				return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
			}
			order.StatusCode = "IR"
		} else {
			if err := db.Exec("UPDATE public.order SET status_code = 'IC' WHERE id = ?", purcd.OrderId).Error; err != nil {
				return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
			}
			order.StatusCode = "IC"
		}
		foln := &FOrderLogNew{
			Note: "generated by system",
		}
		foln.Run(order, db)
	}
	return utils.StatusReturnOK()
}
