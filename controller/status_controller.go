package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var CreateStatus = func(w http.ResponseWriter, r *http.Request) {
	statusController := &StatusController{}
	CreateModelAction(statusController, w, r)
}

var UpdateStatus = func(w http.ResponseWriter, r *http.Request) {
	statusController := &StatusController{}
	UpdateModelAction(statusController, w, r)
}

var ViewStatus = func(w http.ResponseWriter, r *http.Request) {
	statusController := &StatusController{}
	ViewModelAction(statusController, w, r)
}

var ListStatus = func(w http.ResponseWriter, r *http.Request) {
	statusController := &StatusController{}
	ListModelAction(statusController, w, r)
}

type StatusController struct {
	Status model.Status
}

func (statusController *StatusController) Model() model.Model {
	return &statusController.Status
}

func (statusController *StatusController) FNew() functions.SQLFunction {
	return nil
}

func (statusController *StatusController) CreateModel() map[string]interface{} {

	if retval := CreateModel(statusController, func(m model.Model) {
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, statusController.Status)
}

func (statusController *StatusController) ViewModel(id int64) map[string]interface{} {
	status := &model.Status{}
	if retval := ViewModel(id, status); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, status)
}

func (statusController *StatusController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(statusController, &model.Status{}, func(modelSrc, modelTemp model.Model) {
		statusSrc := modelSrc.(*model.Status)
		statusTemp := modelTemp.(*model.Status)

		statusSrc.Name = statusTemp.Name
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (statusController *StatusController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("status", "id ASC", make([]*model.Status, 0), param)
}
