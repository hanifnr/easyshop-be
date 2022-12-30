package scrape

type ShopScrape interface {
	GetProduct() *Product
	GetListProducts() []*Product
}
