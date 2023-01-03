package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Wh struct {
	Id         int64     `json:"id"`
	Trxno      string    `json:"trxno"`
	Date       time.Time `json:"date"`
	ShopId     int64     `json:"shop_id"`
	StatusCode string    `json:"status_code"`
	CreatedAt  time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete   bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
	IsActive   bool      `json:"is_active" gorm:"DEFAULT:TRUE"`
	WhExt      `gorm:"-"`
}

type WhExt struct {
	ShopName string `json:"shop_name"`
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

func (wh *Wh) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("shop").Where("id = ?", wh.ShopId).Scan(&wh.ShopName)
}

func (wh *Wh) SetIsDelete(isDelete bool) {
	wh.IsDelete = isDelete
}
