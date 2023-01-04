package controller

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/controller"
	"exemple.com/swagTest/infra/handler"
	"exemple.com/swagTest/interfaces/repository"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/usecase/interactor"
	"net/http"
)

type UserController struct {
	UserInteractor interactor.UserInteractor
}

func NewUserController(sqlHandler handler.SQLHandler) *UserController {
	return &UserController{
		UserInteractor: interactor.UserInteractor{
			UserRepository: &repository.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (uc *UserController) Store(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})
	manager := controller.NewController[model.User](res, req, false)
	if manager.Errors.Error {
		manager.StopRequest()
		return
	}

	user, err := uc.UserInteractor.Store(manager.Body)
	if err != nil {
		manager.Respons().Build(http.StatusBadRequest, err.Error())
		return
	}

	token, err := middlewares.GenerateJWT(user.Email)
	if err != nil {
		manager.Respons().Build(http.StatusBadRequest, err.Error())
	}

	response["id"] = user.ID
	response["name"] = user.Name
	response["email"] = user.Email
	response["token"] = token

	manager.Respons().Build(http.StatusOK, response)
}

func (uc *UserController) Connect(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})

	// Init controller
	manager := controller.NewController[model.User](res, req, false)
	if manager.Errors.Error {
		manager.StopRequest()
		return
	}

	// Get User with his email
	user, err := uc.UserInteractor.ShowByEmail(manager.Body.Email)
	if err != nil {
		manager.Respons().Build(http.StatusForbidden, err.Error())
		return
	}

	// Check password
	if ok := middlewares.ValidateEncrypt(manager.Body.Password, user.Password); !ok {
		manager.Respons().Build(http.StatusBadRequest, "Bad Password")
		return
	}

	// Generate token
	token, err := middlewares.GenerateJWT(manager.Body.Email)
	if err != nil {
		manager.Respons().Build(http.StatusConflict, err.Error())
		return
	}

	response["token"] = token
	manager.Respons().Build(http.StatusOK, response)
}

func (uc *UserController) Show(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})

	manager := controller.NewController[model.User](res, req)
	if manager.Errors.Error {
		manager.StopRequest()
		return
	}

	user, err := uc.UserInteractor.Show(manager.Body.ID)
	if err != nil {
		manager.Respons().Build(http.StatusBadRequest, err.Error())
		return
	}

	response["id"] = user.ID
	response["name"] = user.Name
	response["email"] = user.Email

	manager.Respons().Build(http.StatusOK, response)
}

func (uc *UserController) Modify(res http.ResponseWriter, req *http.Request) {

	manager := controller.NewController[model.User](res, req)
	if manager.Errors.Error {
		manager.StopRequest()
		return
	}

	user, err := uc.UserInteractor.Show(manager.Body.ID)
	if err != nil {
		manager.Respons().Build(http.StatusBadRequest, err.Error())
		return
	}

	manager.Bind(&user, manager.Body)
	uid, err := uc.UserInteractor.Modify(user)
	if err != nil {
		manager.Respons().Build(http.StatusBadRequest, err.Error())
		return
	}

	manager.Respons().Build(http.StatusOK, uid)
}
