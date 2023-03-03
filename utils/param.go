package utils

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Param struct {
	Page        *int
	StartDate   *time.Time
	EndDate     *time.Time
	Id          *int64
	Code        *string
	Name        *string
	StatusCode  *string
	CustId      *int64
	ShopId      *int64
	Imported    *bool
	PhoneNumber *string
	CountryCode *string
	IsActive    *bool
	IsDelete    *bool
	OrderId     *int64
	Email       *string
	OrderBy     *string
	Taxed       *bool
}

func ProcessParam(r *http.Request) *Param {
	paramPage := ParamToInt(r.URL.Query().Get("page"))
	paramStartDate := ParamToTime(r.URL.Query().Get("startdate"), true)
	paramEndDate := ParamToTime(r.URL.Query().Get("enddate"), false)
	paramId := ParamToInt64(r.URL.Query().Get("id"))
	paramCode := ParamToString(r.URL.Query().Get("code"))
	paramName := ParamToString(r.URL.Query().Get("name"))
	paramStatusCode := ParamToString(r.URL.Query().Get("status_code"))
	paramCustId := ParamToInt64(r.URL.Query().Get("cust_id"))
	paramShopId := ParamToInt64(r.URL.Query().Get("shop_id"))
	paramImported := ParamToBool(r.URL.Query().Get("imported"))
	paramPhoneNumber := ParamToString(r.URL.Query().Get("phone_number"))
	paramCountryCode := ParamToString(r.URL.Query().Get("country_code"))
	paramIsActive := ParamToBool(r.URL.Query().Get("is_active"))
	paramIsDelete := ParamToBool(r.URL.Query().Get("is_delete"))
	paramOrderId := ParamToInt64(r.URL.Query().Get("order_id"))
	paramEmail := ParamToString(r.URL.Query().Get("email"))
	paramOrderBy := ParamToString(r.URL.Query().Get("order_by"))
	paramTaxed := ParamToBool(r.URL.Query().Get("taxed"))

	return &Param{
		Page:        paramPage,
		StartDate:   paramStartDate,
		EndDate:     paramEndDate,
		Id:          paramId,
		Code:        paramCode,
		Name:        paramName,
		StatusCode:  paramStatusCode,
		CustId:      paramCustId,
		ShopId:      paramShopId,
		Imported:    paramImported,
		PhoneNumber: paramPhoneNumber,
		CountryCode: paramCountryCode,
		IsActive:    paramIsActive,
		IsDelete:    paramIsDelete,
		OrderId:     paramOrderId,
		Email:       paramEmail,
		OrderBy:     paramOrderBy,
		Taxed:       paramTaxed,
	}
}

func (param *Param) ProcessFilter(db *gorm.DB) {
	if param.StartDate != nil && param.EndDate != nil {
		db.Where("CAST(date AS DATE) BETWEEN ? AND ?", param.StartDate, param.EndDate)
	}
	if param.Id != nil {
		db.Where("id = ?", param.Id)
	}
	if param.Code != nil {
		db.Where("UPPER(code) = ?", strings.ToUpper(*param.Code))
	}
	if param.Name != nil {
		db.Where("UPPER(name) LIKE ?", PercentText(strings.ToUpper(*param.Name)))
	}
	if param.StatusCode != nil {
		db.Where("UPPER(status_code) = ?", strings.ToUpper(*param.StatusCode))
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
	if param.PhoneNumber != nil {
		db.Where("phone_number = ?", param.PhoneNumber)
	}
	if param.CountryCode != nil {
		db.Where("UPPER(country_code) = UPPER(?)", param.CountryCode)
	}
	if param.IsActive != nil {
		db.Where("is_active = ?", param.IsActive)
	}
	if param.IsDelete != nil {
		db.Where("is_delete = ?", param.IsDelete)
	}
	if param.OrderId != nil {
		db.Where("order_id = ?", param.OrderId)
	}
	if param.Email != nil {
		db.Where("UPPER(email) = UPPER(?)", param.Email)
	}
	if param.Taxed != nil {
		db.Where("taxed = ?", param.Taxed)
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

func (param *Param) SetDefaultDelete() {
	if param.IsDelete == nil {
		isDelete := false
		param.IsDelete = &isDelete
	}
}
