package main

import (
	"exemple.com/swagTest/infra/handler"
	"exemple.com/swagTest/infra/router"
	"log"
	"net/http"
)

func main() {
	sqlHandler, err := handler.NewSQLHandler()
	if err != nil {
		log.Println(err.Error())
		return
	}

	r := router.Dispatch(sqlHandler)

	if err := http.ListenAndServe(":4200", r.Handle); err != nil {
		log.Println(err.Error())
		return
	}
}
