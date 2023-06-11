package models

import "time"

type Post struct {
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	PostID      int64     `json:"post_id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type PostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*Post
	*CommunityDetail `json:"community"`
}
