package model

import "time"

type VoucherLog struct {
	Id        int64     `json:"id" gorm:"primary_key;auto_increment"`
	VoucherId int64     `json:"voucher_id"`
	CustId    int64     `json:"cust_id"`
	RedeemAt  time.Time `json:"redeem_at"`
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
