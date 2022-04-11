package repository

import (
	"database/sql"

	"app/internal/model"
)

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
	Insert(profile *model.Profile) (int, error)
	Select(id int) (model.Profile, error)
	SelectByEmail(email string) (model.Profile, error)
	//Update(profile *model.Profile) error
	//Delete(id int) error
}

type PostRepo interface {
	Insert(post *model.Post) (int, error)
	Select(postId int) (model.Post, error)
	SelectByUserId(userId int) ([]int64, error)
	//Update(post *model.Post) error
	//Delete(userId, postId int) error
}

type CommentRepo interface {
	Insert(cmt *model.Comment) (int, error)
	Select(postId int) ([]model.Comment, error)
	//Update(cmt *model.Comment) error
	//Delete(userId, cmtId int) error
}

type ReactionRepo interface {
	InsertUpdate(userId, postId int, reaction string) error
	Select(postId int) ([]model.Reaction, error)
	SelectReactionOfUser(userId, postId int) (string, error)
	//Update(userId, postId int, reaction string) error
	//Delete(userId, postId int) error
}

type RelationshipRepo interface {
	Select(userId int) ([]model.Relationship, error)
	Friends(userId int) ([]model.Relationship, error)
	Requests(userId int) ([]model.Relationship, error)
}

type NotificationRepo interface {
	Insert(notif *model.Notification) (int, error)
	Select(userId int) ([]model.Notification, error)
}

type AlbumRepo interface {
	Insert(album *model.Album) (int, error)
	Select(id int) (model.Album, error)
	SelectByUserId(userId int) ([]model.Album, error)
}

type PhotoRepo interface {
	Insert(photo *model.Photo) (int, error)
	Select(id int) (model.Photo, error)
	SelectByUserId(userId int) ([]model.Photo, error)
}

func NewRepo(db *sql.DB) (repo *Repo) {
	repo = &Repo{
		Profile:      NewProfileRepo(db),
		Post:         NewPostRepo(db),
		Comment:      NewCommentRepo(db),
		Reaction:     NewReactionRepo(db),
		Relationship: NewRelationshipRepo(db),
		Notification: NewNotificationRepo(db),
		Album:        NewAlbumRepo(db),
		Photo:        NewPhotoRepo(db),
	}
	return
}
