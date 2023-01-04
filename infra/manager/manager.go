package manager

import (
	"encoding/json"
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"reflect"
)

type Manager[T model.User | map[string]interface{}] struct {
	Body        T
	Errors      errors.ErrorsEntity
	Response    http.ResponseWriter
	ResponsImpl *Respons
}

type Respons struct {
	Res  http.ResponseWriter
	Body ResponsBody
}

type ResponsBody struct {
	Success bool
	Status  int
	Data    []interface{}
}

func NewManager[T model.User | map[string]interface{}](res http.ResponseWriter, req *http.Request) *Manager[T] {
	var data T
	json.NewDecoder(req.Body).Decode(&data)

	var newData interface{}
	dataByte, _ := json.Marshal(data)
	json.Unmarshal(dataByte, &newData)

	return &Manager[T]{
		Body:     data,
		Response: res,
	}
}

func (m *Manager[T]) StopRequest() {
	m.Response.Header().Set("Content-Type", "application/json")
	m.Response.WriteHeader(http.StatusBadRequest)

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

func (m *Manager[T]) Bind(target, patch interface{}) {
	var x interface{}
	patchByte, _ := json.Marshal(patch)
	json.Unmarshal(patchByte, &x)

	emet := reflect.ValueOf(x)
	iter := emet.MapRange()

	f := reflect.ValueOf(target).Elem()

	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		caser := cases.Title(language.English)
		maj := caser.String(k.String())
		r := f.FieldByName(maj)

		if !r.IsValid() || !r.CanSet() {
			continue
		}

		if r.Kind() == reflect.Interface {
			continue
		}

		switch r.Kind() {
		case reflect.String:
			if v.Elem().String() == r.String() {
				continue
			}

			if v.Elem().String() == "" {
				continue
			}
			r.Set(v.Elem())
		}
	}
}

func (r *Respons) Build(status int, data ...interface{}) {
	r.Res.Header().Set("Content-Type", "application/json")
	r.Res.WriteHeader(status)

	r.Body.Success = true
	if status >= 400 {
		r.Body.Success = false
	}
	r.Body.Status = status
	r.Body.Data = data

	json.NewEncoder(r.Res).Encode(r.Body)
}
