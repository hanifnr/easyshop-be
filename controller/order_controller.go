package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"net/http"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	CreateTransAction(orderController, w, r)
}

type OrderController struct {
	model.Order
	Orderd  []model.Orderd `json:"detail"`
	Details []model.Model  `json:"-"`
}

func (orderController *OrderController) MasterModel() model.Model {
	return orderController.Order
}

func (orderController *OrderController) DetailsModel() []model.Model {
	return orderController.Details
}

func (orderController *OrderController) CreateTrans() map[string]interface{} {
	return nil
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
