package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"net/http"
	"time"
)

var RegisterEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	CreateModelAction(emailVerifController, w, r)
}

var UpdateEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	UpdateModelAction(emailVerifController, w, r)
}

var ViewEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	ViewModelAction(emailVerifController, w, r)
}

var ListEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	ListModelAction(emailVerifController, w, r)
}

var VerifyEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	emailVerifController.VerifyEmail(w, r)
}

type EmailVerifController struct {
	EmailVerif model.EmailVerif
}

func (emailVerifController *EmailVerifController) Model() model.Model {
	return &emailVerifController.EmailVerif
}

func (emailVerifController *EmailVerifController) FNew() functions.SQLFunction {
	return nil
}
func (emailVerifController *EmailVerifController) FDelete() functions.SQLFunction {
	return nil
}

func (emailVerifController *EmailVerifController) CreateModel() map[string]interface{} {
	db := utils.GetDB()
	emailVerif := &emailVerifController.EmailVerif
	row := db.Where("UPPER(email) = UPPER(?)", emailVerif.Email).Find(emailVerif).RowsAffected
	if row > 0 {
		if emailVerif.Verified {
			return utils.MessageErr(false, utils.ErrRequest, "This email has already verified!")
		}
		return emailVerifController.UpdateModel()
	} else {
		if retval := CreateModel(emailVerifController, func(m model.Model) {
			emailVerif.GenerateCode()
		}); retval.ErrCode != 0 {
			return utils.MessageErr(false, retval.ErrCode, retval.Message)
		}
		return utils.Message(true)
	}
}

func (emailVerifController *EmailVerifController) ViewModel(id int64) map[string]interface{} {
	emailVerif := &model.EmailVerif{}
	if retval := ViewModel(id, emailVerif); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, emailVerif)
}

func (emailVerifController *EmailVerifController) UpdateModel() map[string]interface{} {
	retval, _ := UpdateModel(emailVerifController, &model.EmailVerif{}, func(modelSrc, modelTemp model.Model) {
		emailVerif := modelSrc.(*model.EmailVerif)
		emailVerif.GenerateCode()
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (emailVerifController *EmailVerifController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, emailVerifController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (emailVerifController *EmailVerifController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("emailVerif", "id ASC", make([]*model.EmailVerif, 0), param)
}

func (emailVerifController *EmailVerifController) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	type Email struct {
		Email string
		Code  string
	}
	email := &Email{}
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}

	db := utils.GetDB().Begin()

	emailData := &model.EmailVerif{}
	row := db.Where("UPPER(email) = UPPER(?)", email.Email).Find(&emailData).RowsAffected
	if row > 0 {
		if emailData.Verified {
			utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "This email has already verified!"))
			return
		}
		if email.Code == emailData.VerifCode {
			emailData.Verified = true
			emailData.VerifiedAt = time.Now()
			if err := model.Save(emailData, db); err != nil {
				db.Rollback()
				utils.Respond(w, utils.MessageErr(false, utils.ErrSQLCreate, err.Error()))
				return
			}
			utils.Respond(w, utils.Message(true))
		} else {
			utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "Wrong verification code!"))
		}
	} else {
		utils.Respond(w, utils.MessageErr(false, utils.ErrExist, "Email not yet registered!"))
	}
	db.Commit()
}
