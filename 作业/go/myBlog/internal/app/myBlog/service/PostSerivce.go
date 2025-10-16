package service

import (
	myBlogModel "myBlog/internal/app/myBlog/model"
	pkgModel "myBlog/internal/pkg/model"
)

func GetPostWithComments(userId uint) ([]pkgModel.PostWithComments, error) {
	var posts []myBlogModel.Post

}
