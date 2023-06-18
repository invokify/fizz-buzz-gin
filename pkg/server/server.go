package server

import (
	"fizz-buzz-gin/pkg/server/docs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewServer(timeout time.Duration) *gin.Engine {
	// set the gin mode
	gin.SetMode(gin.DebugMode)

	// instantiate a new gin router
	r := gin.New()

	// limited proxies
	r.SetTrustedProxies([]string{"localhost"})

	// set the default middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Serve swagger documentation
	r.GET("/swagger/*any", setDocumentationInfo, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// set the default route for health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// set the api group routes
	api := r.Group("/api/v1")
	{
		api.GET("/fizz-buzz", fizzBuzzHandler(timeout))
		api.GET("/stats", statisticsHandler())
	}

	return r
}

func setDocumentationInfo(c *gin.Context) {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = c.Request.Host
	docs.SwaggerInfo.BasePath = "/api/v1"
}
