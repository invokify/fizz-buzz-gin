package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	}

	return r
}
