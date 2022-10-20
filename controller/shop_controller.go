package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"

	"gorm.io/gorm"
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
	if retval := View(id, shop, shopController); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, shop)
}

func (shopController *ShopController) ListModel(page int) map[string]interface{} {
	db := utils.GetDB()

	var totalRow int64
	if err := db.Select("count(id)").Table("shop").Scan(&totalRow).Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}

	listShop := make([]*model.Shop, 0)
	var query *gorm.DB
	if page == 0 {
		query = db.Find(&listShop)
	} else {
		offset, limit := utils.GetOffsetLimit(page)
		query = db.Offset(offset).Limit(limit).Find(&listShop)
	}
	if err := query.Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	respPage := utils.RespPage(page, int(totalRow))
	return utils.MessageListData(true, listShop, respPage)
}
