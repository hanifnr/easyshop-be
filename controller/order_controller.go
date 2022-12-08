package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	CreateTransAction(orderController, w, r)
}

var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	UpdateTransAction(orderController, w, r)
}

var ViewOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	ViewTransAction(orderController, w, r)
}

var ListOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	ListTransAction(orderController, w, r)
}

var HandleOrder = func(w http.ResponseWriter, r *http.Request) {
	type Order struct {
		Id    int64
		Value string
		Note  string
	}
	order := &Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	orderController := &OrderController{}
	resp := orderController.HandleOrder(order.Id, order.Value, order.Note)
	utils.Respond(w, resp)
}

var TrackingNumber = func(w http.ResponseWriter, r *http.Request) {
	model.GetSingleColumnUpdate(w, r, func(scu *model.SingleColumnUpdate) map[string]interface{} {
		orderController := &OrderController{}
		return orderController.TrackingNumber(scu.Id, scu.Value)
	})
}

var ListOrderd = func(w http.ResponseWriter, r *http.Request) {
	param := utils.ProcessParam(r)
	orderController := &OrderController{}
	resp := orderController.ListDetail(param)
	utils.Respond(w, resp)
}

var UploadOrderProof = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	orderController.UploadOrderProof(w, r)
}

var LoadOrderProof = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	orderController.LoadOrderProof(w, r)
}

type OrderController struct {
	Order   model.Order    `json:"order"`
	Orderd  []model.Orderd `json:"orderd"`
	Details []model.Model  `json:"-"`
}

func (orderController *OrderController) MasterField() string {
	return "order_id"
}

func (orderController *OrderController) MasterModel() model.Model {
	return &orderController.Order
}

func (orderController *OrderController) DetailsModel() []model.Model {
	return orderController.Details
}

func (orderController *OrderController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(orderController, func(db *gorm.DB) error {
		order := &orderController.Order
		passport, err := utils.QueryStringValue(db.Select("number").Table("passport").Where("cust_id=?", order.CustId))
		if err != nil {
			return err
		}
		order.Trxno = "AUTO"
		order.StatusCode = "W"
		order.Passport = passport
		for i := range orderController.Orderd {
			orderd := &orderController.Orderd[i]
			orderd.Dno = i + 1
			orderController.Details = append(orderController.Details, orderd)
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderController)
}

func (orderController *OrderController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, orderController, func(db *gorm.DB) error {
		err := db.Where("order_id = ?", id).Find(&orderController.Orderd).Error
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderController)
}

func (orderController *OrderController) ListTrans(param *utils.Param) map[string]interface{} {
	return ListTrans("order", "id ASC", make([]*model.Order, 0), param)
}

func (orderController *OrderController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(orderController, &model.Order{}, &model.Orderd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		orderSrc := modelSrc.(*model.Order)
		orderTemp := modelTemp.(*model.Order)

		orderSrc.CustId = orderTemp.CustId
		orderSrc.Date = orderTemp.Date
		orderSrc.PickDate = orderTemp.PickDate
		orderSrc.Total = orderTemp.Total
		orderSrc.Trxno = orderTemp.Trxno
		orderSrc.AddrId = orderTemp.AddrId
		orderSrc.ArrivalDate = orderTemp.ArrivalDate

		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderController)

}
func (orderController *OrderController) FNew() functions.SQLFunction {
	return &functions.FOrderNew{}
}

func (orderController *OrderController) HandleOrder(id int64, status, note string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.StatusCode = strings.ToUpper(status)
		flsoNew := &functions.FOrderLogNew{
			Note: note,
		}
		if retval := flsoNew.Run(order, db); retval.ErrCode != 0 {
			return retval
		}
		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) TrackingNumber(id int64, trackingNumber string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.TrackingNumber = strings.ToUpper(trackingNumber)
		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) ListDetail(param *utils.Param) map[string]interface{} {
	imported := false
	param.Imported = &imported
	return ListJoinModel("orderd", "order_id DESC,dno ASC", make([]*model.Orderd, 0), param, func(query *gorm.DB) {
		query.Joins("JOIN \"order\" ON order_id = \"order\".id")
	}, func(query *gorm.DB) {
		query.Where("status_code IN ('PA','IP')")
	})
}

func (orderController *OrderController) UploadOrderProof(w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	file, retval := utils.GetImageFile(w, r)
	if retval.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
		return
	}
	imageUploader := utils.GetImageUploader()
	currentTime := time.Now()
	fileName := "payment-" + strconv.Itoa(int(id)) + "-" + currentTime.Format("20060102150405")
	if err := imageUploader.UploadFile(file, fileName); err != nil {
		utils.Respond(w, utils.MessageErr(false, utils.ErrIO, err.Error()))
		return
	}
	resp := UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.ProofLink = fileName

		return utils.StatusReturnOK()
	})
	utils.Respond(w, resp)
}

func (orderController *OrderController) LoadOrderProof(w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	order := &orderController.Order
	if retval := ViewModel(id, order); retval.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
		return
	}
	url, retval := utils.GenerateSignedUrl(order.ProofLink)
	if retval.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
		return
	}
	resp := make(map[string]interface{})
	resp["url"] = url
	utils.Respond(w, resp)
}
