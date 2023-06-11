package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/snowflake"
	"go.uber.org/zap"
)

// CreateComment 创建评论
func CreateComment(comment *models.Comment) (err error) {
	// 前端传入的数据, 评论内容, 帖子id, 后端需要的内容评论的id,作者id
	// 评论者的id已经有了
	// 评论id
	comment.CommentID = snowflake.GenID()
	// 根据帖子id查找作者id
	Author, err := mysql.GetUserIDByPostID(comment.PostID)
	if err != nil {
		zap.L().Error("mysql.GetUserIDByPostID failed, err:", zap.Error(err))
		return
	}
	comment.AuthorID = Author.AuthorID
	err = mysql.CreateComment(comment)
	return
}

func GetCommentByPostID(postID int) (data []*models.Comment, err error) {
	return mysql.GetCommentByPostID(postID)
}
