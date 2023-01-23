package scrape

type Product struct {
	ShopId   int64  `json:"shop_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Price    string `json:"price"`
	PriceTax string `json:"price_tax"`
	Url      string `json:"url,omitempty"`
	Size     string `json:"size"`
}
