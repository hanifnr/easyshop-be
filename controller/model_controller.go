package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"reflect"
	"time"

	"gorm.io/gorm"
)

func CreateModel(controller Controller, fDefaultValue func(m model.Model)) utils.StatusReturn {
	return CreateModelWithPostSave(controller, fDefaultValue, func(db *gorm.DB) utils.StatusReturn {
		return utils.StatusReturnOK()
	})
}

func CreateModelWithPostSave(controller Controller, fDefaultValue func(m model.Model), fPostSave func(db *gorm.DB) utils.StatusReturn) utils.StatusReturn {
	fDefaultValue(controller.Model())

	db := utils.GetDB().Begin()
	fNew := controller.FNew()
	m := controller.Model()

	if t, ok := m.(model.TimeField); ok {
		currentTime := time.Now()
		t.SetCreatedAt(currentTime)
		t.SetUpdatedAt(currentTime)
	}

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
	if retval := fPostSave(db); retval.ErrCode != 0 {
		db.Rollback()
		return utils.StatusReturn{ErrCode: retval.ErrCode, Message: retval.Message}
	}
	if modelExt, ok := m.(model.ModelExt); ok {
		modelExt.SetValueModelExt(db)
	}
	db.Commit()

	return utils.StatusReturnOK()
}

func ViewModel(id int64, m model.Model) utils.StatusReturn {
	db := utils.GetDB()
	if err := model.Load(id, m, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}

	if modelExt, ok := m.(model.ModelExt); ok {
		modelExt.SetValueModelExt(db)
	}

	return utils.StatusReturnOK()
}

func UpdateModel(controller Controller, m model.Model, fUpdate func(modelSrc model.Model, modelTemp model.Model)) (utils.StatusReturn, model.Model) {
	db := utils.GetDB().Begin()
	modelTemp := controller.Model()
	if err := modelTemp.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}, nil
	}
	if err := model.Load(modelTemp.ID(), m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}, nil
	}

	if t, ok := m.(model.TimeField); ok {
		t.SetUpdatedAt(time.Now())
	}

	fUpdate(m, modelTemp)
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}, nil
	}
	db.Commit()
	return utils.StatusReturnOK(), m
}

func DeleteModel(id int64, controller Controller, fAction func(m model.Model) utils.StatusReturn) utils.StatusReturn {
	db := utils.GetDB().Begin()
	m := controller.Model()
	fDelete := controller.FDelete()

	if retval := ViewModel(id, m); retval.ErrCode != 0 {
		db.Rollback()
		return retval
	}
	if deleteField, ok := m.(model.DeleteField); ok {
		deleteField.SetIsDelete(true)
	}
	if timeField, ok := m.(model.TimeField); ok {
		timeField.SetUpdatedAt(time.Now())
	}
	if retval := fAction(m); retval.ErrCode != 0 {
		db.Rollback()
		return retval
	}
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	if fDelete != nil {
		if retval := fDelete.Run(m, db); retval.ErrCode != 0 {
			db.Rollback()
			return retval
		}
	}
	db.Commit()
	return utils.StatusReturnOK()
}

func RemoveModel(id int64, controller Controller, fAction func(m model.Model) utils.StatusReturn) utils.StatusReturn {
	db := utils.GetDB().Begin()
	m := controller.Model()
	fDelete := controller.FDelete()

	if retval := fAction(m); retval.ErrCode != 0 {
		db.Rollback()
		return retval
	}
	if fDelete != nil {
		if retval := fDelete.Run(m, db); retval.ErrCode != 0 {
			db.Rollback()
			return retval
		}
	}
	if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLDelete, Message: err.Error()}
	}
	db.Commit()
	return utils.StatusReturnOK()
}

func ListModel(table, order string, m interface{}, list interface{}, param *utils.Param) map[string]interface{} {
	ProcessDefaultModelParam(m, param)
	return ListJoinModel(table, order, list, param, func(query *gorm.DB) {}, func(query *gorm.DB) {})
}

func ListJoinModel(table, order string, list interface{}, param *utils.Param, fJoin func(query *gorm.DB), fFilter func(query *gorm.DB)) map[string]interface{} {
	db := utils.GetDB()

	respPage, err := utils.QueryListFind(table, order, &list, param, fJoin, fFilter)
	if err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	ProcessExtField(list, db)
	return utils.MessageListData(true, list, respPage)
}

func ListModelColumns(columns, table, order string, list interface{}, param *utils.Param) map[string]interface{} {
	db := utils.GetDB()

	respPage, err := utils.QueryListScan(columns, table, order, &list, param)
	if err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	ProcessExtField(list, db)
	return utils.MessageListData(true, list, respPage)
}

func ProcessExtField(list interface{}, db *gorm.DB) {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)

		for i := 0; i < s.Len(); i++ {
			if v, ok := s.Index(i).Interface().(model.ModelExt); ok {
				v.SetValueModelExt(db)
			}
		}
	}
}

func UpdateFieldModel(id int64, controller Controller, fAction func(m model.Model) utils.StatusReturn) map[string]interface{} {
	return UpdateFieldModelWithPostSave(id, controller, fAction, func() utils.StatusReturn {
		return utils.StatusReturnOK()
	})
}

func UpdateFieldModelWithPostSave(id int64, controller Controller, fAction func(m model.Model) utils.StatusReturn, fPostSave func() utils.StatusReturn) map[string]interface{} {
	db := utils.GetDB().Begin()
	m := controller.Model()
	if retval := ViewModel(id, m); retval.ErrCode != 0 {
		db.Rollback()
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	if retval := fAction(m); retval.ErrCode != 0 {
		db.Rollback()
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	if retval := fPostSave(); retval.ErrCode != 0 {
		db.Rollback()
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	db.Commit()
	return utils.MessageData(true, m)
}

func ProcessDefaultModelParam(m interface{}, param *utils.Param) {
	if _, ok := m.(model.DeleteField); ok {
		param.SetDefaultDelete()
	}
}
