package handler

import (
	"exemple.com/swagTest/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLHandler struct {
	Db *gorm.DB
}

func NewSQLHandler() (SQLHandler, error) {
	var sqlHandler SQLHandler

	data, _ := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		data.Db.Host,
		data.Db.Port,
		data.Db.User,
		data.Db.Password,
		data.Db.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return sqlHandler, err
	}

	sqlHandler.Db = db

	return sqlHandler, nil
}
