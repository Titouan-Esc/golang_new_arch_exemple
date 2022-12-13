package manager

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/errors"
	"net/http"
)

type Manager[T model.Model | map[string]interface{}] struct {
	Body     T
	Errors   errors.ErrorsEntity
	Response http.ResponseWriter
}

func NewManager[T model.Model | map[string]interface{}](res http.ResponseWriter, req *http.Request) *Manager[T] {
	var data T
	json.NewDecoder(req.Body).Decode(&data)

	return &Manager[T]{
		Body:     data,
		Response: res,
	}
}

func (m *Manager[T]) StopRequest() {
	m.Response.Header().Set("Content-Type", "application/json")
	m.Response.WriteHeader(http.StatusForbidden)

	json.NewEncoder(m.Response).Encode(m.Errors.Errors)
}
