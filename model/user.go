package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id        int64     `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete  bool      `json:"id_delete"`
}

func (user User) ID() int64 {
	return user.Id
}

func (User) TableName() string {
	return "user"
}

func (user User) Validate() error {
	err := validation.Errors{
		"Name":     validation.Validate(user.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Password": validation.Validate(user.Password, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Email":    validation.Validate(user.Email, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (user *User) SetCreatedAt(time time.Time) {
	user.CreatedAt = time
}

func (user *User) SetUpdatedAt(time time.Time) {
	user.UpdatedAt = time
}

func (user *User) SetIsDelete(isDelete bool) {
	user.IsDelete = isDelete
}
