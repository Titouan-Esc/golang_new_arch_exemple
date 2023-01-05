package repository

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"exemple.com/swagTest/utils"
	"strings"
	"testing"
)

func TestUserRepository(t *testing.T) {
	userId := utils.GenerateId()
	suite := test.Setup()
	repo := NewUserRepository(suite.Handler)

	suite.Data = []test.Data{
		{
			Name: "Create",
			Body: map[string]interface{}{
				"Id":       userId,
				"Name":     "Jean",
				"Email":    "jean@gmail.com",
				"Password": middlewares.HasPassword("jean"),
			},
		},
		{
			Name: "FindEmail",
			Body: map[string]interface{}{
				"Email": "jean@gmail.com",
			},
		},
		{
			Name: "Find",
			Body: map[string]interface{}{
				"Id": userId,
			},
		},
		{
			Name: "Update",
			Body: map[string]interface{}{
				"Id":   userId,
				"Name": "Diego",
			},
		},
		{
			Name: "Delete",
			Body: map[string]interface{}{
				"Id": userId,
			},
		},
	}

	for _, value := range suite.Data {
		var resp model.User
		var err error
		var user model.User
		mapByte, _ := json.Marshal(value.Body)
		json.Unmarshal(mapByte, &user)

		t.Run(value.Name, func(t *testing.T) {
			switch strings.ToUpper(value.Name) {
			case "CREATE":
				resp, err = repo.Create(user)
			case "FINDEMAIL":
				resp, err = repo.FindByEmail(user.Email)
			case "FIND":
				resp, err = repo.Find(user.ID)
			case "UPDATE":
				resp, err = repo.Update(user)
			case "DELETE":
				resp, err = repo.Delete(user)
			}

			if err != nil {
				t.Errorf("Failed to query, got: %v\n", err)
			}

			if resp.ID != userId {
				t.Error("Failed response query is empty")
			}

		})
	}
}
