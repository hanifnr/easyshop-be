package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Whd struct {
	Id      int64   `json:"id"`
	WhId    int64   `json:"wh_id"`
	Dno     int     `json:"dno"`
	PurcId  int64   `json:"purc_id"`
	PurcDno int     `json:"purc_dno"`
	Qtywh   float32 `json:"qtywh" gorm:"DEFAULT:0"`
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
