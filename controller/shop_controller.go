package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"strconv"
)

var ViewShop = func(w http.ResponseWriter, r *http.Request) {
	shopController := &ShopController{}
	ViewModelAction(shopController, w, r)
}

var ListShop = func(w http.ResponseWriter, r *http.Request) {
	shopController := &ShopController{}
	ListModelAction(shopController, w, r)
}

var ListComboShop = func(w http.ResponseWriter, r *http.Request) {
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
	resp := ComboShop(page, param)
	utils.Respond(w, resp)
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
	shop := &shopController.Shop
	if retval := ViewModel(id, shop); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, shop)
}

func (shopController *ShopController) UpdateModel() map[string]interface{} {
	return nil
}

func (shopController *ShopController) ListModel(page int, param *utils.Param) map[string]interface{} {
	return ListModel(page, "shop", make([]*model.Shop, 0), param)
}

func ComboShop(page int, param *utils.Param) map[string]interface{} {
	return GetCombo(page, "shop", param)
}
