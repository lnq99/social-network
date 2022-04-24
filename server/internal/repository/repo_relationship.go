package repository

import (
	"database/sql"

	"app/internal/model"

	"github.com/lib/pq"
)

type relationshipRepoImpl struct {
	db *sql.DB
}

// Функция создания репозитория связей между пользователя
func NewRelationshipRepo(db *sql.DB) RelationshipRepo {
	return &relationshipRepoImpl{db}
}

// Функция считывания связи между пользователями
func scanRelationship(row MultiScanner, c *model.Relationship) error {
	err := row.Scan(
		&c.User1,
		&c.User2,
		&c.Created,
		&c.T,
		&c.Other,
	)
	return err
}

// Функция выбора связи между пользователями по Id
func (r relationshipRepoImpl) selectById(query string, id int) (res []model.Relationship, err error) {
	rows, err := r.db.Query(query, id)
	if err != nil {
		return
	}

	var rel model.Relationship

	defer rows.Close()
	for rows.Next() {
		err = scanRelationship(rows, &rel)
		if err != nil {
			return
		}
		res = append(res, rel)
	}

	return
}

// Функция выбора всех связей пользователя с другими пользователями
func (r relationshipRepoImpl) Select(userId int) ([]model.Relationship, error) {
	return r.selectById(`select * from relationship where user1=$1`, userId)
}

// Функция выбора всех "Друзей"
func (r relationshipRepoImpl) Friends(userId int) ([]model.Relationship, error) {
	return r.selectById(`select * from relationship where user1=$1 and type='friend'`, userId)
}

// Функция выбора всех запросов на "Дружбу"
func (r relationshipRepoImpl) Requests(userId int) ([]model.Relationship, error) {
	return r.selectById(`"select * from relationship where user1=$1 and type='request'`, userId)
}

// Функция получения информации о "Друзьях"
func (r relationshipRepoImpl) FriendsDetail(userId int) (fd string, err error) {
	err = r.db.QueryRow("select friends_json($1)", userId).Scan(&fd)
	return
}

func (r relationshipRepoImpl) MutualFriends(u1, u2 int) (mf []int64, err error) {
	row := r.db.QueryRow("select mutual_friends($1, $2)", u1, u2)
	var arr pq.Int64Array
	err = row.Scan(&arr)
	mf = arr
	return
}

// Функция получения связи между 2 пользователями
func (r relationshipRepoImpl) SelectRelationshipWith(u1, u2 int) (t string) {
	err := r.db.QueryRow("select type from relationship where user1=$1 and user2=$2", u1, u2).Scan(&t)
	if err != nil {
		return ""
	}
	return
}

// Функция измения отношения
func (r relationshipRepoImpl) ChangeType(u1, u2 int, t string) (err error) {
	query := `insert into relationship(user1, user2, type) values($1, $2, $3)
	on conflict (user1, user2) do update set type=$3`
	res, err := r.db.Exec(query, u1, u2, t)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}

// Функция удаления связей
func (r relationshipRepoImpl) Delete(u1, u2 int) (err error) {
	query := `delete from relationship where user1=$1 and user2=$2`
	res, err := r.db.Exec(query, u1, u2)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return
}
