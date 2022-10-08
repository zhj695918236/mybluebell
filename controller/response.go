package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  interface{} `json:"msg"`
	Code ResCode     `json:"code"`
	Data interface{} `json:"data"`
}

func ResponseErr(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Msg:  code.Msg(),
		Code: code,
		Data: nil,
	})
}

func ResponseErrWithMsg(c *gin.Context, code ResCode, Msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Msg:  Msg,
		Code: code,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Msg:  CodeSuccess.Msg(),
		Code: CodeSuccess,
		Data: data,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Msg:  msg,
		Code: code,
		Data: nil,
	})

}
