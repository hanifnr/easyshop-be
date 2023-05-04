package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ReqOrder struct {
	Id         int64     `json:"id" gorm:"primary_key;auto_increment"`
	Email      string    `json:"email"`
	StatusCode string    `json:"status_code"`
	CreatedAt  time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
}

func (reqOrder ReqOrder) ID() int64 {
	return reqOrder.Id
}

func (ReqOrder) TableName() string {
	return "req_order"
}

func (reqOrder *ReqOrder) GetTrxno() string {
	return ""
}

func (reqOrder *ReqOrder) SetTrxno(trxno string) {
}

func (reqOrder ReqOrder) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(reqOrder.Email, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (reqOrder *ReqOrder) SetCreatedAt(time time.Time) {
	reqOrder.CreatedAt = time
}

func (reqOrder *ReqOrder) SetUpdatedAt(time time.Time) {
}
