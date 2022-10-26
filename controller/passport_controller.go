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

var ListPassport = func(w http.ResponseWriter, r *http.Request) {
	passportController := &PassportController{}
	ListModelAction(passportController, w, r)
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

func (passportController *PassportController) CreateModel() map[string]interface{} {

	if retval := CreateModel(passportController, func(m model.Model) {}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, passportController.Passport)
}

func (passportController *PassportController) ViewModel(id int64) map[string]interface{} {
	passport := &model.Passport{}
	if retval := ViewModel(id, passport); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, passport)
}

func (passportController *PassportController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(passportController, &model.Passport{}, func(modelSrc, modelTemp model.Model) {
		passportSrc := modelSrc.(*model.Passport)
		passportTemp := modelTemp.(*model.Passport)

		passportSrc.Name = passportTemp.Name
		passportSrc.CountryCode = passportTemp.CountryCode
		passportSrc.Number = passportTemp.Number
		passportSrc.Nationality = passportTemp.Nationality
		passportSrc.BirthDate = passportTemp.BirthDate
		passportSrc.IssueDate = passportTemp.IssueDate
		passportSrc.ExpDate = passportTemp.ExpDate
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (passportController *PassportController) ListModel(page int) map[string]interface{} {
	return ListModel(page, "passport", make([]*model.Passport, 0))
}
