package router

import (
	"os"

	"github.com/cesarbrancalhao/WorterLab/api/middleware"
	"github.com/cesarbrancalhao/WorterLab/api/router/routes"
	"github.com/gin-gonic/gin"
)

func GetRouterConfig() (string, string) {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if port == "" || host == "" {
		port = "8080"
		host = "localhost"
	}

	return host, port
}

func Start() {
	host, port := GetRouterConfig()

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Auth())

	routes.Route(r)
	r.Run(host + ":" + port)
}
