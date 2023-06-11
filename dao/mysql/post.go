package mysql

import "blog/models"

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据id查询单个帖子的数据
func GetPostByID(pID int64) (post *models.Post, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post where post_id = ?"
	post = new(models.Post)
	err = db.Get(post, sqlStr, pID)
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int) (posts []*models.Post, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post order by create_time desc limit ?, ?"
	posts = make([]*models.Post, 0)
	err = db.Select(&posts, sqlStr, (page - 1) * size, size)
	return
}

// GetPostListByCommunityID 获取社区帖子
func GetPostListByCommunityID(communityId int) (posts []*models.Post, err error){
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post where community_id = ?"
	posts = make([]*models.Post, 0)
	err = db.Select(&posts, sqlStr, communityId)
	return
}

// GetUserIDByPostID 获取帖子作者
func GetUserIDByPostID(postId int64) (post models.Post, err error) {
	sqlStr := "select author_id from post where post_id = ?"
	err = db.Get(&post, sqlStr, postId)
	return
}

// DeletePostByID 删除帖子
func DeletePostByID(postId int64) (err error) {
	sqlStr := "delete from post where post_id = ?"
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