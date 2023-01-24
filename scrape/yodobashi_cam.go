package scrape

import (
	"easyshop/utils"

	"github.com/gocolly/colly"
)

type Yodobashi struct{}

func (y *Yodobashi) GetProduct(shopId int64, url string) *Product {
	var product *Product
	doScrap(
		url,
		"html",
		func(e *colly.HTMLElement) {
			code := utils.FormatPrice(e.ChildText("section#modItemDetail>div.detail>div#item-info>div.itemCode"))
			name := e.ChildText("section#modItemDetail>div.detail>h1.title")
			image := e.ChildAttr("section#modItemDetail>div.gallery>div.photo>img#first-image", "src")
			price := utils.FormatPrice(e.ChildText("section#modItemDetail>div.detail>div#item-info>div.price>div.normal>span.exclude"))
			priceTax := utils.FormatPrice(e.ChildText("section#modItemDetail>div.detail>div#item-info>div.price>div.normal>span.include>span"))
			size := e.ChildText("section#modItemInfo>div>div>table>tbody>tr:nth-of-type(2)>td")

			product = &Product{
				Code:     code,
				Name:     name,
				Image:    image,
				Price:    price,
				PriceTax: priceTax,
				Size:     size,
			}
		},
	)
	return product
}

func (m *Yodobashi) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	doScrap(
		"https://www.yodobashi.com/?word=%E5%89%8D%E3%81%AE%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%B8%E6%88%BB%E3%82%8B",
		"html",
		func(e *colly.HTMLElement) {
			name := e.ChildText("div.pName.fs14")
			image := ""
			price := ""
			priceTax := ""
			productUrl := ""

			product := &Product{
				ShopId:   11,
				Name:     name,
				Image:    image,
				Price:    price,
				PriceTax: priceTax,
				Url:      productUrl,
			}
			result = append(result, product)
		},
	)
	return result
}
