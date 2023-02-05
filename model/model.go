package model

import (
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type SingleStringColumnUpdate struct {
	Id    int64
	Value string
}

type SingleNumericColumnUpdate struct {
	Id    int64
	Value float64
}

type Model interface {
	ID() int64
	TableName() string
	Validate() error
}

type Master interface {
	GetTrxno() string
	SetTrxno(trxno string)
}

type Detail interface {
	SetMasterId(id int64)
}

type TimeField interface {
	SetCreatedAt(time time.Time)
	SetUpdatedAt(time time.Time)
}

type DeleteField interface {
	SetIsDelete(isDelete bool)
}

type ModelExt interface {
	SetValueModelExt(db *gorm.DB)
}

func GetSingleColumnUpdate(w http.ResponseWriter, r *http.Request, model interface{}, fAction func() map[string]interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := fAction()
	utils.Respond(w, resp)
}
