package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	CreateModelAction(custController, w, r)
}

var ViewCust = func(w http.ResponseWriter, r *http.Request) {
	custController := &CustController{}
	ViewModelAction(custController, w, r)
}

var ListCust = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, GetListCust())
}

type CustController struct {
	Cust model.Cust
}

func (cust *CustController) Model() model.Model {
	return &cust.Cust
}

func (cust *CustController) FNew() functions.SQLFunction {
	return &functions.FCustNew{}
}

func (custController *CustController) CreateModel() map[string]interface{} {
	currentTime := time.Now()
	cust := &custController.Cust
	cust.Status = "W"
	cust.CreatedAt = currentTime
	cust.UpdatedAt = currentTime
	if retval := Save(custController); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, "Saved succesfully!", custController.Cust)
}

func (custController *CustController) ViewModel(id int64) map[string]interface{} {
	cust := &model.Cust{}
	if retval := View(id, cust, custController); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, "View data", cust)
}

func GetListCust() map[string]interface{} {
	db := utils.GetDB()
	listCust := make([]*model.Cust, 0)
	if err := db.Find(&listCust).Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	return utils.MessageData(true, "List customer", listCust)
}
