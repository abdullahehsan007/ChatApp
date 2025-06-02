package router

import (
	"chatapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routerImpl) handleSignUp(ctx *gin.Context) {
	var user model.Info
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := r.service.SignUp(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "user_name": user.Username})
}
