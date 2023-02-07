package controllers

import (
	"easyshop/etc"
	"easyshop/utils"
	"net/http"
)

var GetProps = func(w http.ResponseWriter, r *http.Request) {
	resp := etc.ReadProps()
	utils.Respond(w, utils.MessageData(true, resp))
}
