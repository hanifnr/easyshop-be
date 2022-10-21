package controllers

import (
	"easyshop/model"
	"easyshop/utils"
)

func CreateModel(controller Controller, fDefaultValue func(m model.Model)) utils.StatusReturn {
	fDefaultValue(controller.Model())
	db := utils.GetDB().Begin()
	fNew := controller.FNew()
	m := controller.Model()
	if err := m.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}
	if fNew != nil {
		if retval := fNew.Run(m, db); retval.ErrCode != 0 {
			db.Rollback()
			return retval
		}
	}
	if err := model.Create(m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	db.Commit()
	return utils.StatusReturnOK()
}

func ViewModel(id int64, m model.Model) utils.StatusReturn {
	db := utils.GetDB().Begin()
	if err := model.Load(id, m, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	return utils.StatusReturnOK()
}

func UpdateModel(controller Controller, m model.Model, fUpdate func(modelSrc model.Model, modelTemp model.Model)) (utils.StatusReturn, model.Model) {
	db := utils.GetDB()
	modelTemp := controller.Model()

	if err := model.Load(modelTemp.ID(), m, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}, nil
	}
	fUpdate(m, modelTemp)
	if err := model.Save(m, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}, nil
	}
	return utils.StatusReturnOK(), m
}
