package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mybluebell/logic"
)

func CommunityHandler(c *gin.Context) {

	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseErr(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}
