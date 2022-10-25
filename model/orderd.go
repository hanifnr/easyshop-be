package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Orderd struct {
	OrderId   int64   `json:"order_id"`
	Dno       int     `json:"dno"`
	ShopId    int64   `json:"shop_id"`
	Productid string  `json:"productid"`
	Name      string  `json:"name"`
	Qty       float32 `json:"qty" gorm:"DEFAULT:0"`
	Qtypurc   float32 `json:"qtypurc" gorm:"DEFAULT:0"`
	Qtywh     float32 `json:"qtywh" gorm:"DEFAULT:0"`
	Price     float64 `json:"price" gorm:"DEFAULT:0"`
	Subtotal  float64 `json:"subtotal" gorm:"DEFAULT:0"`
}

func (Orderd) TableName() string {
	return "orderd"
}

func (orderd Orderd) Validate() error {
	err := validation.Errors{
		"Dno":        validation.Validate(orderd.Dno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Shop Id":    validation.Validate(orderd.ShopId, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Product Id": validation.Validate(orderd.Productid, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Name":       validation.Validate(orderd.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Qty":        validation.Validate(orderd.Qty, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Price":      validation.Validate(orderd.Price, validation.NotNil.Error(utils.FIELD_NOTNIL)),
		"Subtotal":   validation.Validate(orderd.Subtotal, validation.NotNil.Error(utils.FIELD_NOTNIL)),
	}.Filter()

	return err
}

func (orderd Orderd) ID() int64 {
	return -1
}

func (orderd *Orderd) SetMasterId(id int64) {
	orderd.OrderId = id
}
