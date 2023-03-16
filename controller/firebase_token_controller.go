package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateFirebaseToken = func(w http.ResponseWriter, r *http.Request) {
	firebaseTokenController := &FirebaseTokenController{}
	CreateModelAction(firebaseTokenController, w, r)
}

var UpdateFirebaseToken = func(w http.ResponseWriter, r *http.Request) {
	firebaseTokenController := &FirebaseTokenController{}
	UpdateModelAction(firebaseTokenController, w, r)
}

var ViewFirebaseToken = func(w http.ResponseWriter, r *http.Request) {
	firebaseTokenController := &FirebaseTokenController{}
	ViewModelAction(firebaseTokenController, w, r)
}

var ListFirebaseToken = func(w http.ResponseWriter, r *http.Request) {
	firebaseTokenController := &FirebaseTokenController{}
	ListModelAction(firebaseTokenController, w, r)
}

type FirebaseTokenController struct {
	FirebaseToken model.FirebaseToken
}

func (firebaseTokenController *FirebaseTokenController) Model() model.Model {
	return &firebaseTokenController.FirebaseToken
}

func (firebaseTokenController *FirebaseTokenController) FNew() functions.SQLFunction {
	return nil
}
func (firebaseTokenController *FirebaseTokenController) FDelete() functions.SQLFunction {
	return nil
}

func (firebaseTokenController *FirebaseTokenController) CreateModel() map[string]interface{} {
	isTokenExist, firebaseToken := TokenExist(firebaseTokenController.FirebaseToken.Token)
	if !isTokenExist {
		if retval := CreateModel(firebaseTokenController, func(m model.Model) {
			currentTime := time.Now()

			firebaseToken := m.(*model.FirebaseToken)
			firebaseToken.CreatedAt = currentTime
			firebaseToken.UpdatedAt = currentTime
		}); retval.ErrCode != 0 {
			return utils.MessageErr(false, retval.ErrCode, retval.Message)
		}
	} else {
		return utils.MessageData(true, firebaseToken)
	}
	return utils.MessageData(true, firebaseTokenController.FirebaseToken)
}

func (firebaseTokenController *FirebaseTokenController) ViewModel(id int64) map[string]interface{} {
	firebaseToken := &model.FirebaseToken{}
	if retval := ViewModel(id, firebaseToken); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, firebaseToken)
}

func (firebaseTokenController *FirebaseTokenController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(firebaseTokenController, &model.FirebaseToken{}, func(modelSrc, modelTemp model.Model) {
		firebaseTokenSrc := modelSrc.(*model.FirebaseToken)
		firebaseTokenTemp := modelTemp.(*model.FirebaseToken)

		firebaseTokenSrc.Uid = firebaseTokenTemp.Uid
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (firebaseTokenController *FirebaseTokenController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, firebaseTokenController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (firebaseTokenController *FirebaseTokenController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("firebase_token", "id ASC", &firebaseTokenController.FirebaseToken, make([]*model.FirebaseToken, 0), param)
}

func TokenExist(token string) (bool, *model.FirebaseToken) {
	db := utils.GetDB()

	firebaseToken := &model.FirebaseToken{}
	rows := db.Where("token = ?", token).Find(&firebaseToken).RowsAffected
	return rows > 0, firebaseToken
}
