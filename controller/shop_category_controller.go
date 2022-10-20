package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"

	"gorm.io/gorm"
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

func (shopCategoryController *ShopCategoryController) CreateModel() map[string]interface{} {
	return nil
}

func (shopCategoryController *ShopCategoryController) ViewModel(id int64) map[string]interface{} {
	shopCategory := &model.ShopCategory{}
	if retval := View(id, shopCategory, shopCategoryController); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, shopCategory)
}

func (shopCategoryController *ShopCategoryController) ListModel(page int) map[string]interface{} {
	db := utils.GetDB()

	var totalRow int64
	if err := db.Select("count(id)").Table("shop").Scan(&totalRow).Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}

	listShopCategory := make([]*model.ShopCategory, 0)
	var query *gorm.DB
	if page == 0 {
		query = db.Find(&listShopCategory)
	} else {
		offset, limit := utils.GetOffsetLimit(page)
		query = db.Offset(offset).Limit(limit).Find(&listShopCategory)
	}
	if err := query.Error; err != nil {
		return utils.MessageErr(false, utils.ErrSQLList, err.Error())
	}
	respPage := utils.RespPage(page, int(totalRow))
	return utils.MessageListData(true, listShopCategory, respPage)
}
