package driver

import "database/sql"

// Драйвер Базы Данных
type DB struct {
	SQL *sql.DB
}

var db = &DB{}

/*
Функция подключения к БД.
dbdriver - драйвер БД
host - хост сервера
port - порт сервера
dbname - название БД
*/

func Connect(dbdriver, host, port, user, password, dbname string) *DB {
	return db
}
