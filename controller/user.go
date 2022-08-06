package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mybluebell/logic"
	"mybluebell/models"
)

func SignUpHandler(c *gin.Context){
	// 1.
	p := new(models.ParamSignup)
	if err:=c.ShouldBindHeader(c);err!=nil{
		//
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		return
	}
	// 2.
	if err:=logic.SignUp(p)
}
func LoginHandler(c *gin.Context){

}