package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"

	"gorm.io/gorm"
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
	currentTime := time.Now()

	if retval := CreateModel(custController, func(m model.Model) {
		cust := m.(*model.Cust)
		cust.Status = "W"
		cust.CreatedAt = currentTime
		cust.UpdatedAt = currentTime
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
		custSrc.Passport = custTemp.Passport
		custSrc.UpdatedAt = time.Now()
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (custController *CustController) ListModel(page int) map[string]interface{} {
	db := utils.GetDB()

	var totalRow int64
	if err := db.Select("count(id)").Table("cust").Scan(&totalRow).Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}

	listCust := make([]*model.Cust, 0)
	var query *gorm.DB
	if page == 0 {
		query = db.Find(&listCust)
	} else {
		offset, limit := utils.GetOffsetLimit(page)
		query = db.Offset(offset).Limit(limit).Find(&listCust)
	}
	if err := query.Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	respPage := utils.RespPage(page, int(totalRow))
	return utils.MessageListData(true, listCust, respPage)
}
