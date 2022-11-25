package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"strconv"
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

var DeleteAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	DeleteModelAction(addrController, w, r)
}

var ListAddr = func(w http.ResponseWriter, r *http.Request) {
	addrController := &AddrController{}
	ListModelAction(addrController, w, r)
}

var ListComboAddr = func(w http.ResponseWriter, r *http.Request) {
	paramPage := r.URL.Query().Get("page")
	if paramPage == "" {
		paramPage = "0"
	}
	page, err := strconv.Atoi(paramPage)
	if err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	param := utils.ProcessParam(r)
	resp := ComboAddr(page, param)
	utils.Respond(w, resp)
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

func (addrController *AddrController) FDelete() functions.SQLFunction {
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

func (addrController *AddrController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, addrController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (addrController *AddrController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("addr", "id ASC", make([]*model.Addr, 0), param)
}

func ComboAddr(page int, param *utils.Param) map[string]interface{} {
	return GetCombo(page, "addr", "id ASC", param)
}
