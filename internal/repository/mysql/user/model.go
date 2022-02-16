package user

import (
	"time"
)

type User struct {
	Id        int32     // 主键
	Username  string    // 用户名
	Password  string    // 密码
	Nickname  string    // 昵称
	IsDeleted int32     // 是否删除 1:是  -1:否
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
