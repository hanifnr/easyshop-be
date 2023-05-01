package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"

	"gorm.io/gorm"
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
		for i := range reqOrderController.ReqOrderd {
			reqOrder := &reqOrderController.ReqOrderd[i]
			reqOrder.Dno = i + 1
			reqOrderController.Details = append(reqOrderController.Details, &reqOrderController.ReqOrderd[i])
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
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
