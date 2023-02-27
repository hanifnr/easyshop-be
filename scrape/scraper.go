package scrape

import (
	"easyshop/model"
	"easyshop/utils"

	"github.com/gocolly/colly"
)

func GetProduct(shopId int64, url string) map[string]interface{} {
	var data *Product
	switch shopId {
	case 1:
		matsukiyo := &Matsukiyo{}
		data = matsukiyo.GetProduct(shopId, url)
	case 4:
		tokyuHands := &TokyuHands{}
		data = tokyuHands.GetProduct(shopId, url)
	case 5:
		loft := &Loft{}
		data = loft.GetProduct(shopId, url)
	case 12:
		abcMart := &ABCMart{}
		data = abcMart.GetProduct(shopId, url)
	case 18:
		bandai := &Bandai{}
		data = bandai.GetProduct(shopId, url)
	case 19:
		easyShop := &EasyShop{}
		data = easyShop.GetProduct(shopId, url)
	}
	if data.Code != "" {
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
		// case 1:
		// 	matsukiyo := &Matsukiyo{}
		// 	products := matsukiyo.GetListProduct(name)
		// 	result = append(result, products...)
		// case 4:
		// 	tokyuHands := &TokyuHands{}
		// 	products := tokyuHands.GetListProduct(name)
		// 	result = append(result, products...)
		// case 5:
		// 	loft := &Loft{}
		// 	products := loft.GetListProduct(name)
		// 	result = append(result, products...)
		// // case 9:
		// // 	bicCam := &BicCam{}
		// // 	products := bicCam.GetListProduct(name)
		// // 	result = append(result, products...)
		// // case 11:
		// // 	yodobashi := &Yodobashi{}
		// // 	products := yodobashi.GetListProduct(name)
		// // 	result = append(result, products...)
		// case 12:
		// 	abc := &ABCMart{}
		// 	products := abc.GetListProduct(name)
		// 	result = append(result, products...)
		case 18:
			bandai := &Bandai{}
			products := bandai.GetListProduct(name)
			result = append(result, products...)
		}
	}

	return result
}

func GetTopProducts() []*Product {
	result := make([]*Product, 0)

	db := utils.GetDB()

	listShop := make([]*model.Shop, 0)
	db.Where("is_active = TRUE").Find(&listShop)

	for _, shop := range listShop {
		switch shop.Id {
		case 1:
			matsukiyo := &Matsukiyo{}
			products := matsukiyo.GetTopProduct()
			result = append(result, products...)
		}
	}

	return result
}

func GetEasyShopProducts(category string, page int) ([]*Product, map[string]interface{}) {
	easyshop := &EasyShop{}
	return easyshop.GetProducts(category, page)
}

func doScrap(url, selector string, f func(e *colly.HTMLElement)) {
	c := colly.NewCollector()

	c.OnHTML(selector, f)

	c.OnResponse(func(r *colly.Response) {})

	c.Visit(url)
}

func DoScrapResponse(url, selector string, fH func(e *colly.HTMLElement), fR func(f *colly.Response)) {
	c := colly.NewCollector()

	c.OnHTML(selector, fH)

	c.OnResponse(fR)

	c.Visit(url)
}

func GetPage(page int) (int, int) {
	start, _ := utils.GetOffsetLimit(page)
	return start + 1, start + 20
}
