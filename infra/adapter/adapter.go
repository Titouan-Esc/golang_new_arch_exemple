package adapter

import (
	"exemple.com/swagTest/domain/model"
	"exemple.com/swagTest/infra/manager"
	"net/http"
)

func NewAdapter[T model.Model | map[string]interface{}](res http.ResponseWriter, req *http.Request, token ...bool) *manager.Manager[T] {
	var newManager *manager.Manager[T]
	newManager = manager.NewManager[T](res, req)

	testToken := true
	for _, v := range token {
		testToken = v
	}

	if testToken {
		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" {
			newManager.Errors.AddError("E", "NO_TOKEN", "Request doesn't contain key API")
			return newManager
		}

	}

	return newManager
}
