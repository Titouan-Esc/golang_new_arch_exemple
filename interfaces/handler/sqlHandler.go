package handler

type SQLHandler interface {
	Query(string, ...interface{}) (Row, error)
}

type Row interface {
	Scan(...interface{}) error
	Close() error
	Next() bool
}
