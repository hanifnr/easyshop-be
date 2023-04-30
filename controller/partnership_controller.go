package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var CreatePartnership = func(w http.ResponseWriter, r *http.Request) {
	partnershipController := &PartnershipController{}
	CreateModelAction(partnershipController, w, r)
}

var UpdatePartnership = func(w http.ResponseWriter, r *http.Request) {
	partnershipController := &PartnershipController{}
	UpdateModelAction(partnershipController, w, r)
}

var ViewPartnership = func(w http.ResponseWriter, r *http.Request) {
	partnershipController := &PartnershipController{}
	ViewModelAction(partnershipController, w, r)
}

var ListPartnership = func(w http.ResponseWriter, r *http.Request) {
	partnershipController := &PartnershipController{}
	ListModelAction(partnershipController, w, r)
}

var DeletePartnership = func(w http.ResponseWriter, r *http.Request) {
	partnershipController := &PartnershipController{}
	DeleteModelAction(partnershipController, w, r)
}

var ListComboPartnership = func(w http.ResponseWriter, r *http.Request) {
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
	param := utils.ProcessParam(r)
	resp := ComboPartnership(page, param)
	utils.Respond(w, resp)
}

var ListComboPartnershipType = func(w http.ResponseWriter, r *http.Request) {
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
	param := utils.ProcessParam(r)
	resp := ComboPartnershipType(page, param)
	utils.Respond(w, resp)
}

var ApprovePartnership = func(w http.ResponseWriter, r *http.Request) {
	type Approval struct {
		Id    int64
		Value bool
		Note  string
	}
	approval := &Approval{}
	if err := json.NewDecoder(r.Body).Decode(&approval); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	partnershipController := &PartnershipController{}
	resp := partnershipController.ApprovePartnership(approval.Id, approval.Value, approval.Note)
	utils.Respond(w, resp)
}

type PartnershipController struct {
	Partnership model.Partnership
}

func (partnershipController *PartnershipController) Model() model.Model {
	return &partnershipController.Partnership
}

func (partnershipController *PartnershipController) FNew() functions.SQLFunction {
	return nil
}
func (partnershipController *PartnershipController) FDelete() functions.SQLFunction {
	return nil
}

func (partnershipController *PartnershipController) CreateModel() map[string]interface{} {

	if retval := CreateModel(partnershipController, func(m model.Model) {
		currentTime := time.Now()

		partnership := m.(*model.Partnership)
		partnership.CreatedAt = currentTime
		partnership.UpdatedAt = currentTime
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, partnershipController.Partnership)
}

func (partnershipController *PartnershipController) ViewModel(id int64) map[string]interface{} {
	partnership := &model.Partnership{}
	if retval := ViewModel(id, partnership); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, partnership)
}

func (partnershipController *PartnershipController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(partnershipController, &model.Partnership{}, func(modelSrc, modelTemp model.Model) {
		partnershipSrc := modelSrc.(*model.Partnership)
		partnershipTemp := modelTemp.(*model.Partnership)

		partnershipSrc.Name = partnershipTemp.Name
		partnershipSrc.PartnershipTypeId = partnershipTemp.PartnershipTypeId
		partnershipSrc.SocialMedia = partnershipTemp.SocialMedia
		partnershipSrc.PhoneNumber = partnershipTemp.PhoneNumber
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (partnershipController *PartnershipController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, partnershipController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (partnershipController *PartnershipController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("partnership", "id ASC", &partnershipController.Partnership, make([]*model.Partnership, 0), param)
}

func ComboPartnership(page int, param *utils.Param) map[string]interface{} {
	return GetCombo(page, "partnership", "name ASC", param)
}

func ComboPartnershipType(page int, param *utils.Param) map[string]interface{} {
	return GetCombo(page, "partnership_type", "name ASC", param)
}

func (partnershipController *PartnershipController) ApprovePartnership(id int64, approved bool, note string) map[string]interface{} {
	return UpdateFieldModel(id, partnershipController, func(m model.Model) utils.StatusReturn {
		partnership := m.(*model.Partnership)
		partnership.Approved = &approved
		partnership.Note = note
		return utils.StatusReturnOK()
	})
}
