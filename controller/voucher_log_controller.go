package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var CreateVoucherLog = func(w http.ResponseWriter, r *http.Request) {
	voucherLogController := &VoucherLogController{}
	CreateModelAction(voucherLogController, w, r)
}

var UpdateVoucherLog = func(w http.ResponseWriter, r *http.Request) {
	voucherLogController := &VoucherLogController{}
	UpdateModelAction(voucherLogController, w, r)
}

var ViewVoucherLog = func(w http.ResponseWriter, r *http.Request) {
	voucherLogController := &VoucherLogController{}
	ViewModelAction(voucherLogController, w, r)
}

var ListVoucherLog = func(w http.ResponseWriter, r *http.Request) {
	voucherLogController := &VoucherLogController{}
	ListModelAction(voucherLogController, w, r)
}

type VoucherLogController struct {
	VoucherLog model.VoucherLog
}

func (voucherLogController *VoucherLogController) Model() model.Model {
	return &voucherLogController.VoucherLog
}

func (voucherLogController *VoucherLogController) FNew() functions.SQLFunction {
	return nil
}
func (voucherLogController *VoucherLogController) FDelete() functions.SQLFunction {
	return nil
}

func (voucherLogController *VoucherLogController) CreateModel() map[string]interface{} {

	if retval := CreateModel(voucherLogController, func(m model.Model) {

	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, voucherLogController.VoucherLog)
}

func (voucherLogController *VoucherLogController) ViewModel(id int64) map[string]interface{} {
	voucherLog := &model.VoucherLog{}
	if retval := ViewModel(id, voucherLog); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, voucherLog)
}

func (voucherLogController *VoucherLogController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(voucherLogController, &model.VoucherLog{}, func(modelSrc, modelTemp model.Model) {
		// voucherLogSrc := modelSrc.(*model.VoucherLog)
		// voucherLogTemp := modelTemp.(*model.VoucherLog)

	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (voucherLogController *VoucherLogController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, voucherLogController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (voucherLogController *VoucherLogController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("voucher_log", "id ASC", &voucherLogController.VoucherLog, make([]*model.VoucherLog, 0), param)
}
