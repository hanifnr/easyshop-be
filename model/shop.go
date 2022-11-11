package model

import "time"

type Shop struct {
	Id              int64     `json:"id"`
	ShopCategoryId  int64     `json:"shop_category_id"`
	Name            string    `json:"name"`
	ImagePath       string    `json:"image_path"`
	Baseurl         string    `json:"baseurl"`
	Keyurl          string    `json:"keyurl"`
	ScrapeItemName  string    `json:"scrape_item_name"`
	ScrapeItemPrice string    `json:"scrape_item_price"`
	CreatedAt       time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
	Isactive        bool      `json:"isactive" gorm:"DEFAULT:TRUE"`
	ScrapeImage     string    `json:"scrape_image"`
}

func (shop Shop) ID() int64 {
	return shop.Id
}

func (Shop) TableName() string {
	return "shop"
}

func (Shop) Validate() error {
	return nil
}
