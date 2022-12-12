package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Whd struct {
	WhId    int64   `json:"wh_id"`
	Dno     int     `json:"dno"`
	PurcId  int64   `json:"purc_id"`
	PurcDno int     `json:"purc_dno"`
	Qty     float32 `json:"qty" gorm:"DEFAULT:0"`
	WhdExt  `gorm:"-"`
}

type WhdExt struct {
	OrderTrxno string `json:"order_trxno"`
	PurcdWhd
}

type PurcdWhd struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
}

func (Whd) TableName() string {
	return "whd"
}

func (whd Whd) Validate() error {
	err := validation.Errors{
		"Dno":      validation.Validate(whd.Dno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Purc Id":  validation.Validate(whd.PurcId, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Purc Dno": validation.Validate(whd.PurcDno, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (whd Whd) ID() int64 {
	return -1
}

func (whd *Whd) SetMasterId(id int64) {
	whd.WhId = id
}

func (whd *Whd) SetValueModelExt(db *gorm.DB) {
	db.Select("trxno").Table("order").Where("id = (SELECT order_id FROM purcd WHERE purc_id = ? AND dno = ?)", whd.PurcId, whd.PurcDno).Scan(&whd.OrderTrxno)
	db.Select("product_id, name").Table("purcd").Where("purc_id = ? AND dno = ?", whd.PurcId, whd.PurcDno).Scan(&whd.PurcdWhd)
}
