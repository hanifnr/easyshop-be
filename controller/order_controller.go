package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	CreateTransAction(orderController, w, r)
}

type OrderController struct {
	Order   model.Order    `json:"order"`
	Orderd  []model.Orderd `json:"detail"`
	Details []model.Model  `json:"-"`
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
	return nil
}
func (orderController *OrderController) ListTrans(page int) map[string]interface{} {
	return nil
}
func (orderController *OrderController) UpdateTrans() map[string]interface{} {
	return nil
}
func (orderController *OrderController) FNew() functions.SQLFunction {
	return nil
}
