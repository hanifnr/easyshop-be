package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Addr struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	CustId      int64     `json:"cust_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	ZipCode     string    `json:"zip_code"`
	CountryCode string    `json:"country_code"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	FullAddress string    `json:"full_address"`
	IsDelete    bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
	CreatedAt   time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	AddrExt     `gorm:"-"`
}

type AddrExt struct {
	CustName string `json:"cust_name"`
}

func (addr Addr) ID() int64 {
	return addr.Id
}

func (Addr) TableName() string {
	return "addr"
}

func (addr Addr) Validate() error {
	err := validation.Errors{
		"Name":         validation.Validate(addr.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Phone Number": validation.Validate(addr.PhoneNumber, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Zip Code":     validation.Validate(addr.ZipCode, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Country Code": validation.Validate(addr.CountryCode, validation.Required.Error(utils.FIELD_REQUIRED)),
		"City":         validation.Validate(addr.City, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Full Address": validation.Validate(addr.FullAddress, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (addr *Addr) SetCreatedAt(time time.Time) {
	addr.CreatedAt = time
}

func (addr *Addr) SetUpdatedAt(time time.Time) {
	addr.UpdatedAt = time
}

func (addr *Addr) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("cust").Where("id = ?", addr.CustId).Scan(&addr.CustName)
}

func (addr *Addr) SetIsDelete(isDelete bool) {
	addr.IsDelete = isDelete
}
