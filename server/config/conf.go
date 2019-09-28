package config

type Config struct {
	Port string
	Db   Database
}

type Database struct {
	port     string
	passwd   string
	username string
	host     string
	dbName   string
}
