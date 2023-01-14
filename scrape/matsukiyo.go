package scrape

import (
	"easyshop/utils"
	"fmt"
	"regexp"

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
			image := e.ChildAttr("div > div > div > ul > li:nth-of-type(1) > a", "style")
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
	doScrap(
		url,
		"div.goodsDetail",
		func(e *colly.HTMLElement) {
			size := e.ChildText("div:nth-of-type(7) > div > p:nth-of-type(2)")
			size = regexp.MustCompile(`^([^0-9])+`).ReplaceAllString(size, "")
			if size == "" {
				size = e.ChildText("div:nth-of-type(8) > div > p:nth-of-type(2)")
				size = regexp.MustCompile(`^([^0-9])+`).ReplaceAllString(size, "")
			}

			product.Size = size
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
