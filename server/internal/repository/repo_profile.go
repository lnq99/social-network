package repository

import (
	"database/sql"
	"reflect"

	"app/internal/model"
	"app/pkg/logger"

	"github.com/lib/pq"
)

type ProfileRepoImpl struct {
	DB *sql.DB
}

func NewProfileRepo(db *sql.DB) ProfileRepo {
	return &ProfileRepoImpl{db}
}

func (r *ProfileRepoImpl) parseProfile(row *sql.Row) (profile model.Profile, err error) {
	s := reflect.ValueOf(&profile).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	err = row.Scan(columns...)

	return
}

func (r *ProfileRepoImpl) Select(id int) (model.Profile, error) {
	row := r.DB.QueryRow("select * from Profile where id=$1 limit 1", id)
	return r.parseProfile(row)
}

func (r *ProfileRepoImpl) SelectByEmail(e string) (model.Profile, error) {
	row := r.DB.QueryRow("select * from Profile where email=$1 limit 1", e)
	return r.parseProfile(row)
}

func (r *ProfileRepoImpl) SelectFeed(id, limit, offset int) (feed []int64, err error) {
	row := r.DB.QueryRow("select feed($1, $2, $3)", id, limit, offset)

	var arr pq.Int64Array
	err = row.Scan(&arr)
	feed = arr

	return
}

func (r *ProfileRepoImpl) SearchName(id int, s string) (res string, err error) {
	// log.Println(id, s)
	if len(s) >= 2 {
		err = r.DB.QueryRow("select search_name($1, $2)", id, s).Scan(&res)
	}
	if err != nil {
		// log.Println(err)
		err = nil
		res = "[]"
	}
	return
}

func (r *ProfileRepoImpl) Insert(p model.Profile) (err error) {
	query := `insert into Profile(name, gender, birthdate, email, salt, hash)
	values ($1, $2, $3, $4, $5, $6)`
	res, err := r.DB.Exec(query, p.Name, p.Gender, p.Birthdate, p.Email, p.Salt, p.Hash)
	if err == nil {
		return handleRowsAffected(res)
	}
	//logger.Err(err.Error())
	return
}

func (r *ProfileRepoImpl) SetAvatar(p model.Photo) (err error) {
	// avatarS not set for better performance
	query := `update Profile set avartarL=$1 where id=$2`
	res, err := r.DB.Exec(query, p.Url, p.UserId)
	if err == nil {
		return handleRowsAffected(res)
	}
	logger.Err(err.Error())
	return
}

func (r *ProfileRepoImpl) ChangeIntro(id int, intro string) (err error) {
	query := `update Profile set intro=$1 where id=$2`
	res, err := r.DB.Exec(query, intro, id)
	if err == nil {
		return handleRowsAffected(res)
	}
	logger.Err(err.Error())
	return
}
