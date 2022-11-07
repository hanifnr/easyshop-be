package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"reflect"
	"time"

	"gorm.io/gorm"
)

func CreateModel(controller Controller, fDefaultValue func(m model.Model)) utils.StatusReturn {
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
	db.Commit()

	if modelExt, ok := m.(model.ModelExt); ok {
		modelExt.SetValueModelExt(db)
	}

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

func ListModel(page int, table string, list interface{}, param *utils.Param) map[string]interface{} {
	db := utils.GetDB()

	respPage, err := utils.QueryListFind(page, table, &list, param)
	if err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	ProcessExtField(list, db)
	return utils.MessageListData(true, list, respPage)
}

func ListModelColumns(page int, columns, table string, list interface{}, param *utils.Param) map[string]interface{} {
	db := utils.GetDB()

	respPage, err := utils.QueryListScan(page, columns, table, &list, param)
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
