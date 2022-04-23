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
	Select(id int) (model.Profile, error)
	SelectByEmail(e string) (model.Profile, error)
	SelectFeed(id, limit, offset int) ([]int64, error)
	SearchName(id int, s string) (string, error)
	Insert(cmt model.Profile) error
	SetAvatar(cmt model.Photo) error
	ChangeIntro(id int, intro string) error
}

type PostRepo interface {
	Select(postId int) (model.Post, error)
	SelectReaction(userId int) ([]int64, error)
	SelectByUserId(userId int) ([]int64, error)
	Insert(post model.Post) error
	Delete(userId, postId int) error
}

type CommentRepo interface {
	Select(postId int) ([]model.Comment, error)
	Insert(cmt model.Comment) error
}

type ReactionRepo interface {
	Select(postId int) ([]model.Reaction, error)
	SelectByUserPost(userId, postId int) (string, error)
	UpdateReaction(userId, postId int, t string) error
}

type RelationshipRepo interface {
	Select(id int) ([]model.Relationship, error)
	Friends(id int) ([]model.Relationship, error)
	Requests(id int) ([]model.Relationship, error)
	FriendsDetail(id int) (string, error)
	MutualFriends(u1, u2 int) ([]int64, error)
	SelectRelationshipWith(u1, u2 int) string
	ChangeType(u1, u2 int, t string) error
	Delete(u1, u2 int) error
}

type NotificationRepo interface {
	Select(userId int) ([]model.Notification, error)
	Insert(notif model.Notification) error
}

type AlbumRepo interface {
	Select(userId int) ([]model.Album, error)
	SelectByUserIdAndAlbumName(userId int, album string) (int, error)
}

type PhotoRepo interface {
	Select(id int) (model.Photo, error)
	SelectByUserId(userId int) ([]model.Photo, error)
	Insert(photo model.Photo) (int64, error)
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
