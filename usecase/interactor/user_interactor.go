package interactor

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/handler"
	repository2 "exemple.com/swagTest/interfaces/repository"
	"exemple.com/swagTest/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func NewUserInteractor(handle handler.SQLHandler) *UserInteractor {
	return &UserInteractor{
		UserRepository: repository2.UserRepository{
			SQLHandler: handle,
		},
	}
}

func (ui *UserInteractor) Store(user model.User) (usr model.User, err error) {
	usr, err = ui.UserRepository.Create(user)
	return
}

func (ui *UserInteractor) Show(uid string) (user model.User, err error) {
	user, err = ui.UserRepository.Find(uid)
	return
}

func (ui *UserInteractor) Modify(user model.User) (usr model.User, err error) {
	usr, err = ui.UserRepository.Update(user)
	return
}

func (ui *UserInteractor) Destroy(user model.User) (usr model.User, err error) {
	usr, err = ui.UserRepository.Delete(user)
	return
}

func (ui *UserInteractor) ShowByEmail(mail string) (user model.User, err error) {
	user, err = ui.UserRepository.FindByEmail(mail)
	return
}
