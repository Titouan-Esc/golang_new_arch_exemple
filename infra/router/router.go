package router

import (
	"exemple.com/swagTest/interfaces/controller"
	"exemple.com/swagTest/interfaces/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

type Router struct {
	Handle *chi.Mux
}

var router = Router{chi.NewRouter()}

func Dispatch(sqlHandler handler.SQLHandler) *Router {
	userController := controller.NewUserController(sqlHandler)
	router.Handle.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.AddRoute("POST", "/user/create", userController.Store)
	router.AddRoute("POST", "/user/login", userController.Connect)

	return &router
}

func (r *Router) AddRoute(action, url string, fonc http.HandlerFunc) {
	r.Handle.MethodFunc(action, url, fonc)
}
