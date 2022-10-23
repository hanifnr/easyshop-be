package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Id              int64     `json:"id"`
	Trxno           string    `json:"trxno"`
	Date            time.Time `json:"date"`
	Cust_id         int64     `json:"cust_id"`
	Proof_link      string    `json:"proof_link"`
	Pick_date       time.Time `json:"pick_date"`
	Tracking_number string    `json:"tracking_number"`
	Status          string    `json:"status"`
	Total           float64   `json:"total" gorm:"DEFAULT:0"`
	CreatedAt       time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (Order) TableName() string {
	return "order"
}

func (order Order) Validate() error {
	err := validation.Errors{
		"Trxno":     validation.Validate(order.Trxno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Date":      validation.Validate(order.Date, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Cust Id":   validation.Validate(order.Cust_id, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Pick Date": validation.Validate(order.Pick_date, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Total":     validation.Validate(order.Total, validation.NotNil.Error(utils.FIELD_NOTNIL), utils.ValidateNumeric()),
	}.Filter()

	return err
}

func (order Order) ID() int64 {
	return order.Id
}
