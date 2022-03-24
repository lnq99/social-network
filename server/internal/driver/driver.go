package driver

import "database/sql"

type DB struct {
	SQL *sql.DB
}

var db = &DB{}

func Connect(dbdriver, host, port, user, password, dbname string) *DB {
	return db
}
