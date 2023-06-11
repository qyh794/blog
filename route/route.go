package route

import (
	"blog/controllers"
	"blog/logger"
	"blog/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("api/blog")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// 获取社区列表
		v1.GET("community", controllers.CommunityHandler)
		// 根据id获取社区信息
		v1.GET("community/:id", controllers.CommunityDetailHandler)
		// 创建帖子
		v1.POST("createpost", controllers.CreatePostHandler)
		// 删除帖子
		v1.DELETE("deletepost/:postid", controllers.DeletePostHandler)
		// 根据id获取帖子详情
		v1.GET("postdetail/:id", controllers.GetPostDetailHandler)
		// 根据社区获取帖子
		v1.GET("postlistbycommunity/:communityid", controllers.PostListByCommunity)
		// 获取所有帖子信息
		v1.GET("postlist", controllers.GetPostListHandler)
		// 根据时间或者分数获取帖子列表
		v1.GET("postlistbytimeorscore/:order", controllers.GetPostListByTimeOrScore)
		// 给帖子投票
		v1.POST("/vote", controllers.PostVoteControllers)
		// 创建评论
		v1.POST("/comment", controllers.CreateCommentHandler)
		// 获取帖子评论
		v1.GET("/comment/:id", controllers.GetCommentListHandler)
	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
