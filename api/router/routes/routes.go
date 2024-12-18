package routes

import "github.com/gin-gonic/gin"

var RoutesList = []func(r *gin.Engine){
	DefaultRoutes,
	// WordRoutes, todo
}

func DefaultRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func Route(r *gin.Engine) {
	for _, route := range RoutesList {
		route(r)
	}
}
