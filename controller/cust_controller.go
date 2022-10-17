package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var CreateCust = func(w http.ResponseWriter, r *http.Request) {
	cust := &model.Cust{}
	CreateModelAction(cust, w, r)
}

var ListCust = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, model.GetListCust())
}
