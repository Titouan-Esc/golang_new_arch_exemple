package repository

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/handler"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/utils"
)

type UserRepository struct {
	SQLHandler handler.SQLHandler
}

func (u UserRepository) FindByEmail(mail string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Login(user model.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Find(uid string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Create(model model.User) (uid string, err error) {
	model.ID = utils.GenerateId()
	model.Password = middlewares.HasPassword(model.Password)
	if retour := u.SQLHandler.Db.Table("users").Create(&model); retour.Error != nil {
		return "", retour.Error
	}

	return model.ID, nil
}

func (u UserRepository) Update(model model.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(uid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
