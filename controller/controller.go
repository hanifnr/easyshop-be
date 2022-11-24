package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller interface {
	Model() model.Model
	CreateModel() map[string]interface{}
	ViewModel(id int64) map[string]interface{}
	ListModel(param *utils.Param) map[string]interface{}
	UpdateModel() map[string]interface{}
	DeleteModel(id int64) map[string]interface{}
	FNew() functions.SQLFunction
	FDelete() functions.SQLFunction
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
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	resp := controller.ViewModel(int64(id))
	utils.Respond(w, resp)
}

func ListModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	param := utils.ProcessParam(r)
	resp := controller.ListModel(param)
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

func DeleteModelAction(controller Controller, w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	resp := controller.DeleteModel(int64(id))
	utils.Respond(w, resp)
}

func GetInt64Param(param string, w http.ResponseWriter, r *http.Request) (int64, error) {
	params := mux.Vars(r)
	value, err := strconv.ParseInt(params[param], 10, 64)
	if err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return -1, err
	}
	return value, err
}

func GetStringQuery(param string, isRequired bool, w http.ResponseWriter, r *http.Request) (string, error) {
	return GetStringQueryCheck(param, isRequired, func(result string) error { return nil }, w, r)
}

func GetStringQueryCheck(param string, isRequired bool, fCheck func(result string) error, w http.ResponseWriter, r *http.Request) (string, error) {
	result := r.URL.Query().Get(param)
	if isRequired && result == "" {
		err := fmt.Errorf("param %s required", param)
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return "", err
	}
	if err := fCheck(result); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return "", err
	}
	return result, nil
}
