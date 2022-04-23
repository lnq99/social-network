package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type ReactionServiceImpl struct {
	repo repository.ReactionRepo
}

func NewReactionService(repo repository.ReactionRepo) ReactionService {
	return &ReactionServiceImpl{repo}
}

func (r *ReactionServiceImpl) Get(postId int) (res []model.Reaction, err error) {
	return r.repo.Select(postId)
}

func (r *ReactionServiceImpl) GetByUserPost(userId, postId int) (string, error) {
	return r.repo.SelectReactionOfUser(userId, postId)
}

func (r *ReactionServiceImpl) UpdateReaction(userId, postId int, t string) error {
	return r.repo.InsertUpdate(userId, postId, t)
}
