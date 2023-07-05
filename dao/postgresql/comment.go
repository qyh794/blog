package postgresql

import "blog/models"

func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into 
    comment(comment_id ,content, post_id, author_id, parent_id) 
values
    ($1,$2,$3,$4,$5)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID, comment.AuthorID, comment.ParentID)
	return err
}

func GetCommentByPostID(postID int) (data []*models.Comment, err error) {
	sqlStr := `select 
    comment_id, content, post_id, author_id, parent_id, create_time 
from 
    comment 
where 
    post_id = $1`
	data = make([]*models.Comment, 0)
	err = db.Select(&data, sqlStr, postID)
	return
}
