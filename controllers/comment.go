package controllers

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreateCommentHandler 创建评论
func CreateCommentHandler(c *gin.Context) {
	// 判断用户是否登录,获取当前评论的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 定义模型
	var comment models.Comment
	// 获取前端参数并绑定到定义的模型中
	err = c.BindJSON(&comment)
	if err != nil {
		zap.L().Error("BindJSON failed", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	// 生成评论id
	comment.ParentID = userID
	if err = logic.CreateComment(&comment); err != nil {
		zap.L().Error("logic.CreatePost failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// GetCommentListHandler 根据帖子id获取评论列表
func GetCommentListHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		zap.L().Error("get postID with invalid param", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	comments, err := logic.GetCommentByPostID(postID)
	if err != nil {
		zap.L().Error("logic.GetCommentByPostID failed", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, comments)
}
