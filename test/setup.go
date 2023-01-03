package test

import (
	handler2 "exemple.com/swagTest/infra/handler"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	Handler handler2.SQLHandler
	Data    []Data
}

type Data struct {
	Name string
	Body map[string]interface{}
}

func Setup() *Suite {
	var (
		handler handler2.SQLHandler
		s       *Suite
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"test",
	)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	handler = handler2.SQLHandler{Db: db}

	s = &Suite{
		Handler: handler,
	}

	return s
}
