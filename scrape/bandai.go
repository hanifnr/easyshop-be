package scrape

import (
	"easyshop/utils"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Bandai struct{}

func (b *Bandai) GetProduct(shopId int64, url string) *Product {
	var product *Product
	baseUrl := "https://www.bandai.co.jp/catalog/item.php?jan_cd="
	doScrap(
		url,
		"div.pg-productInfo",
		func(e *colly.HTMLElement) {
			code := strings.Replace(url, baseUrl, "", -1)
			name := e.ChildText("div.pg-productInfo__name>h2")
			image := e.ChildAttr("div.pg-productInfo__pic>img", "src")
			price := ""
			priceTax := utils.FormatPrice(e.ChildText("div.pg-productInfo__price>span"))
			size := ""

			product = &Product{
				Code:     code,
				Name:     name,
				Image:    image,
				Price:    price,
				PriceTax: priceTax,
				Size:     strings.Replace(size, "(パッケージ)", "", -1),
			}
		},
	)
	return product
}

func (m *Bandai) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	baseUrl := "https://www.bandai.co.jp/catalog/"
	doScrap(
		fmt.Sprintf("https://www.bandai.co.jp/catalog/search.php?freeword=%s&freebutton=#result", name),
		"a.pg-card",
		func(e *colly.HTMLElement) {
			name := e.ChildText("div.pg-card__name")
			image := e.ChildAttr("div.pg-card__thumbnail>img", "src")
			price := ""
			priceTax := utils.FormatPrice(e.ChildText("div.pg-card__detail"))
			priceTax = priceTax[:len(priceTax)-6]
			productUrl := baseUrl + e.Attr("href")

			product := &Product{
				ShopId:   18,
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
