package controller

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/adapter"
	"exemple.com/swagTest/interfaces/handler"
	"exemple.com/swagTest/interfaces/repository"
	"exemple.com/swagTest/usecase/interactor"
	"log"
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
	manager := adapter.NewAdapter[model.Model](res, req)
	if manager.Errors.Error {
		manager.StopRequest()
		return
	}

	log.Println(manager.Body)
}
