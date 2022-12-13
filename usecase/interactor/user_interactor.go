package interactor

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (ui *UserInteractor) Store(user model.Model) (id string, err error) {
	id, err = ui.UserRepository.Create(user)
	return
}

func (ui *UserInteractor) Show(uid string) (user model.Model, err error) {
	user, err = ui.UserRepository.Find(uid)
	return
}

func (ui *UserInteractor) Modify(user model.Model) (id string, err error) {
	id, err = ui.UserRepository.Update(user)
	return
}

func (ui *UserInteractor) Destroy(uid string) (id string, err error) {
	id, err = ui.UserRepository.Delete(uid)
	return
}

func (ui *UserInteractor) Connect(user model.Model) (token string, err error) {
	token, err = ui.UserRepository.Login(user)
	return
}