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

var CreatePurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	CreateTransAction(purcController, w, r)
}

var UpdatePurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	UpdateTransAction(purcController, w, r)
}

var ViewPurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	ViewTransAction(purcController, w, r)
}

var ListPurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	ListTransAction(purcController, w, r)
}

var HandlePurc = func(w http.ResponseWriter, r *http.Request) {
	type PurcStatus struct {
		Id     int64
		Status string
	}
	purcStatus := &PurcStatus{}
	if err := json.NewDecoder(r.Body).Decode(&purcStatus); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	purcController := &PurcController{}
	resp := purcController.HandlePurc(purcStatus.Id, purcStatus.Status)
	utils.Respond(w, resp)
}

type PurcController struct {
	Purc    model.Purc    `json:"purc"`
	Purcd   []model.Purcd `json:"purcd"`
	Details []model.Model `json:"-"`
}

func (purcController *PurcController) MasterField() string {
	return "purc_id"
}

func (purcController *PurcController) MasterModel() model.Model {
	return &purcController.Purc
}

func (purcController *PurcController) DetailsModel() []model.Model {
	return purcController.Details
}

func (purcController *PurcController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(purcController, func(db *gorm.DB) error {
		purc := &purcController.Purc
		purc.Status = "P"
		for i := range purcController.Purcd {
			purcd := &purcController.Purcd[i]
			type OrderProduct struct {
				Productid string
				Name      string
			}
			orderProduct := &OrderProduct{}
			if err := db.Debug().Select("productid,name").Table("orderd").Where("order_id=? AND dno=?", purcd.OrderId, purcd.Dno).Scan(&orderProduct).Error; err != nil {
				return err
			}
			purcd.Productid = orderProduct.Productid
			purcd.Name = orderProduct.Name
			purcController.Details = append(purcController.Details, &purcController.Purcd[i])
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)
}

func (purcController *PurcController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, purcController, func(db *gorm.DB) error {
		err := db.Where("purc_id = ?", id).Find(&purcController.Purcd).Error
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)
}

func (purcController *PurcController) ListTrans(page int, param *Param) map[string]interface{} {
	return ListTrans(page, "purc", make([]*model.Purc, 0), param)
}

func (purcController *PurcController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(purcController, &model.Purc{}, &model.Purcd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		purcSrc := modelSrc.(*model.Purc)
		purcTemp := modelTemp.(*model.Purc)
		purcSrc.Trxno = purcTemp.Trxno
		purcSrc.Date = purcTemp.Date
		purcSrc.ShopId = purcTemp.ShopId
		purcSrc.Total = purcTemp.Total
		for i := range purcController.Purcd {
			purcd := &purcController.Purcd[i]
			type OrderProduct struct {
				productid string
				name      string
			}
			orderProduct := &OrderProduct{}
			if err := db.Select("productid,name").Table("orderd").Where("order_id=? AND dno=?", purcd.OrderId, purcd.Dno).Scan(&orderProduct).Error; err != nil {
				return err
			}
			purcd.Productid = orderProduct.productid
			purcd.Name = orderProduct.name
			purcController.Details = append(purcController.Details, purcd)
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)

}
func (purcController *PurcController) FNew() functions.SQLFunction {
	return nil
}

func (purcController *PurcController) HandlePurc(id int64, status string) map[string]interface{} {
	db := utils.GetDB().Begin()
	purc := &model.Purc{}
	if retval := ViewModel(id, purc); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	purc.Status = strings.ToUpper(status)
	if err := model.Save(purc, db); err != nil {
		db.Rollback()
		return utils.MessageErr(false, utils.ErrSQLSave, err.Error())
	}
	db.Commit()
	return utils.MessageData(true, purc)
}
