package service

import (
	"sync"

	"app/internal/model"
	"app/internal/repository"
)

var (
	once     sync.Once
	services *Services
)

type Services struct {
	Profile      ProfileService
	Post         PostService
	Comment      CommentService
	Reaction     ReactionService
	Relationship RelationshipService
	Notification NotificationService
	Photo        PhotoService
	Feed         FeedService
}

type IntroBody struct {
	Intro string `json:"intro"`
}

type CommentBody struct {
	PostId   int    `json:"postId"`
	ParentId int    `json:"parentId"`
	Content  string `json:"content"`
}

type PostBody struct {
	Tags     string `json:"tags"`
	Content  string `json:"content"`
	AtchType string `json:"atchType"`
	AtchId   int    `json:"atchId,omitempty"`
	AtchUrl  string `json:"atchUrl"`
}

type ProfileBody struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileService interface {
	Get(id int) (model.Profile, error)
	GetByEmail(e string) (model.Profile, error)
	SearchName(id int, s string) (string, error)

	Register(ProfileBody) error
	SetAvatar(model.Photo) error
	ChangeIntro(id int, intro IntroBody) error
}

type PostService interface {
	Get(postId int) (model.Post, error)
	GetReaction(postId int) ([]int64, error)
	GetByUserId(userId int) ([]int64, error)
	Post(userId int, body PostBody) error
	Delete(userId int, postId int) error
}

type CommentService interface {
	GetTree(postId int) (string, error)
	Add(userId int, body CommentBody) error
}

type ReactionService interface {
	Get(postId int) ([]model.Reaction, error)
	GetByUserPost(userId, postId int) (string, error)
	UpdateReaction(userId, postId int, t string) error
}

type RelationshipService interface {
	Get(id int) ([]model.Relationship, error)
	Friends(id int) ([]model.Relationship, error)
	Requests(id int) ([]model.Relationship, error)
	FriendsDetail(id int) (string, error)
	MutualFriends(u1, u2 int) ([]int64, error)
	GetRelationshipWith(u1, u2 int) string
	ChangeType(u1, u2 int, t string) error
}

type NotificationService interface {
	Get(userId int) ([]model.Notification, error)
	Add(notif model.Notification) error
}

type PhotoService interface {
	GetAlbumByUserId(userId int) ([]model.Album, error)
	GetAlbumId(userId int, album string) (int, error)

	GetPhoto(id int) (model.Photo, error)
	GetPhotoByUserId(userId int) ([]model.Photo, error)

	UploadPhotoToAlbum(p model.Photo, album string) (int, error)
	UploadPhoto(model.Photo) (int, error)
	SetAvatar(model.Photo) error
}

type FeedService interface {
	GetFeed(id, limit, offset int) (feed []int64, err error)
	// Get(id int, tBegin, tEnd string) (newBegin, newEnd string, posts []model.Post)
}

func GetServices(repo *repository.Repo) *Services {
	once.Do(func() {
		services = &Services{
			Profile:      NewProfileService(repo.Profile),
			Post:         NewPostService(repo.Post),
			Comment:      NewCommentService(repo.Comment),
			Reaction:     NewReactionService(repo.Reaction),
			Relationship: NewRelationshipService(repo.Relationship),
			Notification: NewNotificationService(repo.Notification),
			Photo:        NewPhotoService(repo.Photo, repo.Album),
			Feed:         NewFeedService(repo.Profile, repo.Relationship),
		}
	})
	return services
}
