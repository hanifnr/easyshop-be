package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Purcd struct {
	PurcId    int64   `json:"purc_id"`
	Dno       int     `json:"dno"`
	ProductId string  `json:"product_id"`
	Name      string  `json:"name"`
	OrderId   int64   `json:"order_id"`
	OrderDno  int     `json:"order_dno"`
	Qty       float32 `json:"qty" gorm:"DEFAULT:0"`
	Qtywh     float32 `json:"qtywh" gorm:"DEFAULT:0"`
	Price     float64 `json:"price" gorm:"DEFAULT:0"`
	Subtotal  float64 `json:"subtotal" gorm:"DEFAULT:0"`
	Imported  bool    `json:"imported" gorm:"DEFAULT:FALSE"`
}

func (Purcd) TableName() string {
	return "purcd"
}

func (purcd Purcd) Validate() error {
	err := validation.Errors{
		"Dno":       validation.Validate(purcd.Dno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Order Id":  validation.Validate(purcd.OrderId, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Order Dno": validation.Validate(purcd.OrderDno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Qty":       validation.Validate(purcd.Qty, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Price":     validation.Validate(purcd.Price, validation.NotNil.Error(utils.FIELD_NOTNIL)),
		"Subtotal":  validation.Validate(purcd.Subtotal, validation.NotNil.Error(utils.FIELD_NOTNIL)),
	}.Filter()

	return err
}

func (purcd Purcd) ID() int64 {
	return -1
}

func (purcd *Purcd) SetMasterId(id int64) {
	purcd.PurcId = id
}
