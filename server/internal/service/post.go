package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type PostServiceImpl struct {
	repo repository.PostRepo
}

// Функция создания сервиса публикации
func NewPostService(repo repository.PostRepo) PostService {
	return &PostServiceImpl{repo}
}

// Функция  получения публикации
func (r *PostServiceImpl) Get(postId int) (post model.Post, err error) {
	return r.repo.Select(postId)
}

// Функция получения публикации пользователя
func (r *PostServiceImpl) GetByUserId(userId int) ([]int64, error) {
	return r.repo.SelectByUserId(userId)
}

// Функция получения реакции на публикацию
func (r *PostServiceImpl) GetReaction(postId int) ([]int64, error) {
	return r.repo.SelectReaction(postId)
}

// Функция добавления публикации
func (r *PostServiceImpl) Post(userId int, body PostBody) (err error) {
	post := model.Post{
		UserId:   userId,
		Tags:     body.Tags,
		Content:  body.Content,
		AtchType: body.AtchType,
		AtchId:   body.AtchId,
		AtchUrl:  body.AtchUrl,
	}
	if post.AtchType == "photo" {
		photoId, err := services.Photo.UploadPhoto(model.Photo{
			UserId: post.UserId,
			Url:    post.AtchUrl,
		})

		if err != nil {
			return err
		}
		post.AtchId = int(photoId)
	}
	_, err = r.repo.Insert(&post)
	return
}

// Функция удаления публикации
func (r *PostServiceImpl) Delete(userId int, postId int) error {
	return r.repo.Delete(userId, postId)
}
