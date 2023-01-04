package interactor

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/middlewares"
	"exemple.com/swagTest/test"
	"exemple.com/swagTest/utils"
	"strings"
	"testing"
)

func TestUserInteractor_Store(t *testing.T) {
	userId := utils.GenerateId()
	suite := test.Setup()
	ui := NewUserInteractor(suite.Handler)

	suite.Data = []test.Data{
		{
			Name: "Store",
			Body: map[string]interface{}{
				"Id":       userId,
				"Name":     "Titouan",
				"Email":    "titouan@gmail.com",
				"Password": middlewares.HasPassword("titouan"),
			},
		},
		{
			Name: "ShowEmail",
			Body: map[string]interface{}{
				"Email": "titouan@gmail.com",
			},
		},
		{
			Name: "Show",
			Body: map[string]interface{}{
				"id": userId,
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
			case "STORE":
				resp, err = ui.Store(user)
			case "SHOWEMAIL":
				resp, err = ui.ShowByEmail(user.Email)
			case "SHOW":
				resp, err = ui.Show(user.ID)
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
