package repository

// MultiScanner created to support sql.Row and sql.Rows
type MultiScanner interface {
	Scan(dest ...any) error
}
