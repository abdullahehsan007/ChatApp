package router

import (
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) RoutersSetup(router *gin.Engine) {
	router.POST("/signup", r.service.handleSignUp)
	router.POST("/login", r.service.Login())
	router.POST("/auth", r.service.Authorize())
	router.POST("/ref", r.service.Refresh())
	router.POST("/send", r.service.Authorize(), r.service.SendMessage())
	router.Run(":8080")
}
