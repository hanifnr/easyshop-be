package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

const EMAIL_VERIF_REGISTER = 0
const EMAIL_VERIF_AUTH = 1

type EmailVerif struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	Email       string    `json:"email"`
	VerifCode   string    `json:"verif_code"`
	Verified    bool      `json:"verified"`
	VerifiedAt  time.Time `json:"verified_at"`
	GeneratedAt time.Time `json:"generated_at"`
	AuthCode    string    `json:"auth_code"`
}

func (email EmailVerif) ID() int64 {
	return email.Id
}

func (EmailVerif) TableName() string {
	return "email_verif"
}

func (email EmailVerif) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(email.Email, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (email *EmailVerif) GenerateCode(mode int) {
	if mode == EMAIL_VERIF_REGISTER {
		email.VerifCode = utils.RandInt(4)
	} else if mode == EMAIL_VERIF_AUTH {
		email.AuthCode = utils.RandInt(4)
	}
	email.GeneratedAt = time.Now()
}
