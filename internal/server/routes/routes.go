package routes

import (
	"gin-template/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AuthHandler *handlers.AuthHandler
}

func SetupRoutes(router *gin.Engine, h *Handlers) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	v1.GET("/auth/login", h.AuthHandler.Login)
	v1.GET("/auth/logout", h.AuthHandler.Logout)

}
