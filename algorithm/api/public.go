package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/x/multiservice/algorithm/handlers"
)

func AddAuthRoutes(r *gin.Engine) {
	r.POST("/algorithm", handler.Algorithm)
}
