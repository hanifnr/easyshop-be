package model

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Order struct {
	Id             int64      `json:"id"`
	Trxno          string     `json:"trxno"`
	Date           time.Time  `json:"date"`
	CustId         int64      `json:"cust_id"`
	ProofLink      string     `json:"proof_link"`
	PickDate       *time.Time `json:"pick_date"`
	TrackingNumber string     `json:"tracking_number"`
	StatusCode     string     `json:"status_code"`
	Total          float64    `json:"total" gorm:"DEFAULT:0"`
	CreatedAt      time.Time  `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	IsDelete       bool       `json:"is_delete" gorm:"DEFAULT:FALSE"`
	Passport       string     `json:"passport"`
	AddrId         int64      `json:"addr_id"`
	ArrivalDate    *time.Time `json:"arrival_date"`
	OrderExt       `gorm:"-"`
}

type OrderExt struct {
	CustName   string `json:"cust_name"`
	AddrName   string `json:"addr_name"`
	StatusName string `json:"status_name"`
}

func (Order) TableName() string {
	return "order"
}

func (order Order) Validate() error {
	err := validation.Errors{
		"Trxno":   validation.Validate(order.Trxno, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Date":    validation.Validate(order.Date, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Cust Id": validation.Validate(order.CustId, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Total":   validation.Validate(order.Total, validation.Required.Error(utils.FIELD_NOTNIL)),
	}.Filter()

	return err
}

func (order Order) ID() int64 {
	return order.Id
}

func (order *Order) GetTrxno() string {
	return order.Trxno
}

func (order *Order) SetTrxno(trxno string) {
	order.Trxno = trxno
}

func (order *Order) SetCreatedAt(time time.Time) {
	order.CreatedAt = time
}

func (order *Order) SetUpdatedAt(time time.Time) {
	order.UpdatedAt = time
}

func (order *Order) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("cust").Where("id = ?", order.CustId).Scan(&order.CustName)
	db.Select("name").Table("addr").Where("id = ?", order.AddrId).Scan(&order.AddrName)
	db.Select("name").Table("status").Where("code = ?", order.StatusCode).Scan(&order.StatusName)
}

func (order *Order) SetIsDelete(isDelete bool) {
	order.IsDelete = isDelete
}
