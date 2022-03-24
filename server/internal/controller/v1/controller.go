package v1

import (
	"app/config"
	"app/internal/service"
)

type Controller struct {
	services *service.Services
}

func NewController(services *service.Services, conf *config.Config) *Controller {
	return &Controller{}
}
