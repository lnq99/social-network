package repository

import (
	"database/sql"

	"app/internal/model"
)

type albumRepoImpl struct {
	db *sql.DB
}

// Функция создания нового альбома для пользователя
func NewAlbumRepo(db *sql.DB) AlbumRepo {
	return &albumRepoImpl{db}
}

// Функция считывания альбома
func scanAlbum(row MultiScanner, c *model.Album) error {
	err := row.Scan(
		&c.Id,
		&c.UserId,
		&c.Descr,
		&c.Created,
	)
	return err
}

// Функция считывания информации о альмобе по ID пользователя
func (r albumRepoImpl) selectById(query string, id int) (res []model.Album, err error) {
	rows, err := r.db.Query(query, id)
	if err != nil {
		return
	}

	var rel model.Album

	defer rows.Close()
	for rows.Next() {
		err = scanAlbum(rows, &rel)
		if err != nil {
			return
		}
		res = append(res, rel)
	}

	return
}

// Функция добавления новых данных в альбом пользователя
func (r albumRepoImpl) Insert(a *model.Album) (id int, err error) {
	query := `insert into Album(id, userId, desc, created) values ($1, $2, $3, $4)`
	row := r.db.QueryRow(query, a.Id, a.UserId, a.Descr, a.Created)
	err = row.Scan(&id)
	return
}

// Функция выбора альбома
func (r albumRepoImpl) Select(id int) (album model.Album, err error) {
	row := r.db.QueryRow("select * from Album where id=$1 limit 1", id)
	err = scanAlbum(row, &album)
	return
}

// Функция выбора альбома по usersID
func (r albumRepoImpl) SelectByUserId(userId int) ([]model.Album, error) {
	return r.selectById(`select * from Album where UserId=$1`, userId)
}
