package controllers

import (
	"easyshop/scrape"
	"easyshop/utils"
	"net/http"
)

var ListProduct = func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	resp := scrape.GetListProducts(name)
	utils.Respond(w, utils.MessageData(true, resp))
}

var ViewProduct = func(w http.ResponseWriter, r *http.Request) {
	shopId := utils.ParamToInt64(r.URL.Query().Get("shop_id"))
	url := r.URL.Query().Get("url")
	resp := scrape.GetProduct(*shopId, url)
	utils.Respond(w, resp)
}

var GetTopProduct = func(w http.ResponseWriter, r *http.Request) {
	resp := scrape.GetTopProducts()
	utils.Respond(w, utils.MessageData(true, resp))
}
