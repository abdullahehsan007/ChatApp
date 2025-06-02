package router

import (
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) RoutersSetup(router *gin.Engine) {
	router.POST("/signup", r.handleSignUp)
	router.POST("/login", r.Login())
	router.POST("/auth", r.Authorize())
	router.POST("/ref", r.Refresh())
	router.POST("/send", r.Authorize(),r.SendMessage())
	router.Run(":8080")
}
