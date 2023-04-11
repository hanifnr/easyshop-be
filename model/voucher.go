package model

import (
	"database/sql"
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Voucher struct {
	Id            int64        `json:"id" gorm:"primary_key;auto_increment"`
	Code          string       `json:"code"`
	Amount        string       `json:"amount"`
	Qty           *float32     `json:"qty,omitempty"`
	QtyUsed       float32      `json:"qty_used,omitempty"`
	Startdate     sql.NullTime `json:"startdate"`
	Enddate       sql.NullTime `json:"enddate"`
	PartnershipId *int64       `json:"partnership_id,omitempty"`
	Note          string       `json:"note"`
	IsDelete      bool         `json:"is_delete"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

func (voucher Voucher) ID() int64 {
	return voucher.Id
}

func (Voucher) TableName() string {
	return "voucher"
}

func (voucher Voucher) Validate() error {
	err := validation.Errors{
		"Code": validation.Validate(voucher.Code, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (voucher *Voucher) SetCreatedAt(time time.Time) {
	voucher.CreatedAt = time
}

func (voucher *Voucher) SetUpdatedAt(time time.Time) {
	voucher.UpdatedAt = time
}

func (voucher *Voucher) SetValueModelExt(db *gorm.DB) {
}

func (voucher *Voucher) SetIsDelete(isDelete bool) {
	voucher.IsDelete = isDelete
}
