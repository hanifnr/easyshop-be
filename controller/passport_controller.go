package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var CreatePassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	CreateModelAction(passportController, w, r)
}

var UpdatePassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	UpdateModelAction(passportController, w, r)
}

var ViewPassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	ViewModelAction(passportController, w, r)
}

var DeletePassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	DeleteModelAction(passportController, w, r)
}

var ListPassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	ListModelAction(passportController, w, r)
}

var ViewPassportCust = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	ViewModelCust(passportController, w, r)
}

type PassportController struct {
	Passport model.Passport
}

func (passportController *PassportController) Model() model.Model {
	return &passportController.Passport
}

func (passportController *PassportController) FNew() functions.SQLFunction {
	return nil
}

func (passportController *PassportController) FDelete() functions.SQLFunction {
	return nil
}

func (passportController *PassportController) CreateModel() map[string]interface{} {

	if retval := CreateModel(passportController, func(m model.Model) {}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, passportController.Passport)
}

func (passportController *PassportController) ViewModel(id int64) map[string]interface{} {
	passport := &passportController.Passport
	if retval := ViewModel(id, passport); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, passport)
}

func (passportController *PassportController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(passportController, &model.Passport{}, func(modelSrc, modelTemp model.Model) {
		passportSrc := modelSrc.(*model.Passport)
		passportTemp := modelTemp.(*model.Passport)

		passportSrc.Number = passportTemp.Number
		passportSrc.Nationality = passportTemp.Nationality
		passportSrc.BirthDate = passportTemp.BirthDate
		passportSrc.StatusResidence = passportTemp.StatusResidence
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (passportController *PassportController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, passportController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (passportController *PassportController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("passport", "id ASC", &passportController.Passport, make([]*model.Passport, 0), param)
}

func ViewModelCust(passportController *PassportController, w http.ResponseWriter, r *http.Request) {
	custId, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	db := utils.GetDB()
	passport := &passportController.Passport
	query := db.Where("cust_id = ?", custId).Find(passport)
	if err := query.Error; err != nil {
		utils.Respond(w, utils.MessageErr(false, utils.ErrSQLLoad, err.Error()))
	} else if query.RowsAffected == 0 {
		utils.Respond(w, utils.MessageErr(false, utils.ErrSQLLoad, utils.RESPONSE_NOT_FOUND))
	} else {
		utils.Respond(w, utils.MessageData(true, passport))
	}
}
