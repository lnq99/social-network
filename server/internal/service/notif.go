package service

import (
	"app/internal/model"
	"app/internal/repository"
)

type NotificationServiceImpl struct {
	repo repository.NotificationRepo
}

// Функция создания сервиса уведомления
func NewNotificationService(repo repository.NotificationRepo) NotificationService {
	return &NotificationServiceImpl{repo}
}

// Функция получения уведомления
func (r *NotificationServiceImpl) Get(userId int) (res []model.Notification, err error) {
	return r.repo.Select(userId)
}

// Функция создания нового уведомления
func (r *NotificationServiceImpl) Add(notif model.Notification) (err error) {
	_, err = r.repo.Insert(&notif)
	return
}
