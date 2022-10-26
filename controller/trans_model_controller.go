package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
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
