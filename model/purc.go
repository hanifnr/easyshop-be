package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Purc struct {
	Id        int64     `json:"id"`
	Trxno     string    `json:"trxno"`
	Date      time.Time `json:"date"`
	ShopId    int64     `json:"shop_id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"  gorm:"DEFAULT:0"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	Isdelete  bool      `json:"isdelete" gorm:"DEFAULT:FALSE"`
	PurcExt   `gorm:"-"`
}

type PurcExt struct {
	ShopName string `json:"shop_name"`
}

func (Purc) TableName() string {
	return "purc"
}

func (purc Purc) Validate() error {
	err := validation.Errors{
		"Trxno":   validation.Validate(purc.Trxno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Date":    validation.Validate(purc.Date, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Total":   validation.Validate(purc.Total, validation.Required.Error(utils.FIELD_NOTNIL)),
		"Shop Id": validation.Validate(purc.ShopId, validation.Required.Error(utils.FIELD_NOTNIL)),
	}.Filter()

	return err
}

func (purc Purc) ID() int64 {
	return purc.Id
}

func (purc *Purc) GetTrxno() string {
	return purc.Trxno
}

func (purc *Purc) SetTrxno(trxno string) {
	purc.Trxno = trxno
}

func (purc *Purc) SetCreatedAt(time time.Time) {
	purc.CreatedAt = time
}

func (purc *Purc) SetUpdatedAt(time time.Time) {
	purc.UpdatedAt = time
}

func (purc *Purc) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("shop").Where("id = ?", purc.ShopId).Scan(&purc.ShopName)
}
