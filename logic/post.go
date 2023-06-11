package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.PostID = snowflake.GenID()
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	// 创建帖子应该在redis中记录帖子的时间戳和帖子分数
	err = redis.CreatePost(p.PostID)
	return
}

func GetPostDetailByID(pID int64) (data *models.PostDetail, err error) {
	// 帖子详情应该包含什么
	// 帖子的id, title, content, authorid, community_id, create_time
	// 帖子所属社区的信息
	// 帖子的作者名字
	post, err := mysql.GetPostByID(pID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID failed, err:", zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed, err:", zap.Int64("authorid:", post.AuthorID), zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed, err:", zap.Error(err))
		return
	}

	data = &models.PostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: &community,
	}
	return
}

func GetPostList(page, size int) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed, err:", zap.Error(err))
			continue
		}

		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed, err:", zap.Error(err))
			continue
		}

		postDetailOne := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: &community,
		}
		data = append(data, postDetailOne)
	}
	return
}

func GetPostListByCommunity(communityId int) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostListByCommunityID(communityId)
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed, err:", zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailByID(int64(communityId))
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed, err:", zap.Error(err))
			continue
		}
		postDetailOne := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: &community,
		}
		data = append(data, postDetailOne)
	}
	return
}

func DeletePostByID(userID, postIdOfDelete int64) (err error) {
	// 2.根据要删除的帖子id查询其创建者的id
	authorId, err := mysql.GetUserIDByPostID(postIdOfDelete)
	if err != nil {
		zap.L().Error("mysql.GetUserIDByPostID failed, err:", zap.Error(err))
		return
	}
	// 3.判断帖子创建者的id与当前登录的用户id是否相同
	// 4.不相同则返回相应删除失败，权限不足

	if userID != authorId.AuthorID {
		return errors.New("权限不足")
	}
	// 5.相同则进入mysql进行根据帖子id删除帖子
	err = mysql.DeletePostByID(postIdOfDelete)
	if err != nil {
		zap.L().Error("mysql.DeletePostByID failed, err:", zap.Error(err))
		return
	}
	return
}
