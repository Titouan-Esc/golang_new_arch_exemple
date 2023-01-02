package repository

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	suite := test.Setup[model.User]()
	repo := NewUserRepository(suite.Handler)

	suite.Data = model.User{
		Name:     "Jean",
		Email:    "jean@gmail.com",
		Password: middlewares.HasPassword("jean"),
	}

	resp, err := repo.Create(suite.Data)
	if err != nil {
		t.Errorf("Failed to insert in db, got: %v\n", err)
	}

	if resp == "" {
		t.Error("Failed user id is empty")
	}
}
