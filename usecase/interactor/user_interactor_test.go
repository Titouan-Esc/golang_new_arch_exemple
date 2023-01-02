package interactor

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"testing"
)

func TestUserInteractor_Store(t *testing.T) {
	suite := test.Setup[model.User]()
	ui := NewUserInteractor(suite.Handler)

	suite.Data = model.User{
		Name:     "Titouan",
		Email:    "titouan@gmaiL.com",
		Password: middlewares.HasPassword("titouan"),
	}

	resp, err := ui.Store(suite.Data)
	if err != nil {
		t.Errorf("Failed to insert user in db, got: %v\n", err)
	}

	if resp == "" {
		t.Error("Failed user id is empty")
	}
}
