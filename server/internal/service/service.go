package service

import (
	"app/internal/repository"
	"sync"
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

type ProfileService interface {
}

type PostService interface {
}

type CommentService interface {
}

type ReactionService interface {
}

type RelationshipService interface {
}

type NotificationService interface {
}

type PhotoService interface {
}

type FeedService interface {
}

func GetServices(repo *repository.Repo) *Services {
	once.Do(func() {
		services = &Services{}
	})
	return services
}
