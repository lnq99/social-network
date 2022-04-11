package repository

import (
	"database/sql"

	"app/internal/model"
)

type relationshipRepoImpl struct {
	db *sql.DB
}

func NewRelationshipRepo(db *sql.DB) RelationshipRepo {
	return &relationshipRepoImpl{db}
}

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

func (r relationshipRepoImpl) Select(userId int) ([]model.Relationship, error) {
	return r.selectById(`select * from relationship where user1=$1`, userId)
}

func (r relationshipRepoImpl) Friends(userId int) ([]model.Relationship, error) {
	return r.selectById(`select * from relationship where user1=$1 and type='friend'`, userId)
}

func (r relationshipRepoImpl) Requests(userId int) ([]model.Relationship, error) {
	return r.selectById(`"select * from relationship where user1=$1 and type='request'`, userId)
}
