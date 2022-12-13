package manager

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/errors"
	"net/http"
)

type Manager[T model.Model | map[string]interface{}] struct {
	Body        T
	Errors      errors.ErrorsEntity
	Response    http.ResponseWriter
	ResponsImpl *Respons
}

type Respons struct {
	Res http.ResponseWriter
	ResponsBody
}

type ResponsBody struct {
	Success bool
	Status  int
	Data    []interface{}
}

func NewManager[T model.Model | map[string]interface{}](res http.ResponseWriter, req *http.Request) *Manager[T] {
	var data T
	json.NewDecoder(req.Body).Decode(&data)

	return &Manager[T]{
		Body:     data,
		Response: res,
	}
}

func (m *Manager[T]) StopRequest(status int) {
	m.Response.Header().Set("Content-Type", "application/json")
	m.Response.WriteHeader(status)

	json.NewEncoder(m.Response).Encode(m.Errors.Errors)
}

func (m *Manager[T]) Respons() *Respons {
	if m.ResponsImpl == nil {
		m.ResponsImpl = &Respons{
			Res: m.Response,
		}
	}

	return m.ResponsImpl
}

func (r *Respons) Build(status int, data ...interface{}) {
	r.Res.Header().Set("Content-Type", "application/json")
	r.Res.WriteHeader(status)

	r.Success = true
	if status >= 400 {
		r.Success = false
	}
	r.Status = status
	r.Data = data

	json.NewEncoder(r.Res).Encode(r.ResponsBody)
}
