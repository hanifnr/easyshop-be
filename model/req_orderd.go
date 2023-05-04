package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ReqOrderd struct {
	ReqOrderId   int64  `json:"req_order_id" gorm:"primaryKey"`
	Dno          int    `json:"dno" gorm:"primaryKey"`
	Url          string `json:"url"`
	Approved     bool   `json:"approved"`
	ApprovalNote string `json:"approval_note"`
	Note         string `json:"note"`
}

func (reqOrderd ReqOrderd) ID() int64 {
	return -1
}

func (ReqOrderd) TableName() string {
	return "req_orderd"
}

func (reqOrderd ReqOrderd) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(reqOrderd.Url, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (reqOrderd *ReqOrderd) SetMasterId(id int64) {
	reqOrderd.ReqOrderId = id
}
