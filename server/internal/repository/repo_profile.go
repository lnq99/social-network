package repository

import (
	"database/sql"

	"app/internal/model"

	"github.com/lib/pq"
)

type profileRepoImpl struct {
	db *sql.DB
}

// Функция создания репозитория Профиля
func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &profileRepoImpl{db}
}

// Функция сканирования профиля
func scanProfile(row MultiScanner, p *model.Profile) error {
	return row.Scan(
		&p.Id,
		&p.Name,
		&p.Gender,
		&p.Birthdate,
		&p.Email,
		&p.Phone,
		&p.Salt,
		&p.Hash,
		&p.Created,
		&p.Intro,
		&p.AvatarS,
		&p.AvatarL,
		&p.PostCount,
		&p.PhotoCount,
	)
}

// Функция добавления информации о пользователе
func (r profileRepoImpl) Insert(p *model.Profile) (id int, err error) {
	query := `insert into Profile(name, gender, birthdate, email, salt, hash)
		values ($1, $2, $3, $4, $5, $6) returning id`
	row := r.db.QueryRow(query, p.Name, p.Gender, p.Birthdate, p.Email, p.Salt, p.Hash)
	err = row.Scan(&id)
	return
}

// Функция получения информации о пользователе
func (r profileRepoImpl) Select(id int) (p model.Profile, err error) {
	row := r.db.QueryRow("select * from Profile where id=$1 limit 1", id)
	err = scanProfile(row, &p)
	return
}

// Функция получения информации о пользователе по Email
func (r profileRepoImpl) SelectByEmail(email string) (p model.Profile, err error) {
	row := r.db.QueryRow("select * from Profile where email=$1 limit 1", email)
	err = scanProfile(row, &p)
	return
}

// Функция изменения аватаркипользователя
func (r profileRepoImpl) UpdateAvatar(photo model.Photo) (err error) {
	query := `update Profile set avartarL=$1 where id=$2`
	res, err := r.db.Exec(query, photo.Url, photo.UserId)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

// Функция обновления "О себе"
func (r profileRepoImpl) UpdateIntro(id int, intro string) (err error) {
	query := `update Profile set intro=$1 where id=$2`
	res, err := r.db.Exec(query, intro, id)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

// Функция поиска имения в БД
func (r profileRepoImpl) SearchName(id int, s string) (res string, err error) {
	if len(s) >= 2 {
		err = r.db.QueryRow("select search_name($1, $2)", id, s).Scan(&res)
	}
	if err != nil {
		err = nil
		res = "[]"
	}
	return
}

func (r profileRepoImpl) SelectFeed(id, limit, offset int) (feed []int64, err error) {
	row := r.db.QueryRow("select feed($1, $2, $3)", id, limit, offset)
	var arr pq.Int64Array
	err = row.Scan(&arr)
	feed = arr
	return
}
