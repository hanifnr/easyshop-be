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
		"div.itemcontetnt-container>div.itemcontetnt-container",
		func(e *colly.HTMLElement) {
			code := utils.FormatPrice(e.ChildText("div.itemspec-container>div.tablebox>dl:nth-of-type(1)>dd"))
			name := e.ChildText("div.itemdetail-container>div>div.itemnamebox>ul>li.itemname")
			price := utils.FormatPrice(e.ChildText("div.itemdetail-container>div>div.pricebox>ul>li>dl>dd>ul>li>span.txtprice"))
			priceTax := utils.FormatPrice(e.ChildText("div.itemdetail-container>div>div.pricebox>ul>li>dl>dd>ul>li>span.txtzeinuki>span"))
			size := e.ChildText("div.itemspec-container>div.tablebox>dl:nth-of-type(4)>dd")

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
		"head",
		func(e *colly.HTMLElement) {
			image := e.ChildAttr("meta:nth-of-type(11)", "content")
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

	start, end := utils.GetOffsetLimit(page)
	index := 0
	doScrap(
		url,
		"ul.products.elementor-grid.columns-4>li",
		func(e *colly.HTMLElement) {
			if index >= start && index < end {
				fmt.Println(e)
				name := e.ChildText("h2.woocommerce-loop-product__title")
				image := e.ChildAttr("a>img", "src")
				price := utils.FormatPrice(e.ChildText("span.woocommerce-Price-amount.amount>bdi"))
				priceTax := ""
				productUrl := e.ChildAttr("a", "href")

				product := &Product{
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
	return result, utils.RespPage(page, int(index+1))
}
