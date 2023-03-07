package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	public "github.com/x/multiservice/gateway/handlers"
	"github.com/x/multiservice/gateway/middleware"
)

func AddPublicRoutes(r *gin.Engine) {
	// public Endpoints
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", public.Login)
	r.GET("/authenticate", middleware.JwtAuthMiddleware(), public.Authenticate)
}
