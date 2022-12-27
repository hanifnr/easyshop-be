package model

import (
	"easyshop/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Status struct {
	Id   int64  `json:"id" gorm:"primary_key;auto_increment"`
	Code string `json:"code"`
	Name string `json:"name"`
	Idx  string `json:"idx"`
}

func (status Status) ID() int64 {
	return status.Id
}

func (Status) TableName() string {
	return "status"
}

func (status Status) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(status.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}
