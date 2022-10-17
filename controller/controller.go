package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
)

func CreateModelAction(model model.Model, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		data := utils.Message(false, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := model.CreateModel()
	utils.Respond(w, resp)
}
