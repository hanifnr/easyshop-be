package model

import "time"

type Wh struct {
	Id        int64     `json:"id"`
	Trxno     string    `json:"trxno"`
	Date      time.Time `json:"date"`
	ShopId    int64     `json:"shop_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	Isdelete  bool      `json:"isdelete" gorm:"DEFAULT:FALSE"`
}

func (Wh) TableName() string {
	return "wh"
}
