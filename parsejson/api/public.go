package api

import (
	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/parsejson/handlers"
)

func AddAuthRoutes(r *gin.Engine) {
	// public Endpoints
	r.GET("/parsejson", handlers.ParseJSON)
}
