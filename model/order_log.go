package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderLog struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	OrderId     int64     `json:"order_id"`
	StatusCode  string    `json:"status_code"`
	Note        string    `json:"note"`
	Date        time.Time `json:"date"`
	OrderLogExt `gorm:"-"`
}

type OrderLogExt struct {
	StatusName string `json:"status_name"`
}

func (orderLog OrderLog) ID() int64 {
	return orderLog.Id
}

func (OrderLog) TableName() string {
	return "order_log"
}

func (orderLog OrderLog) Validate() error {
	return nil
}

func (orderLog *OrderLog) SetValueModelExt(db *gorm.DB) {
	db.Select("name").Table("status").Where("code = ?", orderLog.StatusCode).Scan(&orderLog.StatusName)
}
