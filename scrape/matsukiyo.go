package scrape

import (
	"easyshop/utils"
	"fmt"

	"github.com/gocolly/colly"
)

type Matsukiyo struct{}

func (m *Matsukiyo) GetProduct(shopId int64, url string) *Product {
	var product *Product
	baseUrlImage := "https://www.matsukiyo.co.jp"
	doScrap(
		url,
		"div.ctBox01.clearfix",
		func(e *colly.HTMLElement) {
			code := utils.FormatPrice(e.ChildText("div.goodsBox.main > p.cpde"))
			name := e.ChildText("div.goodsBox.main > div.spHide > h3")
			image := e.ChildAttr("div > div > div > ul > li > a", "style")
			price := utils.FormatPrice(e.ChildText("div.goodsBox.main > p.price > span > span:first-of-type"))
			priceTax := utils.FormatPrice(e.ChildText("div.goodsBox.main > p.price > span > span.small"))

			image = baseUrlImage + image[22:len(image)-3]

			product = &Product{
				Code:     code,
				Name:     name,
				Image:    image,
				Price:    price,
				PriceTax: priceTax,
			}
		},
	)
	return product
}

func (m *Matsukiyo) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	doScrap(
		fmt.Sprintf("https://www.matsukiyo.co.jp/store/online/search?text=%s", name),
		"ul#itemList > li",
		func(e *colly.HTMLElement) {
			baseUrl := "https://www.matsukiyo.co.jp"

			name := e.ChildText("a > p.ttl")
			image := e.ChildAttr("p.img > a > img", "src")
			price := utils.FormatPrice(e.ChildText("div > div > p.price:first-of-type"))
			priceTax := utils.FormatPrice(e.ChildText("div > div > p.price.inTax"))
			productUrl := e.ChildAttr("p.img > a", "href")

			product := &Product{
				ShopId:   1,
				Name:     name,
				Image:    baseUrl + image,
				Price:    price,
				PriceTax: priceTax,
				Url:      baseUrl + productUrl,
			}
			result = append(result, product)
		},
	)
	return result
}
