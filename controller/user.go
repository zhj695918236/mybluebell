package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mybluebell/dao/mysql"
	"mybluebell/logic"
	"mybluebell/models"
)

func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignup)
	if err := c.ShouldBindJSON(p); err != nil {
		//
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseErr(c, CodeUserExist)
			return
		}
		ResponseErr(c, CodeServerBusy)
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("logic Login failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErr(c, CodeInvalidParam)
			return
		}
		ResponseErrWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Translate(trans)))
		return
	}
	// 2.登录
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic Login failed", zap.String("username", p.Username), zap.Error(err))
		// 这里错误分两种，password error  or  usernotexist
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseErr(c, CodeUserNotExist)
			return
		}
		ResponseErr(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
