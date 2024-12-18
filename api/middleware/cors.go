package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/cesarbrancalhao/WorterLab/pkg/util"
	"github.com/gin-gonic/gin"
)

const maxAge = "3600" // 1 hour

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := getOrigin()
		allowed := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

		if origin != "" && util.SliceCheck(allowed, origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Max-Age", maxAge)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func getOrigin() string {
	origin := os.Getenv("ORIGIN")
	allowed := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ";")

	if util.SliceCheck(allowed, origin) {
		return origin
	}

	return ""
}
