package scrape

import (
	"easyshop/utils"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type ABCMart struct{}

func (a *ABCMart) GetProduct(shopId int64, url string) *Product {
	var product *Product
	doScrap(
		url,
		"div.stick-container",
		func(e *colly.HTMLElement) {
			code := e.ChildText("dl.spec-about>dd:nth-of-type(1)")
			name := e.ChildText("li.name")
			price := ""
			priceTax := utils.FormatPrice(e.ChildText("li.price"))
			size := ""

			product = &Product{
				Code:     code,
				Name:     name,
				Price:    price,
				PriceTax: priceTax,
				Size:     strings.Replace(size, "(パッケージ)", "", -1),
			}
		},
	)
	doScrap(
		url,
		"div.ph_item>div.wrap_ph",
		func(e *colly.HTMLElement) {
			image := e.DOM.Children().Children().Nodes[0].Attr[1].Val
			product.Image = image
		},
	)
	return product
}

func (m *ABCMart) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	link, _ := url.ParseQuery("keyword=" + name)
	doScrap(
		"https://gs.abc-mart.net/shop/goods/search.aspx?"+link.Encode()+"&fsstock=all",
		"div.list-item",
		func(e *colly.HTMLElement) {
			name := e.ChildText("p.name")
			image := e.ChildAttr("img.ph-item", "src")
			price := ""
			priceTax := utils.FormatPrice(e.ChildText("p.price"))
			productUrl := e.ChildAttr("a", "href")

			product := &Product{
				ShopId:   12,
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
