package test

import (
	"exemple.com/swagTest/domain/model"
	handler2 "exemple.com/swagTest/infra/handler"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite[T model.User | map[string]interface{}] struct {
	Handler handler2.SQLHandler
	Data    T
}

func Setup[T model.User | map[string]interface{}]() *Suite[T] {
	var (
		handler handler2.SQLHandler
		s       *Suite[T]
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

	s = &Suite[T]{
		Handler: handler,
	}

	return s
}
