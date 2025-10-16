package model

import (
	"time"
)

type Post struct {
	ID         uint      `gorm:"primary_key;AUTO_INCREMENT" json:"帖子ID"`
	UserId     uint      `gorm:"not null" json:"帖子所属人"`
	Title      string    `grom:"type:text" json:"帖子所属人"`
	Content    string    `gorm:"type:text" json:"文章内容"` // 文章内容
	CommentNum uint      `gorm:"-" json:"评论数量"`         //非数据库字段
	Status     uint      `gorm:"type:int;default:0" json:"帖子状态"`
	CreatedAt  time.Time `json:"创建时间"`
	UpdatedAt  time.Time `json:"更新时间"`
	Ignored    string    `json:"数据状态"`
}
