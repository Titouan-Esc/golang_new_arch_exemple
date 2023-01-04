package interactor

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"strings"
	"testing"
)

func TestUserInteractor_Store(t *testing.T) {
	suite := test.Setup()
	ui := NewUserInteractor(suite.Handler)

	suite.Data = []test.Data{
		{
			Name: "Store",
			Body: map[string]interface{}{
				"Name":     "Titouan",
				"Email":    "titouan@gmail.com",
				"Password": middlewares.HasPassword("titouan"),
			},
		},
		{
			Name: "FindEmail",
			Body: map[string]interface{}{
				"Email": "oscar@gmail.com",
			},
		},
	}

	for _, value := range suite.Data {

		var user model.User
		mapByte, _ := json.Marshal(value.Body)
		json.Unmarshal(mapByte, &user)

		t.Run(value.Name, func(t *testing.T) {

			switch strings.ToUpper(value.Name) {
			case "STORE":
				resp, err := ui.Store(user)
				if err != nil {
					t.Errorf("Failed to insert user in db, got: %v\n", err)
				}

				if resp.ID == "" {
					t.Error("Failed user id is empty")
				}
			case "FINDEMAIL":
				resp, err := ui.ShowByEmail(user.Email)
				if err != nil {
					t.Errorf("Failed to find user by email, got: %v\n", err)
				}

				if resp.ID == "" {
					t.Error("Failed user not found")
				}
			}
		})
	}
}
