package model

import (
	"easyshop/utils"
	"math/rand"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type EmailVerif struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	Email       string    `json:"email"`
	VerifCode   string    `json:"verif_code"`
	Verified    bool      `json:"verified"`
	VerifiedAt  time.Time `json:"verified_at"`
	GeneratedAt time.Time `json:"generated_at"`
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

func (email *EmailVerif) GenerateCode() {
	email.VerifCode = strconv.Itoa(rand.Intn(10000))
	email.GeneratedAt = time.Now()
}
