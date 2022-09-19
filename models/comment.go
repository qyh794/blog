package models

import "time"

type Comment struct {
	PostID    int64     `db:"post_id" json:"post_id"`
	ParentID  int64     `db:"parent_id" json:"parent_id"`
	AuthorID  int64     `db:"author_id" json:"author_id"`
	CommentID int64     `db:"comment_id" json:"comment_id"`
	Content   string    `db:"content" json:"content"`
	CreatTime time.Time `db:"create_time" json:"creat_time"`
}
