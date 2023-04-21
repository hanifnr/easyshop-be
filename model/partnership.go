package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Partnership struct {
	Id                int64     `json:"id" gorm:"primary_key;auto_increment"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	PartnershipTypeId int64     `json:"partnership_type_id"`
	SocialMedia       string    `json:"social_media"`
	PhoneNumber       string    `json:"phone_number"`
	IsDelete          bool      `json:"is_delete"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (partnership Partnership) ID() int64 {
	return partnership.Id
}

func (Partnership) TableName() string {
	return "partnership"
}

func (partnership Partnership) Validate() error {
	err := validation.Errors{
		"Name": validation.Validate(partnership.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (partnership *Partnership) SetCreatedAt(time time.Time) {
	partnership.CreatedAt = time
}

func (partnership *Partnership) SetUpdatedAt(time time.Time) {
	partnership.UpdatedAt = time
}

func (partnership *Partnership) SetValueModelExt(db *gorm.DB) {
}

func (partnership *Partnership) SetIsDelete(isDelete bool) {
	partnership.IsDelete = isDelete
}
