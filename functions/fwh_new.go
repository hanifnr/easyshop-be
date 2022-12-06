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
		if err := db.Exec("UPDATE purcd SET qtywh = ? WHERE purc_id = ? AND dno = ?",
			whd.Qty, whd.PurcId, whd.PurcDno).Error; err != nil {
			return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
		}
		var importedPurcd int
		db.Select("COUNT(*)").Table("purcd").Where("purc_id = ? AND imported = FALSE", whd.PurcId).Scan(&importedPurcd)
		if importedPurcd == 0 {
			if err := db.Exec("UPDATE purc SET status_code = 'IR' WHERE id = ?", whd.PurcId).Error; err != nil {
				return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
			}
		}
	}
	return utils.StatusReturnOK()
}
