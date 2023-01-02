package model

type Env struct {
	Db DB `json:"db"`
}

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	DBName   string `json:"dbname"`
}
