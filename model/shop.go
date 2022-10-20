package model

import "time"

type Shop struct {
	Id             int64     `json:"id"`
	ShopCategoryId int64     `json:"shop_category_id"`
	Name           string    `json:"name"`
	ImagePath      string    `json:"image_path"`
	Baseurl        string    `json:"baseurl"`
	Keyurl         string    `json:"keyurl"`
	ScrapItemName  string    `json:"scrap_item_name"`
	ScrapItemPrice string    `json:"scrap_item_price"`
	CreatedAt      time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (Shop) TableName() string {
	return "shop"
}

func (Shop) Validate() error {
	return nil
}
