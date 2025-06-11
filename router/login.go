package router

import (
	authservice "chatapp/api/auth_service"
	"chatapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TokenString string
var Refresh string

func (h *serviceImpl) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var entry model.User
		if err := ctx.ShouldBindJSON(&entry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}
		id, err := h.authService.Login(entry.Email, entry.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		TokenString, Refresh, err = authservice.CreateToken(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":       "Login Successful",
			"Token":         TokenString,
			"Refresh Token": Refresh,
		})
	}
}
