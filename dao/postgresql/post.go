package postgresql

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `
insert into 
    "post" (post_id, title, content, author_id, community_id) 
values
    ($1,$2,$3,$4,$5)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		zap.L().Error("sql error")
	}
	return
}

// GetPostById 根据id查询单个帖子的数据
func GetPostByID(pID int64) (post *models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time, update_time 
from 
    "post" 
where 
    "post_id" = $1`
	post = new(models.Post)
	err = db.Get(post, sqlStr, pID)
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int) (posts []*models.Post, err error) {
	sqlStr := `select
    post_id, title, content, author_id, community_id, create_time, update_time 
from 
    "post" 
order by 
    create_time 
desc 
LIMIT 
    $1 
offset 
    $2`
	posts = make([]*models.Post, 0)
	err = db.Select(&posts, sqlStr, page, size)
	return
}

// GetPostListByCommunityID 获取社区帖子
func GetPostListByCommunityID(communityId int) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time, update_time 
from 
    "post" 
where 
    community_id = $1`
	posts = make([]*models.Post, 0)
	err = db.Select(&posts, sqlStr, communityId)
	return
}

// GetUserIDByPostID 获取帖子作者
func GetUserIDByPostID(postId int64) (post models.Post, err error) {
	sqlStr := `select 
    author_id 
from 
    "post"
where 
    post_id = $1`
	err = db.Get(&post, sqlStr, postId)
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	// 查询出来的结果可以自己手动排序，也可以使用mysql的内置排序函数进行排序
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time 
from 
    post 
where post_id 
          in ($1) 
order by 
    FIND_IN_SET(post_id, $2)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	// 需要在select函数中修改切片的值，需要取址符
	err = db.Select(&postList, query, args...)
	return
}

// DeletePostByID 删除帖子
func DeletePostByID(postId int64) (err error) {
	sqlStr := `delete from 
           post 
       where 
           post_id = $1`
	ret, err := db.Exec(sqlStr, postId)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return ErrorDelete
	}
	return
}
