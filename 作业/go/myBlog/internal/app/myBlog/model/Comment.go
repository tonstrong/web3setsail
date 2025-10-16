package model

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"评论ID"`
	PostID    uint      `gorm:"size:200;not null" json:"帖子ID"`
	Content   string    `gorm:"type:text" json:"评论内容"`
	ParentID  uint      `gorm:"not null" json:"父评论ID"`
	CreatedAt time.Time `json:"创建时间"`
	UpdatedAt time.Time `json:"更新时间"`
	Ignored   string    `json:"数据状态"`
}
