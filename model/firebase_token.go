package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type FirebaseToken struct {
	Id        int64     `json:"id" gorm:"primary_key;auto_increment"`
	Uid       string    `json:"uid"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	IsDelete  bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (token FirebaseToken) ID() int64 {
	return token.Id
}

func (FirebaseToken) TableName() string {
	return "firebase_token"
}

func (token FirebaseToken) Validate() error {
	err := validation.Errors{
		"Uid":   validation.Validate(token.Uid, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Token": validation.Validate(token.Token, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Type":  validation.Validate(token.Type, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (token *FirebaseToken) SetCreatedAt(time time.Time) {
	token.CreatedAt = time
}

func (token *FirebaseToken) SetUpdatedAt(time time.Time) {
	token.UpdatedAt = time
}

func (token *FirebaseToken) SetIsDelete(isDelete bool) {
	token.IsDelete = isDelete
}
