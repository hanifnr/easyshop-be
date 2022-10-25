package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
)

func CreateTrans(controller TransController, fDefaultValue func(m model.Model)) utils.StatusReturn {
	fDefaultValue(controller.MasterModel())
	db := utils.GetDB().Begin()
	fNew := controller.FNew()
	m := controller.MasterModel()
	master := m.(model.Master)
	if master.GetTrxno() == "AUTO" {
		trxno := functions.FGetNewNo(m.TableName(), db)
		master.SetTrxno(trxno)
	}
	if err := m.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}
	if err := model.Create(m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	for _, d := range controller.DetailsModel() {
		if err := d.Validate(); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
		}
		detail := d.(model.Detail)
		detail.SetMasterId(m.ID())
		if err := model.Create(d, db); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
		}
	}
	if fNew != nil {
		if retval := fNew.Run(m, db); retval.ErrCode != 0 {
			db.Rollback()
			return retval
		}
	}
	db.Commit()
	return utils.StatusReturnOK()
}
