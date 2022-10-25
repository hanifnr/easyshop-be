package functions

import (
	"easyshop/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func FGetNewNo(name string, db *gorm.DB) string {
	var result string

	nCount := &model.NCount{}
	db.Debug().Where("upper(name) = ?", strings.ToUpper(name)).Find(&nCount)

	numberLen := len(strconv.Itoa(nCount.Number))
	for numberLen < nCount.Length {
		result = result + "0"
		numberLen++
	}
	db.Exec("UPDATE ncount SET number=number+1 WHERE code = ?", nCount.Code)
	return nCount.Code + result + strconv.Itoa(nCount.Number)
}
