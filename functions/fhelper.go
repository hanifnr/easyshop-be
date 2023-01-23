package functions

import "gorm.io/gorm"

func IsDataExist(id int64, model interface{}, db *gorm.DB) bool {
	var isExist bool
	db.Select("COUNT(id) > 0").Model(model).Where("id = ? AND is_delete = FALSE", id).Scan(&isExist)
	return isExist
}
