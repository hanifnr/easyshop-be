package scrape

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetRupiah() float64 {
	var result float64
	doScrap(
		"https://kumiai.remit.co.jp/exchange/",
		"table.rate_table>tbody>tr:nth-of-type(4)>td:nth-of-type(2)",
		func(e *colly.HTMLElement) {
			result, _ = strconv.ParseFloat(strings.Split(e.Text, " ")[0], 64)
			result = result + 3
		},
	)
	return result
}
