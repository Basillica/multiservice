package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/gateway/types/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, cErr := c.Cookie("access_token")
		if cErr != nil {
			c.String(http.StatusUnauthorized, "you do not have the access rights to access this endpoint")
			c.Abort()
			return
		}
		t := auth.Token{
			Token: accessToken,
		}
		tErr := t.ValidateToken()
		if tErr != nil {
			c.String(http.StatusForbidden, "you not have valid access rights to this endpoint")
			c.Abort()
			return
		}

		c.Next()
	}
}
