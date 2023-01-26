package scrape

import (
	"easyshop/utils"
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type Loft struct{}

func (l *Loft) GetProduct(shopId int64, url string) *Product {
	var product *Product
	doScrap(
		url,
		"div.itemcontetnt-container>div.itemcontetnt-container",
		func(e *colly.HTMLElement) {
			code := utils.FormatPrice(e.ChildText("div.itemspec-container>div.tablebox>dl:nth-of-type(1)>dd"))
			name := e.ChildText("div.itemdetail-container>div>div.itemnamebox>ul>li.itemname")
			image := e.ChildText("div.itemdetail-container>div>form>div>ul>li>a>div>img")
			e.ForEach("div.itemdetail-container>div>form>div.selectcolorbox:nth-of-type(1)>ul>li:nth-of-type(1)>a:nth-of-type(1)", func(i int, h *colly.HTMLElement) {
				image := h.Attr("href")
				fmt.Println(image)
			})
			price := utils.FormatPrice(e.ChildText("div.itemdetail-container>div>div.pricebox>ul>li>dl>dd>ul>li>span.txtprice"))
			priceTax := utils.FormatPrice(e.ChildText("div.itemdetail-container>div>div.pricebox>ul>li>dl>dd>ul>li>span.txtzeinuki>span"))
			size := e.ChildText("div.itemspec-container>div.tablebox>dl:nth-of-type(4)>dd")

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

func (m *Loft) GetListProduct(name string) []*Product {
	result := make([]*Product, 0)
	link, _ := url.ParseQuery("q=" + name)
	doScrap(
		"https://www.loft.co.jp/store/goods/search.aspx?search=x&category=&keyword=&"+link.Encode(),
		"ul.itemlist.style-t",
		func(e *colly.HTMLElement) {
			fmt.Println(e)
			var product *Product
			e.ForEach("div.detailbox", func(i int, h *colly.HTMLElement) {
				name := h.ChildText("div>div>ul>li>a.js-enhanced-ecommerce-goods-name")
				price := utils.FormatPrice(h.ChildText("div>div>ul>li.sellingprice>span.txtprice"))
				priceTax := ""
				productUrl := h.ChildAttr("div>div>ul>li>a", "href")

				product = &Product{
					ShopId:   4,
					Name:     name,
					Price:    price,
					PriceTax: priceTax,
					Url:      productUrl,
				}
			})
			e.ForEach("div.imgbox>a>img", func(i int, h *colly.HTMLElement) {
				image := h.Attr("data-src")
				product.Image = "https://www.loft.co.jp" + image
			})

			result = append(result, product)
		},
	)
	// doScrap(
	// 	"https://www.loft.co.jp/store/goods/search.aspx?search=x&category=&keyword=&"+link.Encode(),
	// 	"div.imgbox>a>img",
	// 	func(e *colly.HTMLElement) {
	// 		image := e.Attr("data-src")

	// 		product := &Product{
	// 			ShopId:   4,
	// 			Name:     name,
	// 			Image:    image,
	// 			Price:    price,
	// 			PriceTax: priceTax,
	// 			Url:      productUrl,
	// 		}
	// 		result = append(result, product)
	// 	},
	// )
	return result
}
