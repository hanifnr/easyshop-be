package model

import "time"

type VoucherLog struct {
	Id            int64     `json:"id" gorm:"primary_key;auto_increment"`
	VoucherId     int64     `json:"voucher_id"`
	CustId        int64     `json:"cust_id"`
	OrderId       int64     `json:"order_id"`
	PartnershipId int64     `json:"partnership_id"`
	RewardAmount  float64   `json:"reward_amount"`
	RedeemAt      time.Time `json:"redeem_at"`
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
