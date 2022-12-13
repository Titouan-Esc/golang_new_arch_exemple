package handler

import (
	"database/sql"
	"exemple.com/swagTest/interfaces/handler"
	_ "github.com/lib/pq"
)

type SQLHandler struct {
	Conn *sql.DB
}

type Rows struct {
	Rows *sql.Rows
}

func NewSQLHandler() (handler.SQLHandler, error) {
	sqlHandler := &SQLHandler{}

	dsn := "host=localhost port=5432 password=postgres dbname=test sslmode=disable search_path=public"

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	sqlHandler.Conn = conn

	return sqlHandler, err
}

func (S SQLHandler) Query(query string, args ...interface{}) (handler.Row, error) {
	rows, err := S.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	row := &Rows{
		Rows: rows,
	}

	return row, err
}

func (r Rows) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

func (r Rows) Close() error {
	return r.Rows.Close()
}

func (r Rows) Next() bool {
	return r.Rows.Next()
}