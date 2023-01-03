package repository

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"strings"
	"testing"
)

func TestUserRepository(t *testing.T) {
	suite := test.Setup()
	repo := NewUserRepository(suite.Handler)

	suite.Data = []test.Data{
		{
			Name: "Create",
			Body: map[string]interface{}{
				"Name":     "Jean",
				"Email":    "jean@gmail.com",
				"Password": middlewares.HasPassword("jean"),
			},
		},
		{
			Name: "Login",
			Body: map[string]interface{}{
				"Email":    "jean@gmail.com",
				"Password": "jean",
			},
		},
	}

	for _, value := range suite.Data {

		var user model.User
		mapByte, _ := json.Marshal(value.Body)
		json.Unmarshal(mapByte, &user)

		t.Run(value.Name, func(t *testing.T) {
			switch strings.ToUpper(value.Name) {
			case "CREATE":
				resp, err := repo.Create(user)
				if err != nil {
					t.Errorf("Failed to insert in db, got: %v\n", err)
				}

				if resp == "" {
					t.Error("Failed user id is empty")
				}
			}
		})
	}
}
