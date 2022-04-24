package middleware

import (
	"app/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type logResponseWriter struct {
	gin.ResponseWriter
	logWriter http.ResponseWriter
}

func (r logResponseWriter) Write(b []byte) (int, error) {
	r.logWriter.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r logResponseWriter) WriteHeader(statusCode int) {
	r.logWriter.WriteHeader(statusCode)
	r.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		w := logger.HandleRequest(c.Request)

		c.Writer = logResponseWriter{
			ResponseWriter: c.Writer,
			logWriter:      w,
		}

		c.Next()

		logger.HandleResponse(c.Request, c.Writer)
		logger.Params.ClientIP = c.ClientIP()
		logger.Params.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		logger.LogRequestResponse()
	}
}
