package model

import (
	"myBlog/internal/app/myBlog/model"
)

type PostWithComments struct {
	Post     model.Post      `json:"post"`
	Comments []model.Comment `json:"comments"`
}
