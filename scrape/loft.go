package scrape

import (
	"easyshop/utils"
	"fmt"

	"github.com/gocolly/colly"
)

type Loft struct{}

func (l *Loft) GetProduct(shopId int64, url string) *Product {
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

func (m *Loft) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	doScrap(
		fmt.Sprintf("https://www.loft.co.jp/store/goods/search.aspx?search=x&category=&keyword=&q=%s", name),
		"html>body>div.main-container",
		func(e *colly.HTMLElement) {
			name := e.ChildText("div>div>ul>li>a.js-enhanced-ecommerce-goods-name")
			image := e.ChildAttr("div>div>a>img.lazyloaded", "src")
			price := utils.FormatPrice(e.ChildText("div>div>ul>li.sellingprice>span.txtprice"))
			priceTax := ""
			productUrl := e.ChildAttr("div>div>a", "href")

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
