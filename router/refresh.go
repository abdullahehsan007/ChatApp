package router

import (
	authservice "chatapp/api/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *serviceImpl) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.PostForm("r_token")
		if refreshToken == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.String(http.StatusUnauthorized, "Refresh Token Required")
			return
		}

		email, err := authservice.VerifyRefreshToken(refreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
			return
		}
		newAccessToken, _, err := authservice.CreateToken(email)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not generate new token"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"Message":      "Access Token Generated",
			"Access Token": newAccessToken,
		})

	}
}
