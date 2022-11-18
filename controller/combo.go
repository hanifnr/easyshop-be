package controllers

import "easyshop/utils"

type ComboModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetCombo(page int, table, order string, param *utils.Param) map[string]interface{} {
	return ListModelColumns("id,name", table, order, make([]*ComboModel, 0), param)
}
