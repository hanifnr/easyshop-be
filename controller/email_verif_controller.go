package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

var RegisterEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{Mode: model.EMAIL_VERIF_REGISTER}
	CreateModelAction(emailVerifController, w, r)
}

var UpdateEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	UpdateModelAction(emailVerifController, w, r)
}

var AuthEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{Mode: model.EMAIL_VERIF_AUTH}
	emailVerifController.AuthEmail(w, r)
}

var ViewEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	ViewModelAction(emailVerifController, w, r)
}

var ListEmailVerif = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{}
	ListModelAction(emailVerifController, w, r)
}

var VerifyRegisterEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{Mode: model.EMAIL_VERIF_REGISTER}
	emailVerifController.VerifyEmail(w, r)
}

var VerifyAuthEmail = func(w http.ResponseWriter, r *http.Request) {
	emailVerifController := &EmailVerifController{Mode: model.EMAIL_VERIF_AUTH}
	emailVerifController.VerifyEmail(w, r)
}

type EmailVerifController struct {
	EmailVerif model.EmailVerif
	Mode       int
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
	row := db.Where("UPPER(email) = UPPER(?) AND UPPER(type) = UPPER(?)", emailVerif.Email, emailVerif.Type).Find(emailVerif).RowsAffected
	if row > 0 {
		if emailVerif.Verified {
			return utils.MessageErr(false, utils.ErrRequest, "This email has already verified!")
		}
		return emailVerifController.UpdateModel()
	} else {
		if retval := CreateModel(emailVerifController, func(m model.Model) {
			emailVerif.GenerateCode(emailVerifController.Mode)
		}); retval.ErrCode != 0 {
			return utils.MessageErr(false, retval.ErrCode, retval.Message)
		}
		if emailVerif.Type == "R" {
			SendOtp(emailVerif.Email, emailVerif.AuthCode)
		} else {
			SendOtp(emailVerif.Email, emailVerif.VerifCode)
		}
		return utils.Message(true)
	}
}

func (emailVerifController *EmailVerifController) AuthEmail(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(emailVerifController.Model()); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	db := utils.GetDB()
	emailVerif := &emailVerifController.EmailVerif
	row := db.Where("UPPER(email) = UPPER(?) AND UPPER(type) = UPPER(?) AND (verified = TRUE OR type = 'R')", emailVerif.Email, emailVerif.Type).Find(emailVerif).RowsAffected
	if row > 0 {
		utils.Respond(w, emailVerifController.UpdateModel())
		return
	} else if emailVerif.Type == "R" {
		utils.Respond(w, emailVerifController.CreateModel())
		return
	}
	utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "This email not yet verified!"))
}

func (emailVerifController *EmailVerifController) ViewModel(id int64) map[string]interface{} {
	emailVerif := &model.EmailVerif{}
	if retval := ViewModel(id, emailVerif); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, emailVerif)
}

func (emailVerifController *EmailVerifController) UpdateModel() map[string]interface{} {
	b, delayedTime := emailVerifController.EmailVerif.ValidateTime()
	if b {
		retval, m := UpdateModel(emailVerifController, &model.EmailVerif{}, func(modelSrc, modelTemp model.Model) {
			emailVerif := modelSrc.(*model.EmailVerif)
			emailVerif.GenerateCode(emailVerifController.Mode)
		})
		if retval.ErrCode != 0 {
			return utils.MessageErr(false, retval.ErrCode, retval.Message)
		}
		emailVerif := m.(*model.EmailVerif)
		if emailVerifController.Mode == model.EMAIL_VERIF_REGISTER {
			SendOtp(emailVerif.Email, emailVerif.VerifCode)
		} else if emailVerifController.Mode == model.EMAIL_VERIF_AUTH {
			SendOtp(emailVerif.Email, emailVerif.AuthCode)
		}
		return utils.Message(true)
	} else {
		return utils.MessageErr(false, utils.ErrValidate, fmt.Sprintf("try again after %s", delayedTime))
	}
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
	return ListModel("emailVerif", "id ASC", &emailVerifController.EmailVerif, make([]*model.EmailVerif, 0), param)
}

func (emailVerifController *EmailVerifController) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	type Email struct {
		Email string
		Code  string
		Type  string
	}
	email := &Email{}
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}

	db := utils.GetDB().Begin()

	emailData := &model.EmailVerif{}
	row := db.Where("UPPER(email) = UPPER(?) AND UPPER(type) = UPPER(?)", email.Email, email.Type).Find(&emailData).RowsAffected
	if row > 0 {
		if emailVerifController.Mode == model.EMAIL_VERIF_REGISTER {
			if emailData.Verified {
				utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "This email has already verified!"))
				return
			}
			if email.Code == emailData.VerifCode {
				emailData.Verified = true
				emailData.VerifiedAt = time.Now()
				emailData.WaitTime = 0
				if err := model.Save(emailData, db); err != nil {
					db.Rollback()
					utils.Respond(w, utils.MessageErr(false, utils.ErrSQLCreate, err.Error()))
					return
				}
				utils.Respond(w, utils.Message(true))
			} else {
				utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "Wrong verification code!"))
			}
		} else if emailVerifController.Mode == model.EMAIL_VERIF_AUTH {
			if email.Code == emailData.AuthCode {
				emailData.WaitTime = 0
				db.Save(emailData)
				utils.Respond(w, utils.Message(true))
			} else {
				utils.Respond(w, utils.MessageErr(false, utils.ErrRequest, "Wrong authentication code!"))
			}
		}
	} else {
		utils.Respond(w, utils.MessageErr(false, utils.ErrExist, "Email not yet registered!"))
	}
	db.Commit()
}

func SendOtp(to, code string) {
	templateData := struct {
		Code string
	}{
		code,
	}
	runtime.GOMAXPROCS(1)
	go utils.SendEmailOtp(to, templateData)
}
