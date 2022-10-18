package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Controller interface {
	Model() model.Model
	CreateModel() map[string]interface{}
	FNew() functions.SQLFunction
}

func Save(controller Controller) utils.StatusReturn {
	db := utils.GetDB()
	tx := db.Begin()
	if retval := CreateModel(controller, tx); retval.ErrCode != 0 {
		tx.Rollback()
		return retval
	}
	tx.Commit()
	return utils.StatusReturnOK()
}

func CreateModel(controller Controller, db *gorm.DB) utils.StatusReturn {
	fNew := controller.FNew()
	model := controller.Model()
	if err := model.Validate(); err != nil {
		db.Rollback()
		return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: err.Error()}
	}
	if fNew != nil {
		if retval := fNew.Run(model, db); retval.ErrCode != 0 {
			return retval
		}
	}
	if err := db.Create(model).Error; err != nil {
		return utils.StatusReturn{ErrCode: utils.ErrSQLCreate, Message: err.Error()}
	}
	return utils.StatusReturnOK()
}

func CreateModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(controller.Model()); err != nil {
		data := utils.Message(false, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.CreateModel()
	utils.Respond(w, resp)
}
