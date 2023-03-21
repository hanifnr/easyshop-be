package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var ListNotification = func(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	resp := model.GetNotification(uid)
	utils.Respond(w, resp)
}
