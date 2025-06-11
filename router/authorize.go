package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *serviceImpl) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := r.authService.BearerToken(ctx.GetHeader("Authorization"))

		valid, msg, err := r.authService.Authorize(token)

		if !valid {
			ctx.String(http.StatusUnauthorized, msg)
			ctx.Abort()
			return
		}

		if msg == "valid refresh token" {
			ctx.String(http.StatusOK, "This is a refresh token")
			ctx.Abort()
			return
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		ctx.Next()
	}
}
