package repository

import "exemple.com/swagTest/domain/model"

type UserRepository interface {
	Find(uid string) (model.User, error)
	FindByEmail(mail string) (model.User, error)
	Create(model model.User) (model.User, error)
	Update(model model.User) (model.User, error)
	Delete(model model.User) (model.User, error)
}
