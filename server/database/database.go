package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sqlx.DB
}

var Db Database

func CreateConnection() {
	source := fmt.Sprintf("%s:%s@(%s:%s)/%s", "root", "ggp0898", "localhost", "3306", "egressos_informatica")
	dbConn, err := sqlx.Connect("mysql", source)
	if err != nil {
		panic(err)
	}
	Db.Conn = dbConn
}
