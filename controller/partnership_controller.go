package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	PARTNERSHIP_REQUEST = iota
	PARTNERSHIP_APPROVED
	PARTNERSHIP_REFERRAL
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
		Value string
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
	notifPartnership := getDataNotifPartnership(&partnershipController.Partnership, "")
	SendEmailPartnership(PARTNERSHIP_REQUEST, &partnershipController.Partnership, *notifPartnership)
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
		partnershipSrc.Email = partnershipTemp.Email
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

func (partnershipController *PartnershipController) ApprovePartnership(id int64, approvalStatus string, note string) map[string]interface{} {
	return UpdateFieldModelWithPostSave(id, partnershipController, func(m model.Model) utils.StatusReturn {
		partnership := m.(*model.Partnership)
		partnership.ApprovalStatus = approvalStatus
		partnership.Note = note
		return utils.StatusReturnOK()
	}, func() utils.StatusReturn {
		partnership := partnershipController.Partnership
		if partnership.ApprovalStatus == "A" {
			notifPartnership := getDataNotifPartnership(&partnership, "")
			SendEmailPartnership(PARTNERSHIP_APPROVED, &partnership, *notifPartnership)
		}
		return utils.StatusReturnOK()
	})
}

func SendEmailPartnership(mode int, partnership *model.Partnership, notifPartnership NotifPartnership) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminEmail2 := os.Getenv("ADMIN_EMAIL2")
	runtime.GOMAXPROCS(1)
	switch mode {
	case PARTNERSHIP_REQUEST:
		go utils.SendEmailNotifPartnershipRequest(adminEmail, adminEmail2, partnership.Email, notifPartnership, notifPartnership.Trxdate)
	case PARTNERSHIP_APPROVED:
		go utils.SendEmailNotifPartnershipApproved(adminEmail, adminEmail2, partnership.Email, notifPartnership, notifPartnership.Trxdate)
	case PARTNERSHIP_REFERRAL:
		go utils.SendEmailNotifPartnershipReferral(adminEmail, adminEmail2, partnership.Email, notifPartnership)
	}

}

type NotifPartnership struct {
	Name    string
	Code    string
	Trxdate string
	Social  string
	Phone   string
	Email   string
}

func getDataNotifPartnership(partnership *model.Partnership, code string) *NotifPartnership {
	notifPartnership := &NotifPartnership{
		Name:    partnership.Name,
		Trxdate: utils.FormatTimeToDate(partnership.CreatedAt),
		Social:  partnership.SocialMedia,
		Phone:   partnership.PhoneNumber,
		Email:   partnership.Email,
		Code:    code,
	}

	return notifPartnership
}
