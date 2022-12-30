package scrape

type Product struct {
	ShopId   int64  `json:"shop_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Name     string `json:"name,omitempty"`
	Image    string `json:"image,omitempty"`
	Price    string `json:"price,omitempty"`
	PriceTax string `json:"price_tax,omitempty"`
	Url      string `json:"url,omitempty"`
}
