package utils

import "gorm.io/gorm"

func QueryStringValue(query *gorm.DB) (string, error) {
	var result string
	err := query.Scan(&result).Error
	return result, err
}

func QueryList(page int, columns, table string, list *interface{}, param *Param) (map[string]interface{}, error) {
	var query *gorm.DB
	if page == 0 {
		query = db.Order("id ASC")
	} else {
		offset, limit := GetOffsetLimit(page)
		query = db.Offset(offset).Order("id ASC").Limit(limit)
	}
	if columns != "" {
		query = db.Select(columns)
	}
	param.ProcessFilter(query)
	totalRow := query.Find(list).RowsAffected

	if err := query.Error; err != nil {
		return nil, err
	}
	return RespPage(page, int(totalRow)), nil
}
