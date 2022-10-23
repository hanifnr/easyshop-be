package utils

import "math"

const PAGE_LIMIT = 2

func GetOffsetLimit(page int) (int, int) {
	if page == 1 {
		return 0, PAGE_LIMIT
	} else {
		return (page - 1) * PAGE_LIMIT, PAGE_LIMIT
	}
}

func RespPage(page, totalRow int) map[string]interface{} {
	totalPage := int(math.Ceil(float64(totalRow) / float64(PAGE_LIMIT)))
	return map[string]interface{}{"page": page, "total_page": totalPage, "total_row": totalRow}
}
