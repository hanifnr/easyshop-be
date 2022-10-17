package controllers

import (
	"easyshop/model"
	"net/http"
)

var CreateCust = func(w http.ResponseWriter, r *http.Request) {
	cust := &model.Cust{}
	CreateModelAction(cust, w, r)
}
