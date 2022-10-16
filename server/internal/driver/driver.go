package driver

import (
	"database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
)

// Драйвер Базы Данных
type DB struct {
	SQL *sql.DB
}

var db = &DB{}

const (
	postgresDSFmt = "host=%s port=%s user=%s password=%s dbname='%s' sslmode=disable"
	mysqlDSFmt    = "%s:%s@tcp(%s:%s)/%s"
)

/*
Функция подключения к БД.
dbdriver - драйвер БД
host - хост сервера
port - порт сервера
dbname - название БД
*/

func Connect(dbdriver, host, port, user, password, dbname string) *DB {
	var dataSourceName string

	if dbdriver == "postgres" {
		dataSourceName = fmt.Sprintf(postgresDSFmt, host, port, user, password, dbname)
	} else if dbdriver == "mysql" {
		dataSourceName = fmt.Sprintf(mysqlDSFmt, user, password, host, port, dbname)
	} else {
		panic("Driver not supported!")
	}

	pool, err := sql.Open(dbdriver, dataSourceName)
	if err != nil {
		panic(err)
	}

	db.SQL = pool
	return db
}
