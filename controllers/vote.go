package controllers

import (
	"blog/logic"
	"blog/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostVoteControllers(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvaildParam)
		return
	}
	// 获取当前用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 投票业务
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	} 
	ResponseSuccess(c, nil)
}