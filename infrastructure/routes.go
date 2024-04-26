package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

// NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	httpRouter.RedirectTrailingSlash = true

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}

}
