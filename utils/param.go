package utils

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Param struct {
	Page      *int
	StartDate *time.Time
	EndDate   *time.Time
	Id        *int64
	Name      *string
	Status    *string
	CustId    *int64
	ShopId    *int64
	Imported  *bool
}

func ProcessParam(r *http.Request) *Param {
	paramPage := ParamToInt(r.URL.Query().Get("page"))
	paramStartDate := ParamToTime(r.URL.Query().Get("startdate"), true)
	paramEndDate := ParamToTime(r.URL.Query().Get("enddate"), false)
	paramId := ParamToInt64(r.URL.Query().Get("id"))
	paramName := ParamToString(r.URL.Query().Get("name"))
	paramStatus := ParamToString(r.URL.Query().Get("status"))
	paramCustId := ParamToInt64(r.URL.Query().Get("cust_id"))
	paramShopId := ParamToInt64(r.URL.Query().Get("shop_id"))
	paramImported := ParamToBool(r.URL.Query().Get("imported"))

	return &Param{
		Page:      paramPage,
		StartDate: paramStartDate,
		EndDate:   paramEndDate,
		Id:        paramId,
		Name:      paramName,
		Status:    paramStatus,
		CustId:    paramCustId,
		ShopId:    paramShopId,
		Imported:  paramImported,
	}
}

func (param *Param) ProcessFilter(db *gorm.DB) {
	if param.StartDate != nil && param.EndDate != nil {
		db.Where("CAST(date AS DATE) BETWEEN ? AND ?", param.StartDate, param.EndDate)
	}
	if param.Id != nil {
		db.Where("id = ?", param.Id)
	}
	if param.Name != nil {
		db.Where("UPPER(name) LIKE ?", PercentText(strings.ToUpper(*param.Name)))
	}
	if param.Status != nil {
		db.Where("UPPER(status) = ?", strings.ToUpper(*param.Status))
	}
	if param.CustId != nil {
		db.Where("cust_id = ?", param.CustId)
	}
	if param.ShopId != nil {
		db.Where("shop_id = ?", param.ShopId)
	}
	if param.Imported != nil {
		db.Where("imported = ?", param.Imported)
	}
}

func ParamToString(param string) *string {
	if param == "" {
		return nil
	} else {
		return &param
	}
}

func ParamToInt(param string) *int {
	if param == "" {
		return nil
	}
	value, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}
	return &value
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

func ParamToBool(param string) *bool {
	if param == "" {
		return nil
	} else {
		value, err := strconv.ParseBool(param)
		if err != nil {
			return nil
		}
		return &value
	}
}

func PercentText(text string) string {
	return "%" + text + "%"
}
