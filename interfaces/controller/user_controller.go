package controller

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/controller"
	"exemple.com/swagTest/interfaces/handler"
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
	var user model.Model
	json.NewDecoder(req.Body).Decode(&user)

	uid, err := uc.UserInteractor.Store(user)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(uid)
}

func (uc *UserController) Connect(res http.ResponseWriter, req *http.Request) {
	// Init controller
	manager := controller.NewController[model.Model](res, req, false)
	if manager.Errors.Error {
		manager.StopRequest(http.StatusForbidden)
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
	manager.Respons().Build(http.StatusOK, token)
}
