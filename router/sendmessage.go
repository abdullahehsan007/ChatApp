package router

import (
	"chatapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *routerImpl) SendMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var message model.Message
		if err := ctx.ShouldBindJSON(&message); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}
        
		_,err:= h.service.SendMessage(ctx, message, TokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}
}
