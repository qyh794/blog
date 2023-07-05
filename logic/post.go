package logic

import (
	"blog/cache"
	"blog/dao/postgresql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.PostID = snowflake.GenID()
	err = postgresql.CreatePost(p)
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
	post, err := postgresql.GetPostByID(pID)
	if err != nil {
		zap.L().Error("postgresql.GetPostByID failed, err:", zap.Error(err))
		return
	}
	user, err := postgresql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("postgresql.GetUserByID failed, err:", zap.Int64("authorid:", post.AuthorID), zap.Error(err))
		return
	}

	community, err := postgresql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("postgresql.GetCommunityDetailByID failed, err:", zap.Error(err))
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
	posts, err := postgresql.GetPostList(page, size)
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := postgresql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("postgresql.GetUserByID failed, err:", zap.Error(err))
			continue
		}

		community, err := postgresql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("postgresql.GetCommunityDetailByID failed, err:", zap.Error(err))
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
	posts, err := postgresql.GetPostListByCommunityID(communityId)
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := postgresql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("postgresql.GetUserByID failed, err:", zap.Error(err))
			continue
		}
		community, err := postgresql.GetCommunityDetailByID(int64(communityId))
		if err != nil {
			zap.L().Error("postgresql.GetCommunityDetailByID failed, err:", zap.Error(err))
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

// 根据时间或者分数获取帖子
func GetPostListByTimeOrScore(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 查询缓存
	cacheData, err := cache.GetPostListFromCache()
	// 缓存命中
	if len(cacheData) > 0 {
		zap.L().Info("缓存命中")
		return cacheData, nil
	}
	zap.L().Info("缓存未命中")

	// 根据帖子的分数或者时间在redis中获取帖子的id
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	// 得到帖子的id后，从MySQL中获取帖子详细信息
	posts, err := postgresql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// 查询出帖子的投票情况
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	// 将帖子信息进行组合后返回
	for i, post := range posts {
		// 获取作者信息
		user, err := postgresql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("postgresql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		communityDetail, err := postgresql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("postgresql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("CommunityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[i],
			Post:            post,
			CommunityDetail: &communityDetail,
		}
		data = append(data, postDetail)
	}
	// 添加缓存
	err = redis.AddPostCache(data)
	if err == nil {
		zap.L().Info("添加缓存成功")
	} else {
		zap.L().Error("redis.AddPostCache failed",
			zap.Error(err))
	}
	return
}

func DeletePostByID(userID, postIdOfDelete int64) (err error) {
	// 2.根据要删除的帖子id查询其创建者的id
	authorId, err := postgresql.GetUserIDByPostID(postIdOfDelete)
	if err != nil {
		zap.L().Error("postgresql.GetUserIDByPostID failed, err:", zap.Error(err))
		return
	}
	// 3.判断帖子创建者的id与当前登录的用户id是否相同
	// 4.不相同则返回相应删除失败，权限不足

	if userID != authorId.AuthorID {
		return errors.New("权限不足")
	}
	// 5.相同则进入mysql进行根据帖子id删除帖子
	err = postgresql.DeletePostByID(postIdOfDelete)
	if err != nil {
		zap.L().Error("postgresql.DeletePostByID failed, err:", zap.Error(err))
		return
	}
	return
}
