package middlewares

import (
	"full_login/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "login: Unauthorized")

			c.Abort()
			return
		}
		c.Next()
	}
}
