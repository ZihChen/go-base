package model

import "time"

// Admin 管理者帳號
type Admin struct {
	ID        int       `json:"id" gorm:"column:id;unsigned auto_increment comment '用戶ID';not null;primary_key"`
	Account   string    `json:"account" gorm:"column:account; comment '用戶帳號';not null;unique"`
	EditedBy  string    `json:"edited_by" gorm:"column:edited_by; comment '最後編輯人員'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; comment '資料建立時間'; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at; comment '資料最後更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName 设置 Admin 的表名为 `admin`
func (Admin) TableName() string {
	return "admin"
}
