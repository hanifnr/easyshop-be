package controllers

import (
	"easyshop/model"
	"easyshop/utils"
	"net/http"
)

var GetTaxOffice = func(w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	paramGeneralTotal := r.URL.Query().Get("general_total")
	paramConsumTotal := r.URL.Query().Get("consum_total")
	paramLqTotal := r.URL.Query().Get("lq_total")
	paramLqTotalNum := r.URL.Query().Get("lq_total_num")

	db := utils.GetDB()
	data, statusReturn := model.GetTaxOffice(id, paramGeneralTotal, paramConsumTotal, paramLqTotal, paramLqTotalNum, w, db)

	if statusReturn.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, statusReturn.ErrCode, statusReturn.Message))
		return
	}

	utils.Respond(w, utils.MessageData(true, data))
}
