package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"

	"gorm.io/gorm"
)

func CreateTrans(controller TransController, fDefaultValue func(m model.Model)) utils.StatusReturn {
	fDefaultValue(controller.MasterModel())
	if err := ValidateTrans(controller); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}

	fNew := controller.FNew()
	masterModel := controller.MasterModel()
	listDetail := controller.DetailsModel()

	db := utils.GetDB().Begin()

	master := masterModel.(model.Master)
	if master.GetTrxno() == "AUTO" {
		trxno := functions.FGetNewNo(masterModel.TableName(), db)
		master.SetTrxno(trxno)
	}
	if err := masterModel.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}
	if err := model.Create(masterModel, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	for _, data := range listDetail {
		if err := data.Validate(); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
		}
		detail := data.(model.Detail)
		detail.SetMasterId(masterModel.ID())
		if err := model.Create(data, db); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
		}
	}
	if fNew != nil {
		if retval := fNew.Run(masterModel, db); retval.ErrCode != 0 {
			db.Rollback()
			return retval
		}
	}
	db.Commit()
	return utils.StatusReturnOK()
}

func ViewTrans(id int64, controller TransController, fLoadDetail func(db *gorm.DB) error) utils.StatusReturn {
	db := utils.GetDB().Begin()
	if err := model.Load(id, controller.MasterModel(), db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	if err := fLoadDetail(db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	return utils.StatusReturnOK()
}

func ListTrans(page int, table string, list interface{}) map[string]interface{} {
	db := utils.GetDB()

	var totalRow int64
	if err := db.Select("count(id)").Table(table).Scan(&totalRow).Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}

	var query *gorm.DB
	if page == 0 {
		query = db.Find(&list).Order("id ASC")
	} else {
		offset, limit := utils.GetOffsetLimit(page)
		query = db.Offset(offset).Order("id ASC").Limit(limit).Find(&list)
	}
	if err := query.Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	respPage := utils.RespPage(page, int(totalRow))
	return utils.MessageListData(true, list, respPage)
}
