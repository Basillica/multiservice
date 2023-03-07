package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/gateway/types/appenv"
)

func CORSMiddleware() gin.HandlerFunc {
	alllowedHeaders := "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, x-csrftoken, Authorization, accept, origin, Cache-Control, X-Requested-With, sentry-trace, x_bearer_token, x-bearer-token, content-disposition"
	return func(c *gin.Context) {
		ae := c.MustGet("appenv").(*appenv.AppConfig)

		c.Writer.Header().Set("Access-Control-Allow-Origin", ae.AllowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", alllowedHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization, Set-Cookie")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
