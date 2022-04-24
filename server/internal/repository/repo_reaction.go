package repository

import (
	"database/sql"

	"app/internal/model"
)

type reactionRepoImpl struct {
	db *sql.DB
}

// Функция создания репозитория реакции
func NewReactionRepo(db *sql.DB) ReactionRepo {
	return &reactionRepoImpl{db}
}

// Функция считывания реакции пользователя
func scanReaction(row MultiScanner, c *model.Reaction) error {
	err := row.Scan(
		&c.UserId,
		&c.PostId,
		&c.T,
	)
	return err
}

// Функция добавления новых реакций
func (r reactionRepoImpl) InsertUpdate(userId, postId int, reaction string) error {
	query := `insert into Reaction values ($1, $2, $3)
		on conflict (userId, postId) do update set type = $3`

	_, err := r.db.Exec(query, userId, postId, reaction)
	return err
}

// Функция выбора реакции
func (r reactionRepoImpl) Select(postId int) (res []model.Reaction, err error) {
	rows, err := r.db.Query(`select * from Comment where postId=$1`, postId)
	if err != nil {
		return
	}

	var reaction model.Reaction

	defer rows.Close()
	for rows.Next() {
		err = scanReaction(rows, &reaction)
		if err != nil {
			return
		}
		res = append(res, reaction)
	}

	return
}

// Функция выбора реакций пользователя
func (r reactionRepoImpl) SelectReactionOfUser(userId, postId int) (t string, err error) {
	row := r.db.QueryRow("select type from Reaction where userId=$1 and postId=$2 limit 1", userId, postId)
	err = row.Scan(&t)
	if err != nil {
		err = nil
		t = ""
	}
	return
}
