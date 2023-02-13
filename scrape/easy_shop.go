package scrape

import (
	"easyshop/utils"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type EasyShop struct{}

const COSMETIC = "COSMETIC"
const LIQUOR = "LIQUOR"
const MEDICINE = "MEDICINE"
const NECESSITY = "NECESSITY"
const ELECTRONIC = "ELECTRONIC"
const SUPPLEMENT = "SUPPLEMENT"

func (l *EasyShop) GetProduct(shopId int64, url string) *Product {
	var product *Product
	doScrap(
		url,
		"div:nth-of-type(3) > section:nth-of-type(1) > div:nth-of-type(1)",
		func(e *colly.HTMLElement) {
			code := strings.Replace(utils.FormatPrice(e.ChildText("span.sku_wrapper.detail-container")), "ode", "", -1)
			name := e.ChildText("div:nth-of-type(2) > div > div > div > h1")
			price := utils.FormatPrice(e.ChildText("div:nth-of-type(2) > div > div:nth-of-type(4) > div > p > span > bdi"))
			priceTax := ""
			size := e.ChildText("div:nth-of-type(2) > div > div:nth-of-type(6) > div > div > table > tbody > tr > td.marked-element")

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
		"div[data-id=\"4b679c0\"]>div>div>div>div",
		func(e *colly.HTMLElement) {
			selectDiv := e.DOM
			image := selectDiv.Children().Children().Children().Nodes[0].Attr[0].Val
			product.Image = image
		},
	)
	return product
}

func (m *EasyShop) GetProducts(category string, page int) ([]*Product, map[string]interface{}) {
	result := make([]*Product, 0)

	var url string
	switch strings.ToUpper(category) {
	case COSMETIC:
		url = "https://www.easyshop-jp.com/cosmetics-3"
	case LIQUOR:
		url = "https://www.easyshop-jp.com/liquor"
	case MEDICINE:
		url = "https://www.easyshop-jp.com/medicine"
	case NECESSITY:
		url = "https://www.easyshop-jp.com/daily-necessities"
	case ELECTRONIC:
		url = "https://www.easyshop-jp.com/electronics-2"
	case SUPPLEMENT:
		url = "https://www.easyshop-jp.com/supplement-vitamin-and-food"
	}

	start, end := GetPage(page)
	index := 1
	doScrap(
		url,
		"ul.products.elementor-grid.columns-4>li",
		func(e *colly.HTMLElement) {
			if index >= start && index <= end {
				fmt.Println(e)
				name := e.ChildText("h2.woocommerce-loop-product__title")
				image := e.ChildAttr("a>img", "src")
				price := utils.FormatPrice(e.ChildText("span.woocommerce-Price-amount.amount>bdi"))
				priceTax := ""
				productUrl := e.ChildAttr("a", "href")

				product := &Product{
					Index:    int64(index),
					ShopId:   19,
					Name:     name,
					Image:    image,
					Price:    price,
					PriceTax: priceTax,
					Url:      productUrl,
				}
				result = append(result, product)
			}
			index++
		},
	)
	return result, utils.RespPage(page, int(index-1))
}
