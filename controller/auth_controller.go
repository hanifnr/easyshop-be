package controllers

import (
	u "easyshop/utils"
	"net/http"
)

var BasicTokenController = func(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		u.RespondError(w, u.MessageErr(false, 0000, "Wrong code!"), http.StatusBadRequest)
		return
	}
	data, err := u.CreateToken(code)
	if err != nil {
		u.RespondError(w, u.MessageErr(false, 0000, err.Error()), http.StatusBadRequest)
		return
	}
	resp := u.MessageData(true, data)
	u.Respond(w, resp)
}
