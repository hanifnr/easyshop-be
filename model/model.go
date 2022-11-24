package model

import (
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type SingleColumnUpdate struct {
	Id    int64
	Value string
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

func GetSingleColumnUpdate(w http.ResponseWriter, r *http.Request, fAction func(scu *SingleColumnUpdate) map[string]interface{}) {
	scu := &SingleColumnUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&scu); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	resp := fAction(scu)
	utils.Respond(w, resp)
}
