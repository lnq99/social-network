package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type PhotoServiceImpl struct {
	photo repository.PhotoRepo
	album repository.AlbumRepo
}

// Функция создания сервиса фото
func NewPhotoService(photo repository.PhotoRepo, album repository.AlbumRepo) PhotoService {
	return &PhotoServiceImpl{photo, album}
}

// Функция получения альбома фото пользователя
func (r *PhotoServiceImpl) GetAlbumByUserId(userId int) (res []model.Album, err error) {
	return r.album.SelectByUserId(userId)
}

// Функция получения альбома
func (r *PhotoServiceImpl) GetAlbumId(userId int, album string) (albumId int, err error) {
	albums, err := r.album.SelectByUserId(userId)
	if err != nil {
		for _, a := range albums {
			if a.Descr == album {
				albumId = a.Id
				return
			}
		}
	}
	return
}

// Функция получения фото
func (r *PhotoServiceImpl) GetPhoto(id int) (model.Photo, error) {
	return r.photo.Select(id)
}

// Функция получения фото пользователя
func (r *PhotoServiceImpl) GetPhotoByUserId(userId int) (res []model.Photo, err error) {
	return r.photo.SelectByUserId(userId)
}

// Функция загрузки фото в альбом
func (r *PhotoServiceImpl) UploadPhotoToAlbum(p model.Photo, album string) (photoId int, err error) {
	p.AlbumId, err = r.GetAlbumId(p.UserId, album)
	if err != nil {
		return -1, err
	}
	photoId, err = r.photo.Insert(&p)
	return
}

// Функция-оберка загрузки фото
func (r *PhotoServiceImpl) UploadPhoto(p model.Photo) (photoId int, err error) {
	return r.UploadPhotoToAlbum(p, "Upload")
}

// Функция-обертка загрузки новой аватарки
func (r *PhotoServiceImpl) SetAvatar(p model.Photo) (err error) {
	_, err = r.UploadPhotoToAlbum(p, "Avatar")
	return
}
