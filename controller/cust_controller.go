package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"strings"
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

var ListCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	ListModelAction(custController, w, r)
}

var HandleCust = func(w http.ResponseWriter, r *http.Request) {
	type CustStatus struct {
		Id     int64
		Status string
	}
	custStatus := &CustStatus{}
	if err := json.NewDecoder(r.Body).Decode(&custStatus); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	custController := &CustController{}
	resp := custController.HandleCust(custStatus.Id, custStatus.Status)
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

func (custController *CustController) CreateModel() map[string]interface{} {

	if retval := CreateModel(custController, func(m model.Model) {
		cust := m.(*model.Cust)
		cust.Status = "W"
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, custController.Cust)
}

func (custController *CustController) ViewModel(id int64) map[string]interface{} {
	cust := &model.Cust{}
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

func (custController *CustController) ListModel(page int) map[string]interface{} {
	return ListModel(page, "cust", make([]*model.Cust, 0))
}

func (custController *CustController) HandleCust(id int64, status string) map[string]interface{} {
	db := utils.GetDB().Begin()
	cust := &model.Cust{}
	if retval := ViewModel(id, cust); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	cust.Status = strings.ToUpper(status)
	if err := model.Save(cust, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	db.Commit()
	return utils.MessageData(true, cust)
}
