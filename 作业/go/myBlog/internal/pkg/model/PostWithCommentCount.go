package model

import (
	"myBlog/internal/app/myBlog/model"
)

type PostWithCommentCount struct {
	Post         model.Post `json:"post"`
	CommentCount int        `json:"comment_count"`
}
