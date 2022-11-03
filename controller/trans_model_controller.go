package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"reflect"
	"time"

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

	if t, ok := controller.MasterModel().(model.TimeField); ok {
		currentTime := time.Now()
		t.SetCreatedAt(currentTime)
		t.SetUpdatedAt(currentTime)
	}

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

	if v, ok := controller.MasterModel().(model.ModelExt); ok {
		v.SetValueModelExt(db)
	}

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

func UpdateTrans(controller TransController, m model.Model, d model.Model, fUpdate func(modelSrc model.Model, modelTemp model.Model)) utils.StatusReturn {
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
	fUpdate(m, modelTemp)
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

func ListTrans(page int, table string, list interface{}, param *Param) map[string]interface{} {
	db := utils.GetDB()

	var query *gorm.DB
	if page == 0 {
		query = db.Order("id ASC")
	} else {
		offset, limit := utils.GetOffsetLimit(page)
		query = db.Offset(offset).Order("id ASC").Limit(limit)
	}
	param.ProcessFilter(query)
	totalRow := query.Find(&list).RowsAffected

	if err := query.Error; err != nil {
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
	respPage := utils.RespPage(page, int(totalRow))
	return utils.MessageListData(true, list, respPage)
}
