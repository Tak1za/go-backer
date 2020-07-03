package middlewares

import (
	"github.com/Tak1za/ivar/config"
	"github.com/gin-gonic/gin"
)

func DriverMiddleware(driver *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("driver", driver)
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Next()
	}
}
