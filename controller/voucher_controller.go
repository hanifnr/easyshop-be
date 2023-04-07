package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateVoucher = func(w http.ResponseWriter, r *http.Request) {
	voucherController := &VoucherController{}
	CreateModelAction(voucherController, w, r)
}

var UpdateVoucher = func(w http.ResponseWriter, r *http.Request) {
	voucherController := &VoucherController{}
	UpdateModelAction(voucherController, w, r)
}

var ViewVoucher = func(w http.ResponseWriter, r *http.Request) {
	voucherController := &VoucherController{}
	ViewModelAction(voucherController, w, r)
}

var ListVoucher = func(w http.ResponseWriter, r *http.Request) {
	voucherController := &VoucherController{}
	ListModelAction(voucherController, w, r)
}

var DeleteVoucher = func(w http.ResponseWriter, r *http.Request) {
	voucherController := &VoucherController{}
	DeleteModelAction(voucherController, w, r)
}

type VoucherController struct {
	Voucher model.Voucher
}

func (voucherController *VoucherController) Model() model.Model {
	return &voucherController.Voucher
}

func (voucherController *VoucherController) FNew() functions.SQLFunction {
	return nil
}
func (voucherController *VoucherController) FDelete() functions.SQLFunction {
	return nil
}

func (voucherController *VoucherController) CreateModel() map[string]interface{} {

	if retval := CreateModel(voucherController, func(m model.Model) {
		currentTime := time.Now()

		voucher := m.(*model.Voucher)
		voucher.CreatedAt = currentTime
		voucher.UpdatedAt = currentTime
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, voucherController.Voucher)
}

func (voucherController *VoucherController) ViewModel(id int64) map[string]interface{} {
	voucher := &model.Voucher{}
	if retval := ViewModel(id, voucher); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, voucher)
}

func (voucherController *VoucherController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(voucherController, &model.Voucher{}, func(modelSrc, modelTemp model.Model) {
		voucherSrc := modelSrc.(*model.Voucher)
		voucherTemp := modelTemp.(*model.Voucher)

		voucherSrc.Code = voucherTemp.Code
		voucherSrc.Amount = voucherTemp.Amount
		voucherSrc.Qty = voucherTemp.Qty
		voucherSrc.Startdate = voucherTemp.Startdate
		voucherSrc.Enddate = voucherTemp.Enddate
		voucherSrc.PartnershipId = voucherTemp.PartnershipId
		voucherSrc.Note = voucherTemp.Note
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (voucherController *VoucherController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, voucherController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (voucherController *VoucherController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("voucher", "id ASC", voucherController.Voucher, make([]*model.Voucher, 0), param)
}
