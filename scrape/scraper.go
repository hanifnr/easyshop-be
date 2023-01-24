package scrape

import (
	"easyshop/model"
	"easyshop/utils"

	"github.com/gocolly/colly"
)

func GetProduct(shopId int64, url string) map[string]interface{} {
	switch shopId {
	case 1:
		matsukiyo := &Matsukiyo{}
		data := matsukiyo.GetProduct(shopId, url)
		return utils.MessageData(true, data)
	case 4:
		tokyuHands := &TokyuHands{}
		data := tokyuHands.GetProduct(shopId, url)
		return utils.MessageData(true, data)
	}
	return utils.MessageErr(false, utils.ErrExist, "Product not found!")
}

func GetListProducts(name string) []*Product {
	result := make([]*Product, 0)

	db := utils.GetDB()

	listShop := make([]*model.Shop, 0)
	db.Where("is_active = TRUE").Find(&listShop)

	for _, shop := range listShop {
		switch shop.Id {
		case 1:
			matsukiyo := &Matsukiyo{}
			products := matsukiyo.GetListProduct(name)
			result = append(result, products...)
		case 4:
			tokyuHands := &TokyuHands{}
			products := tokyuHands.GetListProduct(name)
			result = append(result, products...)
			// case 5:
			// 	loft := &Loft{}
			// 	products := loft.GetListProduct(name)
			// 	result = append(result, products...)
			// case 11:
			// 	yodobashi := &Yodobashi{}
			// 	products := yodobashi.GetListProduct(name)
			// 	result = append(result, products...)
		}
	}

	return result
}

func doScrap(url, selector string, f func(e *colly.HTMLElement)) {
	c := colly.NewCollector()

	c.OnHTML(selector, f)

	c.Visit(url)
}
