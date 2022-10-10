package model

type Orderd struct {
	Id        int64   `json:"id"`
	OrderId   int64   `json:"order_id"`
	Dno       int     `json:"dno"`
	ShopId    int64   `json:"shop_id"`
	ProductId string  `json:"productid"`
	Name      string  `json:"name"`
	Qty       float32 `json:"qty" gorm:"DEFAULT:0"`
	Qtypurc   float32 `json:"qtypurc" gorm:"DEFAULT:0"`
	Qtywh     float32 `json:"qtywh" gorm:"DEFAULT:0"`
	Price     float64 `json:"price" gorm:"DEFAULT:0"`
	Subtotal  float64 `json:"subtotal" gorm:"DEFAULT:0"`
}

func (Orderd) TableName() string {
	return "orderd"
}
