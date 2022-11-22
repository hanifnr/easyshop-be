package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"reflect"
	"time"

	"gorm.io/gorm"
)

func CreateTrans(controller TransController, fDefaultValue func(db *gorm.DB) error) utils.StatusReturn {
	fNew := controller.FNew()
	masterModel := controller.MasterModel()

	if t, ok := masterModel.(model.TimeField); ok {
		currentTime := time.Now()
		t.SetCreatedAt(currentTime)
		t.SetUpdatedAt(currentTime)
	}

	db := utils.GetDB().Begin()

	if err := fDefaultValue(db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}

	if err := ValidateTrans(controller); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}

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

	for _, data := range controller.DetailsModel() {
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

	if v, ok := controller.MasterModel().(model.ModelExt); ok {
		v.SetValueModelExt(db)
	}

	db.Commit()

	return utils.StatusReturnOK()
}

func ViewTrans(id int64, controller TransController, fLoadDetail func(db *gorm.DB) error) utils.StatusReturn {
	db := utils.GetDB()

	if err := model.Load(id, controller.MasterModel(), db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}

	if err := fLoadDetail(db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}

	if v, ok := controller.MasterModel().(model.ModelExt); ok {
		v.SetValueModelExt(db)
	}
	return utils.StatusReturnOK()
}

func UpdateTrans(controller TransController, m model.Model, d model.Model, fUpdate func(modelSrc model.Model, modelTemp model.Model, db *gorm.DB) error) utils.StatusReturn {
	db := utils.GetDB().Begin()
	modelTemp := controller.MasterModel()
	if err := modelTemp.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}
	if err := model.Load(modelTemp.ID(), m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	if t, ok := m.(model.TimeField); ok {
		t.SetUpdatedAt(time.Now())
	}
	if err := fUpdate(m, modelTemp, db); err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLSave, Message: err.Error()}
	}
	if err := db.Debug().Where(controller.MasterField()+"= ?", modelTemp.ID()).Delete(d).Error; err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrSQLDelete, Message: err.Error()}
	}
	for _, data := range controller.DetailsModel() {
		if err := data.Validate(); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
		}
		detail := data.(model.Detail)
		detail.SetMasterId(modelTemp.ID())
		if err := model.Create(data, db); err != nil {
			db.Rollback()
			return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
		}
	}
	db.Commit()
	return utils.StatusReturnOK()
}

func ListTrans(table, order string, list interface{}, param *utils.Param) map[string]interface{} {
	db := utils.GetDB()

	respPage, err := utils.QueryListFind(table, order, &list, param)
	if err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}

	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)

		for i := 0; i < s.Len(); i++ {
			if v, ok := s.Index(i).Interface().(model.ModelExt); ok {
				v.SetValueModelExt(db)
			}
		}
	}
	return utils.MessageListData(true, list, respPage)
}

func UpdateFieldMaster(id int64, controller TransController, fAction func(m model.Model, db *gorm.DB) utils.StatusReturn) map[string]interface{} {
	db := utils.GetDB().Begin()
	m := controller.MasterModel()
	if retval := ViewModel(id, m); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	if retval := fAction(m, db); retval.ErrCode != 0 {
		db.Rollback()
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	db.Commit()
	return utils.MessageData(true, m)
}
