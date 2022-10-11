package model

type Purcd struct {
	Id       int64   `json:"id"`
	PurcId   int64   `json:"purc_id"`
	Dno      int     `json:"dno"`
	OrderId  int64   `json:"order_id"`
	OrderDno int     `json:"order_dno"`
	Qty      float32 `json:"qty" gorm:"DEFAULT:0"`
	Qtywh    float32 `json:"qtywh" gorm:"DEFAULT:0"`
	Price    float64 `json:"price" gorm:"DEFAULT:0"`
	Subtotal float64 `json:"subtotal" gorm:"DEFAULT:0"`
}

func (Purcd) TableName() string {
	return "purcd"
}
