package scrape

import (
	"easyshop/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	Id          int64     `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Index       int64     `json:"idx,omitempty" gorm:"-"`
	ShopId      int64     `json:"shop_id,omitempty" gorm:"-"`
	Code        string    `json:"code,omitempty"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Price       string    `json:"price"`
	PriceTax    string    `json:"price_tax"`
	Url         string    `json:"url,omitempty" gorm:"-"`
	Size        string    `json:"size"`
	ReqOrderId  *int64    `json:"req_order_id"`
	ReqOrderDno *int      `json:"req_order_dno"`
	IsDelete    bool      `json:"is_delete" gorm:"DEFAULT:FALSE"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"CURRENT_TIMESTAMP"`
}

func (product Product) ID() int64 {
	return product.Id
}

func (Product) TableName() string {
	return "product"
}

func (product Product) Validate() error {
	err := validation.Errors{
		"Code": validation.Validate(product.Code, validation.Required.Error(utils.FIELD_REQUIRED)),
		"Name": validation.Validate(product.Name, validation.Required.Error(utils.FIELD_REQUIRED)),
	}.Filter()

	return err
}

func (product *Product) SetCreatedAt(time time.Time) {
	product.CreatedAt = time
}

func (product *Product) SetUpdatedAt(time time.Time) {
	product.UpdatedAt = time
}

func (product *Product) SetIsDelete(isDelete bool) {
	product.IsDelete = isDelete
}
