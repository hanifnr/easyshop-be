package model

import "time"

type ShopCategory struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (ShopCategory) TableName() string {
	return "shop_category"
}

func (ShopCategory) Validate() error {
	return nil
}
