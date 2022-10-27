package model

import "time"

type Purc struct {
	Id        int64     `json:"id"`
	Trxno     string    `json:"trxno"`
	Date      time.Time `json:"date"`
	ShopId    int64     `json:"shop_id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"  gorm:"DEFAULT:0"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	Isdelete  bool      `json:"isdelete" gorm:"DEFAULT:FALSE"`
}

func (Purc) TableName() string {
	return "purc"
}
