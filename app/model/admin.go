package model

import "time"

// Admin 管理者帳號
type Admin struct {
	ID        int       `json:"id" gorm:"column:id;not null;unique;primary_key"`
	Account   string    `json:"account" gorm:"column:account;not null"`
	Password  string    `json:"password" gorm:"column:pwd;not null"`
	LoginIP   string    `json:"login_ip" gorm:"column:login_ip"`
	LoginAt   time.Time `json:"login_at" gorm:"column:login_at"`
	Status    bool      `json:"status" gorm:"column:status;not null"`
	GroupID   int       `json:"group_id" gorm:"column:group_id"`
	EditedBy  string    `json:"edited_by" gorm:"column:edited_by"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName 设置 Admin 的表名为 `admin`
func (Admin) TableName() string {
	return "admin"
}
