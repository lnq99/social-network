package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type ReactionServiceImpl struct {
	repo repository.ReactionRepo
}

// Функция создания сервиса реакции
func NewReactionService(repo repository.ReactionRepo) ReactionService {
	return &ReactionServiceImpl{repo}
}

// Функция  получения реакции
func (r *ReactionServiceImpl) Get(postId int) (res []model.Reaction, err error) {
	return r.repo.Select(postId)
}

// Функция получения реакции на публикацию пользователя
func (r *ReactionServiceImpl) GetByUserPost(userId, postId int) (string, error) {
	return r.repo.SelectReactionOfUser(userId, postId)
}

// Функция обновления реакции на публикацию
func (r *ReactionServiceImpl) UpdateReaction(userId, postId int, t string) error {
	return r.repo.InsertUpdate(userId, postId, t)
}
