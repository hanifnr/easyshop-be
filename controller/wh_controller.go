package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var CreateWh = func(w http.ResponseWriter, r *http.Request) {
	whController := &WhController{}
	CreateTransAction(whController, w, r)
}

var UpdateWh = func(w http.ResponseWriter, r *http.Request) {
	whController := &WhController{}
	UpdateTransAction(whController, w, r)
}

var ViewWh = func(w http.ResponseWriter, r *http.Request) {
	whController := &WhController{}
	ViewTransAction(whController, w, r)
}

var DeleteWh = func(w http.ResponseWriter, r *http.Request) {
	whController := &WhController{}
	DeleteTransAction(whController, w, r)
}

var ListWh = func(w http.ResponseWriter, r *http.Request) {
	whController := &WhController{}
	ListTransAction(whController, w, r)
}

var HandleWh = func(w http.ResponseWriter, r *http.Request) {
	type Wh struct {
		Id    int64
		Value string
	}
	wh := &Wh{}
	if err := json.NewDecoder(r.Body).Decode(&wh); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	whController := &WhController{}
	resp := whController.HandleOrder(wh.Id, wh.Value)
	utils.Respond(w, resp)
}

type WhController struct {
	Wh      model.Wh      `json:"wh"`
	Whd     []model.Whd   `json:"whd"`
	Details []model.Model `json:"-"`
}

func (whController *WhController) MasterField() string {
	return "wh_id"
}

func (whController *WhController) MasterModel() model.Model {
	return &whController.Wh
}

func (whController *WhController) DetailsModel() []model.Model {
	return whController.Details
}

func (whController *WhController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(whController, func(db *gorm.DB) error {
		whController.Wh.StatusCode = "IW"
		for i := range whController.Whd {
			whd := &whController.Whd[i]
			whController.Details = append(whController.Details, whd)
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, whController)
}

func (whController *WhController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, whController, func(db *gorm.DB) error {
		err := db.Where("wh_id = ?", id).Find(&whController.Whd).Error
		for i := range whController.Whd {
			whController.Details = append(whController.Details, &whController.Whd[i])
		}
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, whController)
}

func (whController *WhController) ListTrans(param *utils.Param) map[string]interface{} {
	return ListTrans("wh", "id ASC", make([]*model.Wh, 0), param)
}

func (whController *WhController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(whController, &model.Wh{}, &model.Whd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		whSrc := modelSrc.(*model.Wh)
		whTemp := modelTemp.(*model.Wh)
		whSrc.Trxno = whTemp.Trxno
		whSrc.Date = whTemp.Date
		whSrc.ShopId = whTemp.ShopId
		for i := range whController.Whd {
			whController.Details = append(whController.Details, &whController.Whd[i])
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, whController)

}

func (whController *WhController) DeleteTrans(id int64) map[string]interface{} {
	if retval := DeleteTrans(id, whController, func() utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (whController *WhController) FNew() functions.SQLFunction {
	return &functions.FWhNew{}
}

func (whController *WhController) FDelete() functions.SQLFunction {
	return &functions.FWhDelete{}
}

func (whController *WhController) HandleOrder(id int64, status string) map[string]interface{} {
	return UpdateFieldMaster(id, whController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		wh := m.(*model.Wh)
		wh.StatusCode = strings.ToUpper(status)
		return utils.StatusReturnOK()
	})
}
