package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type TransController interface {
	MasterModel() model.Model
	DetailsModel() []model.Model
	CreateTrans() map[string]interface{}
	ViewTrans(id int64) map[string]interface{}
	ListTrans(param *utils.Param) map[string]interface{}
	UpdateTrans() map[string]interface{}
	DeleteTrans(id int64) map[string]interface{}
	FNew() functions.SQLFunction
	FDelete() functions.SQLFunction
	MasterField() string
}

func CreateTransAction(controller TransController, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(controller); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.CreateTrans()
	utils.Respond(w, resp)
}

func ViewTransAction(controller TransController, w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	resp := controller.ViewTrans(int64(id))
	utils.Respond(w, resp)
}

func ListTransAction(controller TransController, w http.ResponseWriter, r *http.Request) {
	param := utils.ProcessParam(r)
	resp := controller.ListTrans(param)
	utils.Respond(w, resp)
}

func UpdateTransAction(controller TransController, w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(controller); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := controller.UpdateTrans()
	utils.Respond(w, resp)
}

func DeleteTransAction(controller TransController, w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	resp := controller.DeleteTrans(int64(id))
	utils.Respond(w, resp)
}

func ValidateTrans(controller TransController) error {
	if len(controller.DetailsModel()) == 0 {
		return errors.New("detail transaction can't be empty")
	}
	return nil
}
