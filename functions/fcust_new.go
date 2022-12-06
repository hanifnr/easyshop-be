package functions

import (
	"easyshop/model"
	"easyshop/utils"
	"strings"

	"gorm.io/gorm"
)

type FCustNew struct{}

func (f FCustNew) Run(m model.Model, db *gorm.DB) utils.StatusReturn {
	cust := m.(*model.Cust)
	var exist int
	if rows := db.Select("1").Table("cust").Where("UPPER(email) = ?", strings.ToUpper(cust.Email)).Scan(&exist).RowsAffected; rows > 0 {
		return utils.StatusReturn{ErrCode: utils.ErrExist, Message: "Email already registered!"}
	}
	if rows := db.Select("1").Table("cust").Where("country_code = ? AND phone_number = ?", cust.CountryCode, cust.PhoneNumber).Scan(&exist).RowsAffected; rows > 0 {
		return utils.StatusReturn{ErrCode: utils.ErrExist, Message: "Phone number already registered!"}
	}
	if rows := db.Select("1").Table("email_verif").Where("email = ? AND verified = TRUE", cust.Email).Scan(&exist).RowsAffected; rows == 0 {
		return utils.StatusReturn{ErrCode: utils.ErrExist, Message: "Email not yet verified!"}
	}
	return utils.StatusReturnOK()
}
