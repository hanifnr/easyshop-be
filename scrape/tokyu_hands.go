package scrape

import (
	"easyshop/utils"
	"fmt"

	"github.com/gocolly/colly"
)

type TokyuHands struct{}

func (t *TokyuHands) GetProduct(shopId int64, url string) *Product {
	var product *Product
	doScrap(
		url,
		"main#detail.content:nth-of-type(1)>div.modContainer.content:nth-of-type(1)",
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

func (m *TokyuHands) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	doScrap(
		fmt.Sprintf("https://hands.net/search/?q=%s", name),
		"ul#search-result-list.list>li.item>div.detail",
		func(e *colly.HTMLElement) {
			name := e.ChildText("div.title>a")
			image := e.ChildAttr("div.photo>a>img", "src")
			price := ""
			priceTax := utils.FormatPrice(e.ChildText("div.meta>div.price"))
			productUrl := e.ChildAttr("div.title>a", "href")

			product := &Product{
				ShopId:   4,
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
