package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"用户Id"`
	UserName  string    `gorm:"type:text;not null;size:20" json:"用户名"`
	Email     string    `gorm:"type:text" json:"邮箱"`
	CreatedAt time.Time `json:"创建时间"`
	UpdatedAt time.Time `json:"更新时间"`
	Ignored   string    `json:"数据状态"`
}
