package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/x/multiservice/gateway/docs"
	"github.com/x/multiservice/gateway/middleware"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.ConfigMiddleware(), middleware.CORSMiddleware())
	r.GET("/gateway/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	AddPublicRoutes(r)
	return r
}
