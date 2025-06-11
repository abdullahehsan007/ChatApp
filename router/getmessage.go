package router

// func (h *routerImpl) GetMessage() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var get model.Get
// 		if err := ctx.ShouldBindJSON(&get); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not in JSON"})
// 			return
// 		}

// 		_, err := h.service.GetMessage(ctx, get, TokenString)
// 		if err != nil {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}
// }
