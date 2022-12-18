package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Passport struct {
	Id              int64     `json:"id" gorm:"primary_key;auto_increment"`
	CustId          int64     `json:"cust_id"`
	Number          string    `json:"number"`
	Nationality     string    `json:"nationality"`
	BirthDate       time.Time `json:"birth_date"`
	StatusResidence string    `json:"status_residence"`
	CreatedAt       time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete        bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
	PassportExt     `gorm:"-"`
}
type PassportExt struct {
	CustName string `json:"cust_name"`
}

func (passport Passport) ID() int64 {
	return passport.Id
}

func (Passport) TableName() string {
	return "passport"
}

func (passport Passport) Validate() error {
	err := validation.Errors{
		"Number":           validation.Validate(passport.Number, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Nationality":      validation.Validate(passport.Nationality, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Birth Date":       validation.Validate(passport.BirthDate, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Status Residence": validation.Validate(passport.StatusResidence, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (passport *Passport) SetCreatedAt(time time.Time) {
	passport.CreatedAt = time
}

func (passport *Passport) SetUpdatedAt(time time.Time) {
	passport.UpdatedAt = time
}

func (passport *Passport) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("cust").Where("id = ?", passport.CustId).Scan(&passport.CustName)
}

func (passport *Passport) SetIsDelete(isDelete bool) {
	passport.IsDelete = isDelete
}
