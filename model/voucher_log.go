package model

import (
	"time"

	"gorm.io/gorm"
)

type VoucherLog struct {
	Id            int64     `json:"id" gorm:"primary_key;auto_increment"`
	VoucherId     int64     `json:"voucher_id"`
	CustId        int64     `json:"cust_id"`
	OrderId       int64     `json:"order_id"`
	PartnershipId int64     `json:"partnership_id"`
	RewardAmount  float64   `json:"reward_amount"`
	RedeemAt      time.Time `json:"redeem_at"`
	VoucherLogExt `gorm:"-"`
}
type VoucherLogExt struct {
	CustEmail string `json:"cust_email"`
}

func (VoucherLog VoucherLog) ID() int64 {
	return VoucherLog.Id
}

func (VoucherLog) TableName() string {
	return "voucher_log"
}

func (VoucherLog VoucherLog) Validate() error {
	return nil
}

func (vl *VoucherLog) SetValueModelExt(db *gorm.DB) {
	db.Select("email").Table("cust").Where("id = ?", vl.CustId).Scan(&vl.CustEmail)
}
