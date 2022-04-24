package middleware

import (
	"app/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(auth *auth.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := 0

		token, err := c.Cookie("token")
		if err == nil {
			id, err = auth.ParseTokenId(token)
			if err == nil {
				c.Set("ID", id)
				// log.Println(id, token, err)
				return
			}
		}

		err = auth.TokenValid(c.Request)
		// log.Println(err)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
