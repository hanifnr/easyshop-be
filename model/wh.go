package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Wh struct {
	Id        int64     `json:"id"`
	Trxno     string    `json:"trxno"`
	Date      time.Time `json:"date"`
	ShopId    int64     `json:"shop_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete  bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
}

func (Wh) TableName() string {
	return "wh"
}

func (wh Wh) ID() int64 {
	return wh.Id
}

func (wh Wh) Validate() error {
	err := validation.Errors{
		"Trxno":   validation.Validate(wh.Trxno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Date":    validation.Validate(wh.Date, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Shop Id": validation.Validate(wh.ShopId, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (wh *Wh) GetTrxno() string {
	return wh.Trxno
}

func (wh *Wh) SetTrxno(trxno string) {
	wh.Trxno = trxno
}

func (wh *Wh) SetCreatedAt(time time.Time) {
	wh.CreatedAt = time
}

func (wh *Wh) SetUpdatedAt(time time.Time) {
	wh.UpdatedAt = time
}
