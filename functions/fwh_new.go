package functions

import (
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

type FWhNew struct{}

func (f FWhNew) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	wh := m.(*model.Wh)

	listWhd := make([]*model.Whd, 0)
	db.Where("wh_id=?", wh.Id).Find(&listWhd)
	for _, whd := range listWhd {
		if err := db.Exec("UPDATE purcd SET qtywh = ?, imported = TRUE WHERE purc_id = ? AND dno = ?",
			whd.Qty, whd.PurcId, whd.PurcDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}

		purcd := &model.Purcd{}
		db.Where("purc_id = ? AND dno = ?", whd.PurcId, whd.PurcDno).Find(&purcd)
		if err := db.Exec("UPDATE orderd SET qtywh = ? WHERE order_id = ? AND dno = ?",
			whd.Qty, purcd.OrderId, purcd.OrderDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}

		//cek apabila item order sudah diimport ke wh semua set status IR
		var importedOrderWh int
		db.Select("COUNT(*)").Table("orderd").Where("order_id = ? AND qtywh = 0", purcd.OrderId).Scan(&importedOrderWh)
		if importedOrderWh == 0 {
			if err := db.Exec("UPDATE public.order SET status_code = 'IR' WHERE id = ?", purcd.OrderId).Error; err != nil {
				return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
			}
		}
	}
	return utils.StatusReturnOK()
}
