package repository

import "exemple.com/swagTest/domain/model"

type UserRepository interface {
	Find(uid string) (model.Model, error)
	FindByEmail(mail string) (model.Model, error)
	Create(model model.Model) (string, error)
	Update(model model.Model) (string, error)
	Delete(uid string) (string, error)
	Login(user model.Model) (string, error)
}
