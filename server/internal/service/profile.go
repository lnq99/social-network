package service

import (
	"app/internal/model"
	"app/internal/repository"
	"app/pkg/auth"
)

type ProfileServiceImpl struct {
	repo repository.ProfileRepo
}

// Функция созадния сервиса профиля
func NewProfileService(repo repository.ProfileRepo) ProfileService {
	return &ProfileServiceImpl{repo}
}

// Функция получения профиля
func (r *ProfileServiceImpl) Get(id int) (model.Profile, error) {
	return r.repo.Select(id)
}

// Функция получения профиля по почте
func (r *ProfileServiceImpl) GetByEmail(e string) (model.Profile, error) {
	return r.repo.SelectByEmail(e)
}

// Функция поиска профиля по имени
func (r *ProfileServiceImpl) SearchName(id int, s string) (string, error) {
	return r.repo.SearchName(id, s)
}

// Функция  регистрации нового пользователя
func (r *ProfileServiceImpl) Register(body ProfileBody) (err error) {
	manager := auth.GetManager()
	salt, hashed := manager.GetHashSalt(body.Password)
	p := model.Profile{
		Email:     body.Email,
		Name:      body.Username,
		Salt:      salt,
		Hash:      hashed,
		Gender:    body.Gender,
		Birthdate: body.Birthdate,
	}
	_, err = r.repo.Insert(&p)
	return
}

// Функция установки аватарки
func (r *ProfileServiceImpl) SetAvatar(p model.Photo) error {
	return r.repo.UpdateAvatar(p)
}

// Функция изменения информации о себе
func (r *ProfileServiceImpl) ChangeIntro(id int, intro IntroBody) error {
	return r.repo.UpdateIntro(id, intro.Intro)
}
