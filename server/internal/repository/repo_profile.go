package repository

import (
	"database/sql"

	"app/internal/model"
)

type profileRepoImpl struct {
	db *sql.DB
}

func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &profileRepoImpl{db}
}

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

func (r profileRepoImpl) Insert(p *model.Profile) (id int, err error) {
	query := `insert into Profile(name, gender, birthdate, email, salt, hash)
		values ($1, $2, $3, $4, $5, $6) returning id`
	row := r.db.QueryRow(query, p.Name, p.Gender, p.Birthdate, p.Email, p.Salt, p.Hash)
	err = row.Scan(&id)
	return
}

func (r profileRepoImpl) Select(id int) (p model.Profile, err error) {
	row := r.db.QueryRow("select * from Profile where id=$1 limit 1", id)
	err = scanProfile(row, &p)
	return
}

func (r profileRepoImpl) SelectByEmail(email string) (p model.Profile, err error) {
	row := r.db.QueryRow("select * from Profile where email=$1 limit 1", email)
	err = scanProfile(row, &p)
	return
}

func (r profileRepoImpl) UpdateAvatar(photo model.Photo) (err error) {
	query := `update Profile set avartarL=$1 where id=$2`
	res, err := r.db.Exec(query, photo.Url, photo.UserId)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

func (r profileRepoImpl) UpdateIntro(id int, intro string) (err error) {
	query := `update Profile set intro=$1 where id=$2`
	res, err := r.db.Exec(query, intro, id)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}
