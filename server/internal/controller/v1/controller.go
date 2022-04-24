package v1

import (
	"os"

	"app/config"
	"app/internal/service"
	"app/pkg/auth"
	"app/pkg/logger"
	"app/pkg/utils"
)

var toInt = utils.ToInt

type Controller struct {
	conf     *config.Config
	services *service.Services
	auth     *auth.Manager
	logger   *logger.Logger
}

func NewController(services *service.Services, conf *config.Config) *Controller {
	f, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		f = os.Stdout
	}

	return &Controller{
		conf:     conf,
		services: services,
		auth:     auth.InitManager("id", conf.SigningKey),
		logger:   logger.LoggerWithWriter(f),
	}
}
