package service

import (
	"app/internal/repository"
)

type FeedServiceImpl struct {
	profile      repository.ProfileRepo
	relationship repository.RelationshipRepo
}

func NewFeedService(
	profile repository.ProfileRepo,
	relationship repository.RelationshipRepo) FeedService {
	return &FeedServiceImpl{profile, relationship}
}

func (r *FeedServiceImpl) GetFeed(id, limit, offset int) (feed []int64, err error) {
	return r.profile.SelectFeed(id, limit, offset)
}
