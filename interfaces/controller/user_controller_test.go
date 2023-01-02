package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestUserController(t *testing.T) {
	jsonHeander := map[string]string{"Content-type": "application/json"}

	data := []struct {
		name    string
		method  string
		url     string
		payload []byte
		headers map[string]string
	}{
		{
			name:    "Create User",
			method:  "POST",
			url:     "http://localhost:4200/user/create",
			payload: []byte(`{"name": "Test", "email":"test@gmail.com", "password":"test"}`),
			headers: jsonHeander,
		},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			var payload io.Reader
			if v.payload != nil {
				payload = bytes.NewBuffer(v.payload)
			}

			req, err := http.NewRequest(v.method, v.url, payload)
			if err != nil {
				t.Fatalf("Failed could not create request, got: %v\n", err)
			}

			for key, value := range v.headers {
				req.Header.Set(key, value)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed could not send request, got: %v\n", err)
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("Failed expected code 200, got: %v\n", res.Status)
			}
		})
	}
}
