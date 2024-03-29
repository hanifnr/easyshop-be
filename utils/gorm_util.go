package utils

import (
	"strings"

	"gorm.io/gorm"
)

func QueryStringValue(query *gorm.DB) (string, error) {
	var result string
	err := query.Scan(&result).Error
	return result, err
}

func QueryListFind(table, order string, list *interface{}, param *Param, fJoin func(query *gorm.DB), fFilter func(query *gorm.DB)) (map[string]interface{}, error) {
	var query *gorm.DB
	var totalRow int64
	var page int

	if param.Page == nil {
		page = 0
	} else {
		page = *param.Page
	}

	query = db.Select("count(*)").Table(table)
	fJoin(query)
	param.ProcessFilter(query)
	fFilter(query)
	query.Scan(&totalRow)
	var orderBy string
	if param.OrderBy != nil {
		orderBy = strings.ToUpper(*param.OrderBy)
		if strings.Contains(orderBy, "DESCEND") {
			orderBy = strings.Replace(orderBy, "DESCEND", "DESC", -1)
		} else if strings.Contains(orderBy, "ASCEND") {
			orderBy = strings.Replace(orderBy, "ASCEND", "ASC", -1)
		}
	} else {
		orderBy = order
	}
	if page == 0 {
		query = db.Order(orderBy)
	} else {
		offset, limit := GetOffsetLimit(page)
		query = db.Offset(offset).Order(orderBy).Limit(limit)
	}
	fJoin(query)
	param.ProcessFilter(query)
	fFilter(query)
	query.Debug().Find(list)

	if err := query.Error; err != nil {
		return nil, err
	}
	return RespPage(page, int(totalRow)), nil
}

func QueryListScan(column, table, order string, list *interface{}, param *Param) (map[string]interface{}, error) {
	var query *gorm.DB
	var totalRow int64
	var page int

	if param.Page == nil {
		page = 0
	} else {
		page = *param.Page
	}

	query = db.Select("count(*)").Table(table)
	param.ProcessFilter(query)
	query.Scan(&totalRow)
	if page == 0 {
		query = db.Select(column).Table(table).Order(order)
	} else {
		offset, limit := GetOffsetLimit(page)
		query = db.Select(column).Table(table).Offset(offset).Order(order).Limit(limit)
	}
	param.ProcessFilter(query)
	query.Scan(list)

	if err := query.Error; err != nil {
		return nil, err
	}
	return RespPage(page, int(totalRow)), nil
}
