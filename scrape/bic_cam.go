package scrape

import (
	"easyshop/utils"

	"github.com/gocolly/colly"
)

type BicCam struct{}

func (y *BicCam) GetProduct(shopId int64, url string) *Product {
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

func (m *BicCam) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	// link, _ := url.ParseQuery("q=" + name)
	doScrap(
		"https://www.biccamera.com/bc/category/?q=%83%5E%83u%83%8B%90%EE+%83z%83%8F%83C%83g",
		"html",
		func(e *colly.HTMLElement) {
			name := e.ChildAttr("li>p.bcs_title>a.bcs_item", "href")
			image := e.ChildAttr("li>p.bcs_image>a>img", "src")
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
