package utils

import "gorm.io/gorm"

func QueryStringValue(query *gorm.DB) (string, error) {
	var result string
	err := query.Scan(&result).Error
	return result, err
}

func QueryListFind(page int, list *interface{}, param *Param) (map[string]interface{}, error) {
	var query *gorm.DB
	if page == 0 {
		query = db.Order("id ASC")
	} else {
		offset, limit := GetOffsetLimit(page)
		query = db.Offset(offset).Order("id ASC").Limit(limit)
	}
	param.ProcessFilter(query)
	totalRow := query.Debug().Find(list).RowsAffected

	if err := query.Error; err != nil {
		return nil, err
	}
	return RespPage(page, int(totalRow)), nil
}

func QueryListScan(page int, column, table string, list *interface{}, param *Param) (map[string]interface{}, error) {
	var query *gorm.DB
	param.ProcessFilter(query)
	if page == 0 {
		query = db.Select(column).Table(table).Order("id ASC")
	} else {
		offset, limit := GetOffsetLimit(page)
		query = db.Select(column).Table(table).Offset(offset).Order("id ASC").Limit(limit)
	}
	totalRow := query.Debug().Scan(list).RowsAffected

	if err := query.Error; err != nil {
		return nil, err
	}
	return RespPage(page, int(totalRow)), nil
}
