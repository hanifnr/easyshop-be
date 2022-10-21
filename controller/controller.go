package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller interface {
	Model() model.Model
	CreateModel() map[string]interface{}
	ViewModel(id int64) map[string]interface{}
	ListModel(page int) map[string]interface{}
	UpdateModel() map[string]interface{}
	FNew() functions.SQLFunction
}

func CreateModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(controller.Model()); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.CreateModel()
	utils.Respond(w, resp)
}

func ViewModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.ViewModel(int64(id))
	utils.Respond(w, resp)
}

func ListModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	paramPage := r.URL.Query().Get("page")
	if paramPage == "" {
		paramPage = "0"
	}
	page, err := strconv.Atoi(paramPage)
	if err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.ListModel(page)
	utils.Respond(w, resp)
}

func UpdateModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(controller.Model()); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.UpdateModel()
	utils.Respond(w, resp)
}
