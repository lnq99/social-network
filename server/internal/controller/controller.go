package controller

import (
	"app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.CorsMiddleware)

	// r := gin.New()
	// r.Use(gin.LoggerWithWriter(io.MultiWriter(os.Stdout, f)))
	// r.Use(gin.Recovery())

	// r.Use(middleware.LoggerMiddleware(ctrl.logger))

	return r
}
