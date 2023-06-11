package controllers

import (
	"blog/logic"
	"blog/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler controllers 负责向前端返回相应
// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数
	// 参数校验
	p := new(models.Post)
	// 获取前端参数并绑定到定义的模型中
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invaild param", zap.Error(err))
		ResponseError(c, CodeInvaildParam)
		return
	}
	// 获取当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 创建帖子
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据帖子获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	postID, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invaild param", zap.Error(err))
		ResponseError(c, CodeInvaildParam)
		return
	}

	data, err := logic.GetPostDetailByID(postID)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取所有帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}

// PostListByCommunity 根据社区分类获取帖子列表
func PostListByCommunity(c *gin.Context) {
	communityIDStr := c.Param("communityid")
	communityId, err := strconv.Atoi(communityIDStr)
	if err != nil {
		zap.L().Error("get communityid with invalid param", zap.Error(err))
		ResponseError(c, CodeInvaildParam)
		return
	}
	data, err := logic.GetPostListByCommunity(communityId)
	if err != nil {
		zap.L().Error("logic.GetPostListByCommunity failed, err:", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListByTimeOrScore 根据时间或者分数获取帖子列表
func GetPostListByTimeOrScore(c *gin.Context) {
	// 获取参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListByTimeOrScore with invalid param", zap.Error(err))
		ResponseError(c, CodeInvaildParam)
		return
	}
	data, err := logic.GetPostListByTimeOrScore(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeSeverBusy)
		return
	}
	ResponseSuccess(c, data)
}

// DeletePostHandler 根据帖子id删除帖子
func DeletePostHandler(c *gin.Context) {
	// 获取帖子参数
	postIdStr := c.Param("postid")
	postIdOfDelete, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("get post_id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvaildParam)
		return
	}
	// 判断用户是否登录
	// 判断登录的用户是否是创建帖子的用户
	// 1.获取当前用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	err = logic.DeletePostByID(userID, postIdOfDelete)
	if err != nil {
		ResponseError(c, CodeDeletePost)
		return
	}
	ResponseSuccess(c, nil)
}
