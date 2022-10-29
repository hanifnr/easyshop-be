package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
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
	type OrderStatus struct {
		Id     int64
		Status string
	}
	orderStatus := &OrderStatus{}
	if err := json.NewDecoder(r.Body).Decode(&orderStatus); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	orderController := &OrderController{}
	resp := orderController.HandleOrder(orderStatus.Id, orderStatus.Status)
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
	if retval := CreateTrans(orderController, func(m model.Model) {
		currentTime := time.Now()

		order := m.(*model.Order)
		order.Status = "W"
		order.CreatedAt = currentTime
		order.UpdatedAt = currentTime

		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
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

func (orderController *OrderController) ListTrans(page int) map[string]interface{} {
	return ListTrans(page, "order", make([]*model.Order, 0))
}

func (orderController *OrderController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(orderController, &model.Order{}, &model.Orderd{}, func(modelSrc, modelTemp model.Model) {
		orderSrc := modelSrc.(*model.Order)
		orderTemp := modelTemp.(*model.Order)

		orderSrc.Cust_id = orderTemp.Cust_id
		orderSrc.Date = orderTemp.Date
		orderSrc.Pick_date = orderTemp.Pick_date
		orderSrc.Total = orderTemp.Total
		orderSrc.Trxno = orderTemp.Trxno

		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
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
	db := utils.GetDB().Begin()
	order := &model.Order{}
	if retval := ViewModel(id, order); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	order.Status = strings.ToUpper(status)
	if err := model.Save(order, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	db.Commit()
	return utils.MessageData(true, order)
}
