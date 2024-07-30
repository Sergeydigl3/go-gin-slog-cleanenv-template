package handlers

import (
	"gin-template/internal/server/services/auth"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	log         *slog.Logger
	authService *auth.Service
}

func NewAuthHandler(log *slog.Logger, authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		log:         log,
		authService: authService,
	}
}

func (s *AuthHandler) Login(c *gin.Context) {
	s.authService.Login()
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}

func (s *AuthHandler) Logout(c *gin.Context) {
	s.authService.Logout()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
