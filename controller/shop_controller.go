package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var ViewShop = func(w http.ResponseWriter, r *http.Request) {
	shopController := &ShopController{}
	ViewModelAction(shopController, w, r)
}

var ListShop = func(w http.ResponseWriter, r *http.Request) {
	shopController := &ShopController{}
	ListModelAction(shopController, w, r)
}

type ShopController struct {
	Shop model.Shop
}

func (shopController *ShopController) Model() model.Model {
	return &shopController.Shop
}

func (shopController *ShopController) FNew() functions.SQLFunction {
	return nil
}

func (shopController *ShopController) CreateModel() map[string]interface{} {
	return nil
}

func (shopController *ShopController) ViewModel(id int64) map[string]interface{} {
	shop := &model.Shop{}
	if retval := ViewModel(id, shop); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, shop)
}

func (shopController *ShopController) UpdateModel() map[string]interface{} {
	return nil
}

func (shopController *ShopController) ListModel(page int) map[string]interface{} {
	return ListModel(page, "shop", make([]*model.Shop, 0))
}
