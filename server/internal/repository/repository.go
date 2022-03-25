package repository

import "database/sql"

type Repo struct {
	Profile      ProfileRepo
	Post         PostRepo
	Comment      CommentRepo
	Reaction     ReactionRepo
	Relationship RelationshipRepo
	Notification NotificationRepo
	Album        AlbumRepo
	Photo        PhotoRepo
}

type ProfileRepo interface {
}

type PostRepo interface {
}

type CommentRepo interface {
}

type ReactionRepo interface {
}

type RelationshipRepo interface {
}

type NotificationRepo interface {
}

type AlbumRepo interface {
}

type PhotoRepo interface {
}

func NewRepo(db *sql.DB) *Repo {
	repo := &Repo{}
	return repo
}
