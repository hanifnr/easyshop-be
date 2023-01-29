package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"
	"time"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	userController := &UserController{}
	CreateModelAction(userController, w, r)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	userController := &UserController{}
	UpdateModelAction(userController, w, r)
}

var ViewUser = func(w http.ResponseWriter, r *http.Request) {
	userController := &UserController{}
	ViewModelAction(userController, w, r)
}

var ListUser = func(w http.ResponseWriter, r *http.Request) {
	userController := &UserController{}
	ListModelAction(userController, w, r)
}

type UserController struct {
	User model.User
}

func (userController *UserController) Model() model.Model {
	return &userController.User
}

func (userController *UserController) FNew() functions.SQLFunction {
	return nil
}
func (userController *UserController) FDelete() functions.SQLFunction {
	return nil
}

func (userController *UserController) CreateModel() map[string]interface{} {

	if retval := CreateModel(userController, func(m model.Model) {
		currentTime := time.Now()

		user := m.(*model.User)
		user.CreatedAt = currentTime
		user.UpdatedAt = currentTime
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, userController.User)
}

func (userController *UserController) ViewModel(id int64) map[string]interface{} {
	user := &model.User{}
	if retval := ViewModel(id, user); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, user)
}

func (userController *UserController) UpdateModel() map[string]interface{} {
	retval, retModel := UpdateModel(userController, &model.User{}, func(modelSrc, modelTemp model.Model) {
		userSrc := modelSrc.(*model.User)
		userTemp := modelTemp.(*model.User)

		userSrc.Name = userTemp.Name
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, retModel)
}

func (userController *UserController) DeleteModel(id int64) map[string]interface{} {
	if retval := DeleteModel(id, userController, func(m model.Model) utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}
func (userController *UserController) ListModel(param *utils.Param) map[string]interface{} {
	return ListModel("user", "id ASC", &userController.User, make([]*model.User, 0), param)
}
