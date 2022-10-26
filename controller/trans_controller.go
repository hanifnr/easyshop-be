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
	ListTrans(page int) map[string]interface{}
	UpdateTrans() map[string]interface{}
	FNew() functions.SQLFunction
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

func ValidateTrans(controller TransController) error {
	if len(controller.DetailsModel()) == 0 {
		return errors.New("detail transaction can't be empty")
	}
	return nil
}
