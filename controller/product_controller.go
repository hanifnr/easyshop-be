package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/scrape"
	"easyshop/utils"
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
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (trendingProductController *TrendingProductController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, trendingProductController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (trendingProductController *TrendingProductController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("trendingProduct", "id ASC", &trendingProductController.TrendingProduct, make([]*scrape.Product, 0), param)
}
