package repository

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/interfaces/handler"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/utils"
)

type UserRepository struct {
	SQLHandler handler.SQLHandler
}

func (u UserRepository) Login(user model.Model) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Find(uid string) (model.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Create(model model.Model) (uid string, err error) {
	id := utils.GenerateId()
	password := middlewares.HasPassword(model.Password)

	rows, err := u.SQLHandler.Query(`INSERT INTO "user" ("id", "name", "email", "password") VALUES ($1, $2, $3, $4) RETURNING "id"`,
		id, model.Name, model.Email, password)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&uid); err != nil {
			return
		}
	}

	return uid, nil
}

func (u UserRepository) Update(model model.Model) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(uid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
