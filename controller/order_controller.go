package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"strings"

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
	model.GetSingleColumnUpdate(w, r, func(scu *model.SingleColumnUpdate) map[string]interface{} {
		orderController := &OrderController{}
		return orderController.HandleOrder(scu.Id, scu.Value)
	})
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
		order.Status = "W"
		order.Passport = passport
		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
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
	return nil
}

func (orderController *OrderController) HandleOrder(id int64, status string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model) {
		order := m.(*model.Order)
		order.Status = strings.ToUpper(status)
	})
}

func (orderController *OrderController) TrackingNumber(id int64, trackingNumber string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model) {
		order := m.(*model.Order)
		order.TrackingNumber = strings.ToUpper(trackingNumber)
	})
}

func (orderController *OrderController) ListDetail(param *utils.Param) map[string]interface{} {
	return ListModel("orderd", "order_id DESC,dno ASC", make([]*model.Orderd, 0), param)
}
