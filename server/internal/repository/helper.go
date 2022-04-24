package repository

import (
	"database/sql"
	"errors"
)

// MultiScanner created to support sql.Row and sql.Rows
type MultiScanner interface {
	Scan(dest ...any) error
}

// Функция обработки ошибки БД в результате запросыа
func handleRowsAffected(res sql.Result) error {
	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("0 row affected")
	}
	return err
}
