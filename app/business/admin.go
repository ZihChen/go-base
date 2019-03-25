package business

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"GoFormat/app/global/structs"
	"GoFormat/app/model"
	"GoFormat/app/repository"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

// AdminBus 管理者Business
type AdminBus struct {
}

var adminSingleton *AdminBus
var adminOnce sync.Once

// AdminIns 獲得單例對象
func AdminIns() *AdminBus {
	adminOnce.Do(func() {
		adminSingleton = &AdminBus{}
	})
	return adminSingleton
}

// UserInfo 用戶資訊
func (a *AdminBus) UserInfo(account string) (userInfo structs.Admin, apiErr errorcode.Error) {
	// 檢查使用者是否存在
	adminRepo := repository.AdminIns()
	admin, apiErr := adminRepo.GetAdmin(account)
	if apiErr != nil {
		return
	}

	// 組合資料
	userInfo.ID = admin.ID
	userInfo.Account = admin.Account
	userInfo.LoginIP = admin.LoginIP
	userInfo.LoginAt = admin.LoginAt.Format("2006-01-02 15:04:05")
	userInfo.Status = admin.Status
	userInfo.GroupID = admin.GroupID
	userInfo.EditedBy = admin.EditedBy
	userInfo.CreatedAt = admin.CreatedAt.Format("2006-01-02 15:04:05")
	userInfo.UpdatedAt = admin.UpdatedAt.Format("2006-01-02 15:04:05")

	return
}

// AdminList 取管理者清單
func (a *AdminBus) AdminList() (lists []structs.AdminList, apiErr errorcode.Error) {
	adminRepo := repository.AdminIns()

	// 取DB資料
	list, apiErr := adminRepo.AdminList()
	if apiErr != nil {
		return
	}

	// 組回傳資料
	listbyte, _ := json.Marshal(list)
	if err := json.Unmarshal(listbyte, &lists); err != nil {
		go helper.WarnLog(fmt.Sprintf("JSON_UNMARSHAL_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("JSON_UNMARSHAL_ERROR")
		return
	}

	// 轉換時間格式
	for k := range list {
		lists[k].LoginAt = list[k].LoginAt.Format("2006-01-02 15:04:05")
		lists[k].CreatedAt = list[k].CreatedAt.Format("2006-01-02 15:04:05")
		lists[k].UpdatedAt = list[k].UpdatedAt.Format("2006-01-02 15:04:05")

		if lists[k].UpdatedAt == "0001-01-01 00:00:00" {
			lists[k].UpdatedAt = ""
		}
	}
	return
}

// LoginRecord 紀錄登入紀錄
func (a *AdminBus) LoginRecord(c *gin.Context, user *structs.Login) (apiErr errorcode.Error) {
	// 檢查使用者是否存在
	adminRepo := repository.AdminIns()
	admin, apiErr := adminRepo.GetAdmin(user.Account)
	if apiErr != nil {
		return
	}

	// 找不到帳號資料，直接回傳帳密錯誤
	if admin.Account == "" {
		apiErr = errorcode.GetAPIError("USER_OR_PASSWORD_ERROR")
		return
	}

	// 檢查用戶是否停用(false:停用，true:啟用)
	if !admin.Status {
		apiErr = errorcode.GetAPIError("USER_ACCOUNT_DISABLE")
		return
	}

	// 比對密碼是否相同
	if ok := helper.CheckPasswordHash(user.Password, admin.Password); ok != true {
		apiErr = errorcode.GetAPIError("USER_OR_PASSWORD_ERROR")
		return
	}

	// session加密(亂數字串+帳號+毫秒時間)
	user.Session = helper.Md5EncryptionWithTime(user.Account)

	// 更新登入記入 + 新增一筆登入session紀錄
	apiErr = adminRepo.UpdateLoginRecord(user)
	if apiErr != nil {
		return
	}

	// 存入Redis
	redis := &repository.Redis{}
	key := fmt.Sprintf("GoFormat:login:%v", user.Session)
	redis.Set(key, user.Account, global.RedisLoginExpire)

	// 存入cookie
	c.SetCookie("session", user.Session, 2628000, "/", c.Request.Host, false, true) // 31天

	return
}

// CheckSessionExist middleware檢查
func (a *AdminBus) CheckSessionExist(session string) (account string, apiErr errorcode.Error) {
	adminRepo := repository.AdminIns()

	// 取DB Session
	admin, apiErr := adminRepo.GetSession(session)
	if apiErr != nil {
		return
	}

	// 比對 Session
	if admin.Session != session {
		apiErr = errorcode.GetAPIError("AUTH_VAILDATE_FAIL")
		return
	}

	return admin.Account, apiErr
}

// DeleteSession 登出後，清除DB Session 資料
func (a *AdminBus) DeleteSession(session string) (apiErr errorcode.Error) {
	adminRepo := repository.AdminIns()

	apiErr = adminRepo.DeleteSession(session)
	if apiErr != nil {
		return
	}

	return
}

// Create 建立管理者帳號
func (a *AdminBus) Create(newAdmin *structs.ADRegister) (apiErr errorcode.Error) {
	// 檢查帳號是否存在
	adminRepo := repository.AdminIns()
	admin, apiErr := adminRepo.GetAdmin(newAdmin.Account)
	if apiErr != nil {
		return
	}

	// 不到帳號資料，直接回傳帳密錯誤
	if admin.Account != "" {
		apiErr = errorcode.GetAPIError("USER_IS_EXIST")
		return
	}

	// 密碼加密
	newAdmin.Password, apiErr = helper.HashPassword(newAdmin.Password)
	if apiErr != nil {
		return
	}

	// 存入DB
	apiErr = adminRepo.CreateAdmin(newAdmin)
	if apiErr != nil {
		return
	}

	return
}

// MenuList 取首頁選單
func (a *AdminBus) MenuList(account string) (list []structs.MenuList, apiErr errorcode.Error) {
	// 取Session對應的用戶帳號
	adminRepo := repository.AdminIns()

	// 取管理者詳細資料
	admin, apiErr := adminRepo.GetAdmin(account)
	if apiErr != nil {
		return
	}

	// 取選單列表
	list = adminRepo.MenuList(&admin)

	return
}

// CheckPermission 檢查帳號權限
func (a *AdminBus) CheckPermission(session string) (account structs.AdminRepo, apiErr errorcode.Error) {
	// 進DB驗證資料
	adminRepo := repository.AdminIns()

	// 取DB Session
	admin, apiErr := adminRepo.CheckPermission(session)
	if apiErr != nil {
		return
	}

	byteData, err := json.Marshal(admin)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("JSON_MARSHAL_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("JSON_MARSHAL_ERROR")
		return
	}

	if err := json.Unmarshal(byteData, &account); err != nil {
		go helper.WarnLog(fmt.Sprintf("JSON_UNMARSHAL_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("JSON_UNMARSHAL_ERROR")
		return
	}

	return
}

// UpdatePassword 更改密碼
func (a *AdminBus) UpdatePassword(newPwd structs.UpdatePassword) (pwd structs.UpdatePwdRepo, apiErr errorcode.Error) {
	// 密碼加密
	newPwd.Password, apiErr = helper.HashPassword(newPwd.Password)
	if apiErr != nil {
		return
	}

	// 更改DB資料
	adminRepo := repository.AdminIns()
	if apiErr = adminRepo.UpdatePassword(newPwd); apiErr != nil {
		return
	}

	// 清除登入紀錄
	if apiErr = clearLoginRecord(newPwd.Account); apiErr != nil {
		return
	}

	return
}

// ResetPassword 將密碼改為隨機亂數產生的預設密碼
func (a *AdminBus) ResetPassword(newPwd structs.UpdatePassword) (pwd structs.UpdatePwdRepo, apiErr errorcode.Error) {
	// 取亂數字串
	newPwd.Password = helper.RanderStr(6)

	// 塞入回傳的struct
	pwd.Password = newPwd.Password

	// 取用戶資料
	adminRepo := repository.AdminIns()
	admin, apiErr := adminRepo.GetAdminInfoByID(newPwd.ID)
	if apiErr != nil {
		return
	}

	// 密碼加密
	newPwd.Password, apiErr = helper.HashPassword(newPwd.Password)
	if apiErr != nil {
		return
	}

	// 更新DB資料
	newPwd.Account = admin.Account
	if apiErr = adminRepo.UpdatePassword(newPwd); apiErr != nil {
		return
	}

	// 清除登入紀錄
	if apiErr = clearLoginRecord(newPwd.Account); apiErr != nil {
		return
	}

	return
}

// UpdateAdmin 更新管理者帳號權限
func (a *AdminBus) UpdateAdmin(admin structs.EditAdmin) (apiErr errorcode.Error) {
	var adminMap = make(map[string]interface{})

	// 將struct 轉換成map
	byteData, err := json.Marshal(admin)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("JSON_MARSHAL_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("JSON_MARSHAL_ERROR")
		return
	}

	if err := json.Unmarshal(byteData, &adminMap); err != nil {
		go helper.WarnLog(fmt.Sprintf("JSON_UNMARSHAL_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("JSON_UNMARSHAL_ERROR")
		return
	}

	adminRepo := repository.AdminIns()
	// 更新帳號資訊
	if apiErr = adminRepo.UpdateAdmin(adminMap); apiErr != nil {
		return
	}

	// 清除該用戶session資料
	adminInfo, apiErr := adminRepo.GetAdminInfoByID(admin.ID)
	if apiErr != nil {
		return
	}

	// 清除登入紀錄
	if apiErr = clearLoginRecord(adminInfo.Account); apiErr != nil {
		return
	}

	return
}

// DeleteAdmin 刪除管理者帳號
func (a *AdminBus) DeleteAdmin(uid int) (apiErr errorcode.Error) {
	// 取帳號資訊
	adminRepo := repository.AdminIns()
	admin, apiErr := adminRepo.GetAdminInfoByID(uid)
	if apiErr != nil {
		return
	}

	// 清除登入紀錄
	if apiErr = clearLoginRecord(admin.Account); apiErr != nil {
		return
	}

	// 刪除DB資料
	if apiErr = adminRepo.DeleteAdmin(uid); apiErr != nil {
		return
	}

	return
}

// ClearExpiredSession 清除過期session
func (a *AdminBus) ClearExpiredSession(time string) (apiErr errorcode.Error) {
	repo := repository.AdminIns()

	// 取DB Session 紀錄並清除
	var sessions []model.AdminSession
	sessions, apiErr = repo.ClearExpiredSession(time)
	if apiErr != nil {
		return
	}

	// 清除Redis紀錄
	redis := repository.Redis{}
	for _, v := range sessions {
		key := fmt.Sprintf("GoFormat:login:%v", v.Session)
		apiErr = redis.Delete(key)
		if apiErr != nil {
			return
		}
	}
	return
}

// clearLoginRecord 清除登入紀錄
func clearLoginRecord(account string) (apiErr errorcode.Error) {
	adminRepo := repository.AdminIns()

	// 取DB使用者全部的session並清除
	var sessions []model.AdminSession
	sessions, apiErr = adminRepo.DeleteAllSession(account)
	if apiErr != nil {
		return
	}

	// 清除Redis紀錄
	redis := repository.Redis{}
	for _, v := range sessions {
		key := fmt.Sprintf("GoFormat:login:%v", v.Session)

		if apiErr = redis.Delete(key); apiErr != nil {
			return
		}

	}

	return
}
