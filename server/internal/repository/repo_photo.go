package repository

import (
	"database/sql"

	"app/internal/model"
)

type photoRepoImpl struct {
	db *sql.DB
}

// Функция создания репозитория Фото
func NewPhotoRepo(db *sql.DB) PhotoRepo {
	return &photoRepoImpl{db}
}

// Функция считывания фото пользователя
func scanPhoto(row MultiScanner, c *model.Photo) error {
	err := row.Scan(
		&c.Id,
		&c.UserId,
		&c.AlbumId,
		&c.Url,
		&c.Created,
	)
	return err
}

// Функция выбора изображения по USERID
func (r photoRepoImpl) selectById(query string, id int) (res []model.Photo, err error) {
	rows, err := r.db.Query(query, id)
	if err != nil {
		return
	}

	var rel model.Photo

	defer rows.Close()
	for rows.Next() {
		err = scanPhoto(rows, &rel)
		if err != nil {
			return
		}
		res = append(res, rel)
	}

	return
}

// Функция добавления нового фото
func (r photoRepoImpl) Insert(p *model.Photo) (id int, err error) {
	query := `insert into Photo(userId, albumId, url) values ($1, $2, $3) returning id`
	row := r.db.QueryRow(query, p.UserId, p.AlbumId, p.Url)
	err = row.Scan(&id)
	return
}

// Функция получения фото
func (r photoRepoImpl) Select(id int) (photo model.Photo, err error) {
	row := r.db.QueryRow("select * from Photo where id=$1 limit 1", id)
	err = scanPhoto(row, &photo)
	return
}

// Функция получения фото по userID
func (r photoRepoImpl) SelectByUserId(userId int) ([]model.Photo, error) {
	return r.selectById(`select * from Photo where UserId=$1`, userId)
}
