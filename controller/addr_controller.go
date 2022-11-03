package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	CreateModelAction(addrController, w, r)
}

var UpdateAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	UpdateModelAction(addrController, w, r)
}

var ViewAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	ViewModelAction(addrController, w, r)
}

var ListAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	ListModelAction(addrController, w, r)
}

type AddrController struct {
	Addr model.Addr
}

func (addrController *AddrController) Model() model.Model {
	return &addrController.Addr
}

func (addrController *AddrController) FNew() functions.SQLFunction {
	return nil
}

func (addrController *AddrController) CreateModel() map[string]interface{} {

	if retval := CreateModel(addrController, func(m model.Model) {
		currentTime := time.Now()

		addr := m.(*model.Addr)
		addr.CreatedAt = currentTime
		addr.UpdatedAt = currentTime
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, addrController.Addr)
}

func (addrController *AddrController) ViewModel(id int64) map[string]interface{} {
	addr := &model.Addr{}
	if retval := ViewModel(id, addr); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, addr)
}

func (addrController *AddrController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(addrController, &model.Addr{}, func(modelSrc, modelTemp model.Model) {
		addrSrc := modelSrc.(*model.Addr)
		addrTemp := modelTemp.(*model.Addr)

		addrSrc.Name = addrTemp.Name
		addrSrc.Email = addrTemp.Email
		addrSrc.CountryCode = addrTemp.CountryCode
		addrSrc.PhoneNumber = addrTemp.PhoneNumber
		addrSrc.ZipCode = addrTemp.ZipCode
		addrSrc.Province = addrTemp.Province
		addrSrc.City = addrTemp.City
		addrSrc.FullAddress = addrTemp.FullAddress
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (addrController *AddrController) ListModel(page int) map[string]interface{} {
	return ListModel(page, "addr", make([]*model.Addr, 0))
}
