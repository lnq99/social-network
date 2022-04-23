package repository

import (
	"database/sql"

	"app/internal/model"
)

type notificationRepoImpl struct {
	db *sql.DB
}

func NewNotificationRepo(db *sql.DB) NotificationRepo {
	return &notificationRepoImpl{db}
}

func scanNotification(row MultiScanner, c *model.Notification) error {
	err := row.Scan(
		&c.Id,
		&c.UserId,
		&c.T,
		&c.Created,
		&c.FromUserId,
		&c.PostId,
		&c.CmtId,
	)
	return err
}

func (r notificationRepoImpl) Insert(notif *model.Notification) (id int, err error) {
	query := `insert into Notification(userId, type, fromUserId, postId, cmtId)
		values ($1, $2, $3, $4, $5) returning id`
	row := r.db.QueryRow(query, notif.UserId, notif.T, notif.FromUserId, notif.PostId, notif.CmtId)
	err = row.Scan(&id)
	return
}

func (r notificationRepoImpl) Select(userId int) (res []model.Notification, err error) {
	rows, err := r.db.Query(`select * from Notification where userId=$1 order by id desc limit 20`, userId)
	if err != nil {
		return
	}

	var notif model.Notification

	defer rows.Close()
	for rows.Next() {
		err = scanNotification(rows, &notif)
		if err != nil {
			return
		}
		res = append(res, notif)
	}

	return
}
