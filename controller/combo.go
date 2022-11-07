package controllers

import "easyshop/utils"

type ComboModel struct {
	Id   int64
	Name string
}

func GetCombo(page int, table string, param *utils.Param) map[string]interface{} {
	return ListModelColumns(page, "id,name", table, make([]*ComboModel, 0), param)
}
