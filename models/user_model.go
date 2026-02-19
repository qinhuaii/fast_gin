package models

type UserModel struct {
	Model
	Username string `gorm:"size:16"json:"username"`
	Nickname string `gorm:"size:32"json:"nickname"`
	Password string `gorm:"size:64"json:"password"`
	RoleID   uint8  `json:"role_id"` // 1：管理员；2：普通用户
}
