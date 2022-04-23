package repository

import (
	"database/sql"

	"app/internal/model"

	"github.com/lib/pq"
)

type postRepoImpl struct {
	db *sql.DB
}

func (r postRepoImpl) Update(p *model.Post) (err error) {
	query := `update Post set Tags=$2, Content=$3, AtchType=$4, AtchId=$5, AtchUrl=$6 where id=$1`
	res, err := r.db.Exec(query, p.Id, p.Tags, p.Content, p.AtchType, p.AtchId, p.AtchUrl)
	if err == nil {
		return handleRowsAffected(res)
	}
	return
}

func (r postRepoImpl) Delete(userId, postId int) (err error) {
	query := `delete from Post where id=$2 and userId=$1`
	res, err := r.db.Exec(query, userId, postId)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

func NewPostRepo(db *sql.DB) PostRepo {
	return &postRepoImpl{db}
}

func scanPost(row MultiScanner, p *model.Post) error {
	var arr pq.Int64Array
	err := row.Scan(
		&p.Id,
		&p.UserId,
		&p.Created,
		&p.Tags,
		&p.Content,
		&p.AtchType,
		&p.AtchId,
		&p.AtchUrl,
		&arr,
		&p.CmtCount,
	)
	p.Reaction = arr
	return err
}

func (r postRepoImpl) Insert(p *model.Post) (id int, err error) {
	query := `insert into Post(userId, tags, content, atchType, atchId, atchUrl)
		values ($1, $2, $3, $4, $5, $6) returning id`
	row := r.db.QueryRow(query, p.UserId, p.Tags, p.Content, p.AtchType, p.AtchId, p.AtchUrl)
	err = row.Scan(&id)
	return
}

func (r postRepoImpl) Select(postId int) (post model.Post, err error) {
	row := r.db.QueryRow("select * from Post where id=$1 limit 1", postId)
	err = scanPost(row, &post)
	return
}

func (r postRepoImpl) SelectByUserId(userId int) (posts []int64, err error) {
	row := r.db.QueryRow("select array(select id from Post where userId=$1 order by created desc)", userId)
	var arr pq.Int64Array
	err = row.Scan(&arr)
	posts = arr
	return
}

func (r postRepoImpl) SelectReaction(postId int) (res []int64, err error) {
	row := r.db.QueryRow("select reaction from Post where id=$1", postId)
	var arr pq.Int64Array
	err = row.Scan(&arr)
	res = arr
	return
}
