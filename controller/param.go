package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Param struct {
	StartDate *time.Time
	EndDate   *time.Time
	Status    *string
	CustId    *int64
	ShopId    *int64
}

func ProcessParam(r *http.Request) *Param {
	paramStartDate := ParamToTime(r.URL.Query().Get("startdate"), true)
	paramEndDate := ParamToTime(r.URL.Query().Get("enddate"), false)
	paramStatus := ParamToString(r.URL.Query().Get("status"))
	paramCustId := ParamToInt64(r.URL.Query().Get("cust_id"))
	paramShopId := ParamToInt64(r.URL.Query().Get("shop_id"))
	return &Param{
		StartDate: paramStartDate,
		EndDate:   paramEndDate,
		Status:    paramStatus,
		CustId:    paramCustId,
		ShopId:    paramShopId,
	}
}

func (param *Param) ProcessFilter(db *gorm.DB) {
	if param.StartDate != nil && param.EndDate != nil {
		db.Where("CAST(date AS DATE) BETWEEN ? AND ?", param.StartDate, param.EndDate)
	}
	if param.Status != nil {
		db.Where("status = ?", strings.ToUpper(*param.Status))
	}
	if param.CustId != nil {
		db.Where("cust_id = ?", param.CustId)
	}
	if param.ShopId != nil {
		db.Where("shop_id = ?", param.ShopId)
	}
}

func ParamToString(param string) *string {
	if param == "" {
		return nil
	} else {
		return &param
	}
}

func ParamToInt64(param string) *int64 {
	if param == "" {
		return nil
	} else {
		value, _ := strconv.ParseInt(param, 10, 64)
		return &value
	}
}

func ParamToTime(param string, isStartDate bool) *time.Time {
	var date string
	if param == "" {
		return nil
	} else {
		if isStartDate {
			date = param + " 00:00:00 +0700 WIB"
		} else {
			date = param + " 23:59:59 +0700 WIB"
		}
	}

	time, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", date)
	return &time
}
