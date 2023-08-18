package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	"gorm.io/gorm"
)

const (
	REQ_ORDER_NOTIFICATION = iota
	REQ_ORDER_APPROVED_NOTIFICATION
	REQ_ORDER_REJECTED_NOTIFICATION
)

var CreateReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	CreateTransAction(reqOrderController, w, r)
}

var UpdateReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	UpdateTransAction(reqOrderController, w, r)
}

var ViewReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	ViewTransAction(reqOrderController, w, r)
}

var ListReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	ListTransAction(reqOrderController, w, r)
}

var DeleteReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	DeleteTransAction(reqOrderController, w, r)
}

var HandleReqOrder = func(w http.ResponseWriter, r *http.Request) {
	type ReqOrder struct {
		Id    int64
		Value string
	}
	req := &ReqOrder{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	reqOrderController := &ReqOrderController{}
	resp := reqOrderController.HandleReqOrder(req.Id, req.Value)
	utils.Respond(w, resp)
}

var ApproveReqOrder = func(w http.ResponseWriter, r *http.Request) {
	type ReqOrder struct {
		Id    int64
		Dno   int
		Value bool
		Note  string
	}
	req := &ReqOrder{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	reqOrderController := &ReqOrderController{}
	resp := reqOrderController.ApproveReqOrder(req.Id, req.Dno, req.Value, req.Note)
	utils.Respond(w, resp)
}

var CountWaitingReqOrder = func(w http.ResponseWriter, r *http.Request) {
	reqOrderController := &ReqOrderController{}
	resp := reqOrderController.CountWaitingReqOrder()
	utils.Respond(w, resp)
}

type ReqOrderController struct {
	ReqOrder  model.ReqOrder    `json:"req_order"`
	ReqOrderd []model.ReqOrderd `json:"req_orderd"`
	Details   []model.Model     `json:"-"`
}

func (reqOrderController *ReqOrderController) MasterField() string {
	return "req_order_id"
}

func (reqOrderController *ReqOrderController) MasterModel() model.Model {
	return &reqOrderController.ReqOrder
}

func (reqOrderController *ReqOrderController) DetailsModel() []model.Model {
	return reqOrderController.Details
}

func (reqOrderController *ReqOrderController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(reqOrderController, func(db *gorm.DB) error {
		reqOrderController.ReqOrder.StatusCode = "W"
		for i := range reqOrderController.ReqOrderd {
			reqOrder := &reqOrderController.ReqOrderd[i]
			reqOrder.Dno = i + 1
			reqOrderController.Details = append(reqOrderController.Details, &reqOrderController.ReqOrderd[i])
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	SendEmailReqOrderNotification(REQ_ORDER_NOTIFICATION, getDataNotifReqOrder(reqOrderController))
	return utils.MessageData(true, reqOrderController)
}

func (reqOrderController *ReqOrderController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, reqOrderController, func(db *gorm.DB) error {
		err := db.Where("req_order_id = ?", id).Find(&reqOrderController.ReqOrderd).Error
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, reqOrderController)
}

func (reqOrderController *ReqOrderController) ListTrans(param *utils.Param) map[string]interface{} {
	return ListTrans("req_order", "id ASC", &reqOrderController.ReqOrder, make([]*model.ReqOrder, 0), param)
}

func (reqOrderController *ReqOrderController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(reqOrderController, &model.ReqOrder{}, &model.ReqOrderd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		reqOrderSrc := modelSrc.(*model.ReqOrder)
		reqOrderTemp := modelTemp.(*model.ReqOrder)
		reqOrderSrc.Email = reqOrderTemp.Email
		for i := range reqOrderController.ReqOrderd {
			reqOrderController.Details = append(reqOrderController.Details, &reqOrderController.ReqOrderd[i])
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, reqOrderController)

}

func (reqOrderController *ReqOrderController) DeleteTrans(id int64) map[string]interface{} {
	if retval := DeleteTrans(id, reqOrderController, func() utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (reqOrderController *ReqOrderController) FNew() functions.SQLFunction {
	return nil
}

func (reqOrderController *ReqOrderController) FDelete() functions.SQLFunction {
	return nil
}

func (reqOrderController *ReqOrderController) HandleReqOrder(id int64, status string) map[string]interface{} {
	return UpdateFieldMasterWithPostSave(id, reqOrderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		req := m.(*model.ReqOrder)
		req.StatusCode = status
		return utils.StatusReturnOK()
	}, func(m model.Model) utils.StatusReturn {
		req := m.(*model.ReqOrder)
		reqOrderController.ViewTrans(req.Id)
		SendEmailReqOrderNotification(REQ_ORDER_APPROVED_NOTIFICATION, getDataNotifReqOrder(reqOrderController))
		return utils.StatusReturnOK()
	})
}

func (reqOrderController *ReqOrderController) ApproveReqOrder(id int64, dno int, value bool, note string) map[string]interface{} {
	db := utils.GetDB().Begin()
	m := &model.ReqOrderd{}
	if err := db.Where("req_order_id = ? AND dno = ?", id, dno).Find(&m).Error; err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLLoad, err.Error())
	}
	m.Approved = &value
	m.ApprovalNote = note
	if err := model.Save(m, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	db.Commit()
	reqOrderController.ViewTrans(id)
	if value {
		SendEmailReqOrderNotification(REQ_ORDER_APPROVED_NOTIFICATION, getDataNotifReqOrderd(reqOrderController, dno))
	} else {
		SendEmailReqOrderNotification(REQ_ORDER_REJECTED_NOTIFICATION, getDataNotifReqOrderd(reqOrderController, dno))
	}
	return utils.MessageData(true, m)
}

func (reqOrderController *ReqOrderController) CountWaitingReqOrder() map[string]interface{} {
	db := utils.GetDB()

	var retval int64
	db.Select("COUNT(*)").Table("req_order").Where("status_code = 'W'").Scan(&retval)

	return utils.MessageData(true, retval)
}

type NotifReqOrder struct {
	Email   string
	Reqlink string
	Trxdate string
	Note    string
}

func SendEmailReqOrderNotification(mode int, notifReqOrder *NotifReqOrder) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminEmail2 := os.Getenv("ADMIN_EMAIL2")

	runtime.GOMAXPROCS(1)
	switch mode {
	case REQ_ORDER_NOTIFICATION:
		go utils.SendEmailNotifReqOrder(adminEmail, adminEmail2, notifReqOrder.Email, notifReqOrder, notifReqOrder.Trxdate)
	case REQ_ORDER_APPROVED_NOTIFICATION:
		go utils.SendEmailNotifReqOrderApproved(adminEmail, adminEmail2, notifReqOrder.Email, notifReqOrder, notifReqOrder.Trxdate)
	case REQ_ORDER_REJECTED_NOTIFICATION:
		go utils.SendEmailNotifReqOrderRejected(adminEmail, adminEmail2, notifReqOrder.Email, notifReqOrder, notifReqOrder.Trxdate)
	}
}

func getDataNotifReqOrder(reqOrderController *ReqOrderController) *NotifReqOrder {
	reqOrder := reqOrderController.ReqOrder

	var reqLink string
	for _, reqOrderd := range reqOrderController.ReqOrderd {
		reqLink = reqLink + reqOrderd.Url + "\n"
	}

	notifReqOrder := &NotifReqOrder{
		Email:   reqOrder.Email,
		Trxdate: utils.FormatTimeToDate(reqOrder.CreatedAt),
		Reqlink: reqLink,
	}

	return notifReqOrder
}

func getDataNotifReqOrderd(reqOrderController *ReqOrderController, dno int) *NotifReqOrder {
	reqOrder := reqOrderController.ReqOrder

	var reqLink, note string
	for _, reqOrderd := range reqOrderController.ReqOrderd {
		if reqOrderd.Dno == dno {
			reqLink = reqOrderd.Url
			note = reqOrderd.ApprovalNote
			break
		}
	}

	notifReqOrder := &NotifReqOrder{
		Email:   reqOrder.Email,
		Trxdate: utils.FormatTimeToDate(reqOrder.CreatedAt),
		Reqlink: reqLink,
		Note:    note,
	}

	return notifReqOrder
}
