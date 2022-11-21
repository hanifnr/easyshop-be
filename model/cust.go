package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Cust struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    bool      `json:"is_active" gorm:"DEFAULT:TRUE"`
	CreatedAt   time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete    bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
}

func (cust Cust) ID() int64 {
	return cust.Id
}

func (Cust) TableName() string {
	return "cust"
}

func (cust Cust) Validate() error {
	err := validation.Errors{
		"Name":         validation.Validate(cust.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Email":        validation.Validate(cust.Email, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Country Code": validation.Validate(cust.CountryCode, validation.Required.Error(utils.FIELD_REQUIRED), utils.ValidateNumeric()),
		"Phone Number": validation.Validate(cust.PhoneNumber, validation.Required.Error(utils.FIELD_REQUIRED), utils.ValidateNumeric()),
	}.Filter()

	return err
}

func (cust *Cust) SetCreatedAt(time time.Time) {
	cust.CreatedAt = time
}

func (cust *Cust) SetUpdatedAt(time time.Time) {
	cust.UpdatedAt = time
}
