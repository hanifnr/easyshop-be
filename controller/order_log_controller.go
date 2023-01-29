package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var CreateOrderLog = func(w http.ResponseWriter, r *http.Request) {
	orderLogController := &OrderLogController{}
	CreateModelAction(orderLogController, w, r)
}

var UpdateOrderLog = func(w http.ResponseWriter, r *http.Request) {
	orderLogController := &OrderLogController{}
	UpdateModelAction(orderLogController, w, r)
}

var ViewOrderLog = func(w http.ResponseWriter, r *http.Request) {
	orderLogController := &OrderLogController{}
	ViewModelAction(orderLogController, w, r)
}

var ListOrderLog = func(w http.ResponseWriter, r *http.Request) {
	orderLogController := &OrderLogController{}
	ListModelAction(orderLogController, w, r)
}

type OrderLogController struct {
	OrderLog model.OrderLog
}

func (orderLogController *OrderLogController) Model() model.Model {
	return &orderLogController.OrderLog
}

func (orderLogController *OrderLogController) FNew() functions.SQLFunction {
	return nil
}

func (orderLogController *OrderLogController) FDelete() functions.SQLFunction {
	return nil
}

func (orderLogController *OrderLogController) CreateModel() map[string]interface{} {

	if retval := CreateModel(orderLogController, func(m model.Model) {}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderLogController.OrderLog)
}

func (orderLogController *OrderLogController) ViewModel(id int64) map[string]interface{} {
	lso := &model.OrderLog{}
	if retval := ViewModel(id, lso); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, lso)
}

func (orderLogController *OrderLogController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(orderLogController, &model.OrderLog{}, func(modelSrc, modelTemp model.Model) {})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (orderLogController *OrderLogController) DeleteModel(id int64) map[string]interface{} {
	return nil
}

func (orderLogController *OrderLogController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("order_log", "date ASC", &orderLogController.OrderLog, make([]*model.OrderLog, 0), param)
}
