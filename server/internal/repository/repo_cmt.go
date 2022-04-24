package repository

import (
	"database/sql"

	"app/internal/model"
)

type commentRepoImpl struct {
	db *sql.DB
}

// Функция создания репозитория Комментарий
func NewCommentRepo(db *sql.DB) CommentRepo {
	return &commentRepoImpl{db}
}

// Функция считывания комментария
func scanComment(row MultiScanner, c *model.Comment) error {
	err := row.Scan(
		&c.Id,
		&c.UserId,
		&c.PostId,
		&c.ParentId,
		&c.Content,
		&c.Created,
	)
	return err
}

// Функция добавления нового комментария
func (r commentRepoImpl) Insert(cmt *model.Comment) (id int, err error) {
	query := `insert into Comment(userId, postId, parentId, content) values ($1, $2, $3, $4)`
	row := r.db.QueryRow(query, cmt.UserId, cmt.PostId, cmt.ParentId, cmt.Content)
	err = row.Scan(&id)
	return
}

// Функция выбора комментария
func (r commentRepoImpl) Select(postId int) (res []model.Comment, err error) {
	rows, err := r.db.Query(`select * from Comment where postId=$1`, postId)
	if err != nil {
		return
	}

	var cmt model.Comment

	defer rows.Close()
	for rows.Next() {
		err = scanComment(rows, &cmt)
		if err != nil {
			return
		}
		res = append(res, cmt)
	}

	return
}
