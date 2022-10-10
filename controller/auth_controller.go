package controllers

import (
	u "easyshop/utils"
	"net/http"
)

var BasicTokenController = func(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		u.RespondError(w, u.Message(false, "Wrong code!"), http.StatusBadRequest)
		return
	}
	data, err := u.CreateToken(code)
	if err != nil {
		u.RespondError(w, u.Message(false, err.Error()), http.StatusBadRequest)
		return
	}
	resp := u.MessageData(true, "sucess", data)
	u.Respond(w, resp)
}
