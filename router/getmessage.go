package router

import (
	"chatapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *serviceImpl) GetMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var get model.Get
		if err := ctx.ShouldBindJSON(&get); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
			return
		}

		message, err := h.messageService.GetMessage(get, TokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		} else if len(message) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "No messages found",
				"message":  []model.Message{},
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Message Successfully Received",
				"message":  message,
			})
		}
	}
}
