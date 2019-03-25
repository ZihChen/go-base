package admin

import (
	"GoFormat/app/business"
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"GoFormat/app/global/structs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Index [後台]取管理者清單
// @Summary [後台]取管理者清單
// @Tags Admin
// @Produce  json
// @Success 200 {object} structs.adminIndex "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/admin [get]
func Index(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取DB資料
	adminBus := business.AdminIns()
	list, apiErr := adminBus.AdminList()
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(list))
}

// Login [後台]登入後台
// @Summary [後台]登入後台
// @Tags Other
// @Produce  json
// @Param body body structs.adminLoginBody true "登入功能"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/login [POST]
func Login(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 組參數
	user := structs.Login{
		LoginIP: c.ClientIP(),
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := errorcode.GetAPIError("BIND_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證參數規則
	if err := helper.ValidateStruct(&user); err != nil {
		apiErr := errorcode.GetAPIError("VAILDATE_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證DB資料 + 存session
	adminBus := business.AdminIns()
	apiErr := adminBus.LoginRecord(c, &user)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// Logout [後台]登出
// @Summary [後台]登出
// @Tags Other
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/logout [PUT]
func Logout(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 cookie 的 session
	session := c.MustGet("session").(string)

	// 砍DB資料
	adminBus := business.AdminIns()
	if apiErr := adminBus.DeleteSession(session); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 清除cookie
	c.SetCookie("session", session, -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, helper.Success(""))
}

// Register [後台]註冊新管理者
// @Summary [後台]註冊新管理者
// @description 註冊只能由Super Admin進行註冊
// @Tags Admin
// @Produce  json
// @Param body body structs.adminRegisterBody true "註冊新管理者"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/admin/register [POST]
func Register(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 Middleware 的 account
	account := c.MustGet("account").(string)

	// 組參數
	ad := structs.ADRegister{
		Status:   true,
		EditedBy: account,
	}

	if err := c.ShouldBindJSON(&ad); err != nil {
		apiErr := errorcode.GetAPIError("BIND_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證參數
	if err := helper.ValidateStruct(&ad); err != nil {
		apiErr := errorcode.GetAPIError("VAILDATE_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證帳號內容
	if ok := helper.ValidateRegex(ad.Account, global.AdminAccount); !ok {
		apiErr := errorcode.GetAPIError("ACCOUNT_RULE_ERROR")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證密碼內容
	if ok := helper.ValidateRegex(ad.Password, global.AdminPassword); !ok {
		apiErr := errorcode.GetAPIError("PASSWORD_RULE_ERROR")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 註冊
	adminBus := business.AdminIns()
	apiErr := adminBus.Create(&ad)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// CategoryMenu [後台]類別選單
// @Summary [後台]類別選單
// @Tags Other
// @Produce  json
// @Success 200 {object} structs.adminCategoryMenu "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/menu [GET]
func CategoryMenu(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 cookie 的 session
	account := c.MustGet("account").(string)

	// 取首頁選單
	adminBus := business.AdminIns()
	menuList, apiErr := adminBus.MenuList(account)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(menuList))
}

// UpdatePassword [後台]更新密碼
// @Summary [後台]更新密碼
// @description 只能更新自己的密碼
// @Tags Other
// @Produce  json
// @Param body body structs.updatePwdRepoBody true "修改自己的密碼"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/update_password [PUT]
func UpdatePassword(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 Middleware 值
	cookie := c.MustGet("session").(string)
	account := c.MustGet("account").(string)

	// 組合參數
	update := structs.UpdatePassword{}
	if err := c.ShouldBindJSON(&update); err != nil {
		apiErr := errorcode.GetAPIError("BIND_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	update.Account = account
	update.EditedBy = account

	// 驗證參數
	if err := helper.ValidateStruct(&update); err != nil {
		apiErr := errorcode.GetAPIError("VAILDATE_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證密碼內容
	if ok := helper.ValidateRegex(update.Password, global.AdminPassword); !ok {
		apiErr := errorcode.GetAPIError("PASSWORD_RULE_ERROR")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 更改密碼
	b := business.AdminIns()
	_, apiErr := b.UpdatePassword(update)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 清除cookie
	c.SetCookie("session", cookie, -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, helper.Success(""))
}

// ResetPassword [後台]重設密碼
// @Summary [後台]重設密碼
// @description 只能由Admin重設密碼 , 重設密碼會回傳隨機字串亂數
// @Tags Admin
// @Produce  json
// @Param id path int true "用戶ID"
// @Success 200 {object} structs.updatePwdRepoBody "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/admin/reset_password/{id} [PUT]
func ResetPassword(c *gin.Context) {
	var err error
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 Middleware 值
	account := c.MustGet("account").(string)

	// 初始化資料
	reset := structs.UpdatePassword{
		EditedBy: account,
	}

	// 取參數 + 轉型態
	id := c.Param("id")
	reset.ID, err = strconv.Atoi(id)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CHANGE_PARAMS_TYPE_FAIL: %v", err))
		apiErr := errorcode.GetAPIError("CHANGE_PARAMS_TYPE_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 更改密碼
	b := business.AdminIns()
	pwd, apiErr := b.ResetPassword(reset)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(pwd))
}

// EditAdmin [後台]更改管理者資訊
// @Summary [後台]更改管理者資訊
// @description 只能由 [超級使用者] 操作
// @Tags Admin
// @Produce  json
// @Param id path int true "用戶ID"
// @Param body body structs.editAdminBody true "編輯帳號權限"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/admin/edit_admin/{id} [PUT]
func EditAdmin(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取 Middleware 值
	account := c.MustGet("account").(string)

	// 初始化資料
	editAdmin := structs.EditAdmin{
		EditBy: account,
	}

	// 取參數 + 轉型態
	var err error
	id := c.Param("id")
	editAdmin.ID, err = strconv.Atoi(id)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CHANGE_PARAMS_TYPE_FAIL: %v", err))
		apiErr := errorcode.GetAPIError("CHANGE_PARAMS_TYPE_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	if err := c.ShouldBindJSON(&editAdmin); err != nil {
		apiErr := errorcode.GetAPIError("BIND_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證參數
	if err := helper.ValidateStruct(&editAdmin); err != nil {
		apiErr := errorcode.GetAPIError("VAILDATE_PARAMS_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 更新帳號資料
	adminBus := business.AdminIns()
	if apiErr := adminBus.UpdateAdmin(editAdmin); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// DeleteAdmin [後台]刪除管理者帳號
// @Summary [後台]刪除管理者帳號
// @description 只能由 [超級使用者] 操作
// @Tags Admin
// @Produce  json
// @Param id path int true "用戶ID"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/admin/delete_admin/{id} [DELETE]
func DeleteAdmin(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	// 取參數 + 轉型態
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CHANGE_PARAMS_TYPE_FAIL: %v", err))
		apiErr := errorcode.GetAPIError("CHANGE_PARAMS_TYPE_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 刪除帳號
	adminBus := business.AdminIns()
	if apiErr := adminBus.DeleteAdmin(userID); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// ClearExpiredSession 清除過期session
// @Summary 清除過期session
// @description 背景功能
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/clear_session [DELETE]
func ClearExpiredSession(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()
	adminBus := business.AdminIns()
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	if apiErr := adminBus.ClearExpiredSession(timeStr); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// GetUserInfo 取用戶資訊(登入後)
// @Summary 取用戶資訊(登入後)
// @Tags Other
// @Produce  json
// @Success 200 {object} structs.adminLogin "用戶資料"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/user_info [GET]
func GetUserInfo(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	account := c.MustGet("account").(string)

	// 驗證DB資料 + 存session
	adminBus := business.AdminIns()
	adminInfo, apiErr := adminBus.UserInfo(account)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 組合回傳資料
	resp := structs.LoginInfo{}
	resp.ID = adminInfo.ID
	resp.Account = adminInfo.Account
	switch adminInfo.GroupID {
	case 0:
		resp.Permission = append(resp.Permission, "PROMOTION_USER")
	case 1:
		resp.Permission = append(resp.Permission, "PROMOTION_ADMIN")
	default:
		resp.Permission = append(resp.Permission, "PROMOTION_USER")
	}

	c.JSON(http.StatusOK, helper.Success(resp))
}
