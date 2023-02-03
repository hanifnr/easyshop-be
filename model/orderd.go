package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Orderd struct {
	OrderId   int64   `json:"order_id"`
	Dno       int     `json:"dno"`
	ShopId    int64   `json:"shop_id"`
	ProductId string  `json:"product_id"`
	Name      string  `json:"name"`
	Qty       float32 `json:"qty" gorm:"DEFAULT:0"`
	Qtypurc   float32 `json:"qtypurc" gorm:"DEFAULT:0"`
	Qtywh     float32 `json:"qtywh" gorm:"DEFAULT:0"`
	Price     float64 `json:"price" gorm:"DEFAULT:0"`
	Subtotal  float64 `json:"subtotal" gorm:"DEFAULT:0"`
	Url       string  `json:"url"`
	Image     string  `json:"image"`
	Imported  bool    `json:"imported" gorm:"DEFAULT:FALSE"`
	Arrived   bool    `json:"arrived" gorm:"DEFAULT:FALSE"`
	Note      string  `json:"note" gorm:"DEFAULT:FALSE"`
	OrderdExt `gorm:"-"`
}

type OrderdExt struct {
	OrderTrxno string `json:"order_trxno"`
}

func (Orderd) TableName() string {
	return "orderd"
}

func (orderd Orderd) Validate() error {
	err := validation.Errors{
		"Shop Id":    validation.Validate(orderd.ShopId, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Product Id": validation.Validate(orderd.ProductId, validation.Required.Error(utils.FIELD_REQUIRED)),
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

func (orderd *Orderd) SetValueModelExt(db *gorm.DB) {
	db.Select("trxno").Table("order").Where("id = ?", orderd.OrderId).Scan(&orderd.OrderTrxno)
}
