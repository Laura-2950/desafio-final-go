package middleware

import (
	"net/http"
	"os"

	"github.com/Laura-2950/desafio-final-go/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.Unauthorized("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.Unauthorized("invalid token"))
			return
		}
		c.Next()
	}
}
