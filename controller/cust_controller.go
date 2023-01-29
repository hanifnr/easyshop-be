package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"strconv"
)

var CreateCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	CreateModelAction(custController, w, r)
}

var UpdateCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	UpdateModelAction(custController, w, r)
}

var ViewCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	ViewModelAction(custController, w, r)
}

var DeleteCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	DeleteModelAction(custController, w, r)
}

var ListCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	ListModelAction(custController, w, r)
}

var ListComboCust = func(w http.ResponseWriter, r *http.Request) {
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
	resp := ComboCust(page, param)
	utils.Respond(w, resp)
}

type CustController struct {
	Cust model.Cust
}

func (custController *CustController) Model() model.Model {
	return &custController.Cust
}

func (custController *CustController) FNew() functions.SQLFunction {
	return &functions.FCustNew{}
}

func (custController *CustController) FDelete() functions.SQLFunction {
	return nil
}

func (custController *CustController) CreateModel() map[string]interface{} {

	if retval := CreateModel(custController, func(m model.Model) {}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, custController.Cust)
}

func (custController *CustController) ViewModel(id int64) map[string]interface{} {
	cust := &custController.Cust
	if retval := ViewModel(id, cust); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, cust)
}

func (custController *CustController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(custController, &model.Cust{}, func(modelSrc, modelTemp model.Model) {
		custSrc := modelSrc.(*model.Cust)
		custTemp := modelTemp.(*model.Cust)

		custSrc.Name = custTemp.Name
		custSrc.Email = custTemp.Email
		custSrc.CountryCode = custTemp.CountryCode
		custSrc.PhoneNumber = custTemp.PhoneNumber
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (custController *CustController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, custController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (custController *CustController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("cust", "id ASC", &custController.Cust, make([]*model.Cust, 0), param)
}

func ComboCust(page int, param *utils.Param) map[string]interface{} {
	return GetCombo(page, "cust", "id ASC", param)
}
