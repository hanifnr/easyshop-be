package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var ViewShopCategory = func(w http.ResponseWriter, r *http.Request) {
	shopCategoryController := &ShopCategoryController{}
	ViewModelAction(shopCategoryController, w, r)
}

var ListShopCategory = func(w http.ResponseWriter, r *http.Request) {
	shopCategoryController := &ShopCategoryController{}
	ListModelAction(shopCategoryController, w, r)
}

type ShopCategoryController struct {
	ShopCategory model.ShopCategory
}

func (shopCategoryController *ShopCategoryController) Model() model.Model {
	return &shopCategoryController.ShopCategory
}

func (shopCategoryController *ShopCategoryController) FNew() functions.SQLFunction {
	return nil
}

func (shopCategoryController *ShopCategoryController) FDelete() functions.SQLFunction {
	return nil
}

func (shopCategoryController *ShopCategoryController) CreateModel() map[string]interface{} {
	return nil
}

func (shopCategoryController *ShopCategoryController) ViewModel(id int64) map[string]interface{} {
	shopCategory := &shopCategoryController.ShopCategory
	if retval := ViewModel(id, shopCategory); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, shopCategory)
}

func (shopCategoryController *ShopCategoryController) DeleteModel(id int64) map[string]interface{} {
	return nil
}

func (shopCategoryController *ShopCategoryController) UpdateModel() map[string]interface{} {
	return nil
}

func (shopCategoryController *ShopCategoryController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("shop_category", "id ASC", make([]*model.ShopCategory, 0), param)
}
