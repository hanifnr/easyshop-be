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
	Passport    string    `json:"passport"`
	Status      string    `json:"status"`
	Isactive    *bool     `json:"isactive" gorm:"DEFAULT:TRUE"`
	CreatedAt   time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (Cust) TableName() string {
	return "cust"
}

func (cust Cust) Validate() error {
	err := validation.Errors{
		"Name":         validation.Validate(cust.Name, validation.Required.Error(FIELD_REQUIRED)),
		"Email":        validation.Validate(cust.Email, validation.Required.Error(FIELD_REQUIRED)),
		"Country Code": validation.Validate(cust.CountryCode, validation.Required.Error(FIELD_REQUIRED)),
		"Phone Number": validation.Validate(cust.PhoneNumber, validation.Required.Error(FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (cust Cust) CreateModel() map[string]interface{} {
	currentTime := time.Now()
	cust.Status = "W"
	cust.CreatedAt = currentTime
	cust.UpdatedAt = currentTime
	if err := Save(&cust); err != nil {
		return utils.Message(false, err.Error())
	}
	return utils.MessageData(true, "Saved succesfully!", cust)
}
