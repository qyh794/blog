package controllers

import (
	"blog/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityList failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 获取url参数
	idstr := c.Param("id")
	id1, err := strconv.Atoi(idstr)
	if err != nil {
		ResponseError(c, CodeInvaildParam)
		return
	}
	id := int64(id1)
	data, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailByID() failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}