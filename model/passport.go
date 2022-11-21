package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Passport struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	CustId      int64     `json:"cust_id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
	Number      string    `json:"number"`
	Nationality string    `json:"nationality"`
	BirthDate   time.Time `json:"birth_date"`
	IssueDate   time.Time `json:"issue_date"`
	ExpDate     time.Time `json:"exp_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete    bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
}

func (passport Passport) ID() int64 {
	return passport.Id
}

func (Passport) TableName() string {
	return "passport"
}

func (passport Passport) Validate() error {
	err := validation.Errors{
		"Name":         validation.Validate(passport.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Country Code": validation.Validate(passport.CountryCode, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Number":       validation.Validate(passport.Number, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Nationality":  validation.Validate(passport.Nationality, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Birth Date":   validation.Validate(passport.BirthDate, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Issue Date":   validation.Validate(passport.IssueDate, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Exp Date":     validation.Validate(passport.ExpDate, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (passport *Passport) SetCreatedAt(time time.Time) {
	passport.CreatedAt = time
}

func (passport *Passport) SetUpdatedAt(time time.Time) {
	passport.UpdatedAt = time
}
