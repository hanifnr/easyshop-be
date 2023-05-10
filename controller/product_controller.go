package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/scrape"
	"easyshop/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var CreateTrendingProduct = func(w http.ResponseWriter, r *http.Request) {
	trendingProductController := &TrendingProductController{}
	CreateModelAction(trendingProductController, w, r)
}

var UpdateTrendingProduct = func(w http.ResponseWriter, r *http.Request) {
	trendingProductController := &TrendingProductController{}
	UpdateModelAction(trendingProductController, w, r)
}

var ViewTrendingProduct = func(w http.ResponseWriter, r *http.Request) {
	trendingProductController := &TrendingProductController{}
	ViewModelAction(trendingProductController, w, r)
}

var ListTrendingProduct = func(w http.ResponseWriter, r *http.Request) {
	trendingProductController := &TrendingProductController{}
	ListModelAction(trendingProductController, w, r)
}

var DeleteTrendingProduct = func(w http.ResponseWriter, r *http.Request) {
	trendingProductController := &TrendingProductController{}
	DeleteModelAction(trendingProductController, w, r)
}

// var CleanProduct = func(w http.ResponseWriter, r *http.Request) {
// 	trendingProductController := &TrendingProductController{}
// 	resp := trendingProductController.RemoveProduct()
// 	utils.Respond(w, resp)
// }

type TrendingProductController struct {
	TrendingProduct scrape.Product
}

func (trendingProductController *TrendingProductController) Model() model.Model {
	return &trendingProductController.TrendingProduct
}

func (trendingProductController *TrendingProductController) FNew() functions.SQLFunction {
	return nil
}
func (trendingProductController *TrendingProductController) FDelete() functions.SQLFunction {
	return nil
}

func (trendingProductController *TrendingProductController) CreateModel() map[string]interface{} {

	if retval := CreateModel(trendingProductController, func(m model.Model) {
		currentTime := time.Now()

		trendingProduct := m.(*scrape.Product)
		trendingProduct.CreatedAt = currentTime
		trendingProduct.UpdatedAt = currentTime
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, trendingProductController.TrendingProduct)
}

func (trendingProductController *TrendingProductController) ViewModel(id int64) map[string]interface{} {
	trendingProduct := &scrape.Product{}
	if retval := ViewModel(id, trendingProduct); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, trendingProduct)
}

func (trendingProductController *TrendingProductController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(trendingProductController, &scrape.Product{}, func(modelSrc, modelTemp model.Model) {
		trendingProductSrc := modelSrc.(*scrape.Product)
		trendingProductTemp := modelTemp.(*scrape.Product)

		trendingProductSrc.Name = trendingProductTemp.Name
		trendingProductSrc.Code = trendingProductTemp.Code
		trendingProductSrc.Image = trendingProductTemp.Image
		trendingProductSrc.Price = trendingProductTemp.Price
		trendingProductSrc.PriceTax = trendingProductTemp.PriceTax
		trendingProductSrc.Size = trendingProductTemp.Size
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (trendingProductController *TrendingProductController) DeleteModel(id int64) map[string]interface{} {
	if retval := RemoveModel(id, trendingProductController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (trendingProductController *TrendingProductController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("product", "id ASC", &trendingProductController.TrendingProduct, make([]*scrape.Product, 0), param)
}

func CleanProduct() {
	db := utils.GetDB().Begin()

	listProduct := make([]*scrape.Product, 0)
	db.Where("EXTRACT(DAY FROM (?::date - created_at))::integer > 1 AND req_order_id IS NOT NULL", time.Now()).Find(&listProduct)

	for _, product := range listProduct {
		if err := db.Delete(&product).Error; err != nil {
			db.Rollback()
		}
	}
	b, _ := json.Marshal(utils.MessageData(true, listProduct))

	fmt.Println("AUTO CLEAN PRODUCT: \n", string(b))
	db.Commit()
}
