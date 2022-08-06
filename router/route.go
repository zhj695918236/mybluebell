package router

import (
	"github.com/gin-gonic/gin"
	"mybluebell/controller"
)

func SetupRouter(mode string) *gin.Engine{
	if mode== gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)
	}
	r :=gin.New()



	v1:=r.Group("/api/v1")

	// signup
	v1.POST("/signup",controller.SignUpHandler)
	// login
	v1.POST("/login",controller.LoginHandler)


}
