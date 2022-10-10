package model

import "time"

type Order struct {
	Id              int64     `json:"id"`
	Trxno           string    `json:"trxno"`
	Date            time.Time `json:"date"`
	Cust_id         int64     `json:"cust_id"`
	Proof_link      string    `json:"proof_link"`
	Pick_date       time.Time `json:"pick_date"`
	Tracking_number string    `json:"tracking_number"`
	Status          string    `json:"status"`
	Total           float64   `json:"total" gorm:"DEFAULT:0"`
	CreatedAt       time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (Order) TableName() string {
	return "order"
}
