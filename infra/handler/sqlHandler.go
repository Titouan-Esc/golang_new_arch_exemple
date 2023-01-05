package handler

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLHandler struct {
	Db *gorm.DB
}

func NewSQLHandler() (SQLHandler, error) {
	var sqlHandler SQLHandler

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"test",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return sqlHandler, err
	}

	sqlHandler.Db = db

	return sqlHandler, nil
}
