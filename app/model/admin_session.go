package model

import "time"

// AdminSession 管理者session紀錄
type AdminSession struct {
	Account    string    `json:"account" gorm:"column:account"`
	Session    string    `json:"session" gorm:"column:session"`
	ExpireTime time.Time `json:"expire_time" gorm:"column:expire_time"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName 设置 AdminSession 的表名为 `admin_session`
func (AdminSession) TableName() string {
	return "admin_session"
}
