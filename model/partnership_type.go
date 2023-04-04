package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PartnershipType struct {
	Id   int64  `json:"id" gorm:"primary_key;auto_increment"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (partnershipType PartnershipType) ID() int64 {
	return partnershipType.Id
}

func (PartnershipType) TableName() string {
	return "partnershipType"
}

func (partnershipType PartnershipType) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(partnershipType.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}
