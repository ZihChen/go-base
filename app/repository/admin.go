package repository

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"GoFormat/app/global/structs"
	"GoFormat/app/model"
	"fmt"
	"sync"
	"time"
)

// AdminRepo 管理者Reposotory
type AdminRepo struct {
}

var adminSingleton *AdminRepo
var adminOnce sync.Once

// AdminIns 獲得單例對象
func AdminIns() *AdminRepo {
	adminOnce.Do(func() {
		adminSingleton = &AdminRepo{}
	})
	return adminSingleton
}

// AdminList 取管理者清單
func (a *AdminRepo) AdminList() (admin []model.Admin, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	defer db.Close()

	if err := db.Find(&admin).Error; err != nil {
		go helper.WarnLog(fmt.Sprintf("GET_ADMIN_LIST_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("GET_ADMIN_LIST_ERROR")
		return
	}
	return
}

// GetAdmin 檢查使用者是否存在
func (a *AdminRepo) GetAdmin(account string) (admin model.Admin, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	defer db.Close()

	// 檢查帳號是否存在
	if err := db.Where("account=?", account).First(&admin).Error; err != nil {
		if fmt.Sprintf("%v", err) != "record not found" {
			apiErr = errorcode.GetAPIError("GET_USER_ACCOUNT_ERROR")
			return
		}
	}

	return
}

// UpdateLoginRecord 更新用戶登入時間 + 新增一筆Session資訊
func (a *AdminRepo) UpdateLoginRecord(user *structs.Login) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	admin := model.Admin{}
	tx := db.Begin()

	// 更新登入時間
	now := time.Now()
	expireTime := now.Add(time.Hour * 24 * 30)

	if err := tx.Model(&admin).Where("account=?", user.Account).Updates(model.Admin{LoginIP: user.LoginIP, LoginAt: now}).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("UPDATE_LOGIN_TIME_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("UPDATE_LOGIN_TIME_ERROR")
		return
	}

	// 建制新Session
	adSession := model.AdminSession{
		Account:    user.Account,
		Session:    user.Session,
		ExpireTime: expireTime,
	}

	if err := tx.Create(&adSession).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("CREATE_SESSION_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("CREATE_SESSION_ERROR")
		return
	}
	tx.Commit()

	return
}

// GetSession 取DB Session資料
func (a *AdminRepo) GetSession(session string) (adminSession model.AdminSession, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	defer db.Close()

	// 檢查帳號是否存在
	if rows := db.Where("session=?", session).First(&adminSession).RowsAffected; rows == 0 {
		go helper.WarnLog(fmt.Sprintf("SESSION_NOT_EXIST: %v", session))
		apiErr = errorcode.GetAPIError("SESSION_NOT_EXIST")
		return
	}

	return
}

// DeleteSession 刪除DB資料
func (a *AdminRepo) DeleteSession(session string) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	adminSession := model.AdminSession{}
	tx := db.Begin()

	if err := tx.Where("session=?", session).Delete(&adminSession).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("DELETE_SESSION_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("DELETE_SESSION_ERROR")
		return
	}

	tx.Commit()
	return
}

// CreateAdmin 新增後台管理員
func (a *AdminRepo) CreateAdmin(user *structs.ADRegister) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	// 組合資料
	admin := model.Admin{
		Account:   user.Account,
		Password:  user.Password,
		Status:    user.Status,
		GroupID:   user.GroupID,
		EditedBy:  user.EditedBy,
		CreatedAt: time.Now(),
	}

	tx := db.Begin()

	// 新增資料
	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("CREATE_ADMIN_USER_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("CREATE_ADMIN_USER_ERROR")
		return
	}

	tx.Commit()
	return
}

// MenuList 組合選單清單
func (a *AdminRepo) MenuList(admin *model.Admin) (lists []structs.MenuList) {

	list := structs.MenuList{}

	// 跑馬燈
	list.Name = "跑馬燈列表"
	list.Route = "announcement-list"
	lists = append(lists, list)

	// 影片
	list.Name = "影片列表"
	list.Route = "video"
	lists = append(lists, list)

	// 風格
	list.Name = "網站主題"
	list.Route = "theme"
	lists = append(lists, list)

	// 活動
	list.Name = "活動列表"
	list.Route = "activities"
	lists = append(lists, list)

	// 輪播圖
	list.Name = "輪播圖列表"
	list.Route = "rotate-list"
	lists = append(lists, list)

	// 標籤
	list.Name = "標籤管理"
	list.Route = "tag-list"
	lists = append(lists, list)

	if admin.GroupID == 1 {
		// 管理管理者帳號
		list.Name = "使用者管理"
		list.Route = "user-management"
		lists = append(lists, list)
	}

	return
}

// CheckPermission 檢查帳號權限
func (a *AdminRepo) CheckPermission(session string) (admin model.Admin, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	defer db.Close()

	// 檢查帳號是否存在
	db = db.Joins("JOIN `admin_session` ON `admin`.`account`=`admin_session`.`account`").Where("session=?", session)
	if rows := db.Where("`admin`.`group_id` =?", 1).First(&admin).RowsAffected; rows == 0 {
		apiErr = errorcode.GetAPIError("PERMISSION_DENIED")
		return
	}

	return
}

// UpdatePassword 修改密碼
func (a *AdminRepo) UpdatePassword(newPwd structs.UpdatePassword) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	admin := model.Admin{
		Password:  newPwd.Password,
		EditedBy:  newPwd.EditedBy,
		UpdatedAt: time.Now(),
	}

	tx := db.Begin()

	// 更新資料
	if err := tx.Model(&admin).Where("account = ?", newPwd.Account).Updates(&admin).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("UPDATE_PASSWORD_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("UPDATE_PASSWORD_ERROR")
		return
	}

	tx.Commit()
	return
}

// GetAdminInfoByID 透過管理者ID取詳細資料
func (a *AdminRepo) GetAdminInfoByID(adminID int) (admin model.Admin, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	defer db.Close()

	// 檢查帳號是否存在
	if err := db.Where("id=?", adminID).First(&admin).Error; err != nil {
		apiErr = errorcode.GetAPIError("GET_USER_ACCOUNT_ERROR")
		return
	}

	return
}

// DeleteAllSession 清除用戶登入的所有session
func (a *AdminRepo) DeleteAllSession(account string) (admin []model.AdminSession, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	tx := db.Begin()

	if err := db.Where("account = ? ", account).Find(&admin).Delete(&admin).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("DELETE_SESSION_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("DELETE_SESSION_ERROR")
		return
	}

	tx.Commit()

	return
}

// UpdateAdmin 更新管理者帳號權限
func (a *AdminRepo) UpdateAdmin(adminData map[string]interface{}) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	admin := model.Admin{}
	// 塞入編輯時間
	adminData["edited_at"] = time.Now()

	// 執行DB行為
	tx := db.Begin()

	if err := tx.Model(&admin).Omit("id").Where("id=?", adminData["id"]).Updates(adminData).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("UPDATE_ADMIN_DATA_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("UPDATE_ADMIN_DATA_ERROR")
		return
	}

	tx.Commit()

	return
}

// DeleteAdmin 刪除帳號資料
func (a *AdminRepo) DeleteAdmin(uid int) (apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	admin := model.Admin{}
	tx := db.Begin()

	if err := tx.Where("id=?", uid).Delete(&admin).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("DELETE_ADMIN_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("DELETE_ADMIN_ERROR")
		return
	}

	tx.Commit()

	return
}

// ClearExpiredSession 清除過期session
func (a *AdminRepo) ClearExpiredSession(time string) (session []model.AdminSession, apiErr errorcode.Error) {
	db, apiErr := model.DBConnection(global.GoFormatMa)
	if apiErr != nil {
		return
	}
	defer db.Close()

	tx := db.Begin()

	if err := tx.Where("expire_time < ?", time).Find(&session).Delete(&session).Error; err != nil {
		tx.Rollback()
		go helper.WarnLog(fmt.Sprintf("DELETE_EXPIRE_SESSION_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("DELETE_EXPIRE_SESSION_ERROR")
		return
	}

	tx.Commit()
	return
}
