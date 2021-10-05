package model

import "time"

// Admin 管理者帳號
type Admin struct {
	ID        int       `json:"id" gorm:"column:id;type:int(11) comment '用戶ID';unsigned auto_increment;not null;primary_key"`
	Account   string    `json:"account" gorm:"column:account;type:varchar(30) comment '用戶帳號';not null;unique"`
	Level     int       `json:"level" gorm:"column:level;type:tinyint(4) comment '用戶權限(0:可編輯管理 1:超級使用者)';default:0"`
	EditedBy  string    `json:"edited_by" gorm:"column:edited_by;type:varchar(30) comment '最後編輯人員'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP comment '資料建立時間'; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP comment '資料最後更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName 设置 Admin 的表名为 `admin`
func (Admin) TableName() string {
	return "admin"
}
