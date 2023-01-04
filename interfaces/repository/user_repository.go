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

func NewUserRepository(sql handler.SQLHandler) *UserRepository {
	return &UserRepository{
		SQLHandler: sql,
	}
}

func (u UserRepository) FindByEmail(mail string) (model.User, error) {
	var user model.User
	if retour := u.SQLHandler.Db.Table("users").Where("email = ?", mail).First(&user); retour.Error != nil {
		return model.User{}, retour.Error
	}
	return user, nil
}

func (u UserRepository) Find(uid string) (model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserRepository) Create(model model.User) (model.User, error) {
	model.ID = utils.GenerateId()
	model.Password = middlewares.HasPassword(model.Password)
	if retour := u.SQLHandler.Db.Table("users").Save(&model); retour.Error != nil {
		return model, retour.Error
	}
	return model, nil
}

func (u UserRepository) Update(model model.User) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(uid string) (string, error) {
	// TODO implement me
	panic("implement me")
}
