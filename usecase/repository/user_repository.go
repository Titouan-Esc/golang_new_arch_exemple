package repository

import "exemple.com/swagTest/domain/model"

type UserRepository interface {
	Find(uid string) (model.User, error)
	FindByEmail(mail string) (model.User, error)
	Create(model model.User) (string, error)
	Update(model model.User) (string, error)
	Delete(uid string) (string, error)
	Login(user model.User) (string, error)
}
