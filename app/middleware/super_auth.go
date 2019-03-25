package middleware

import (
	"GoFormat/app/business"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuperAuth 檢查管理者帳號是否為超級使用者
func SuperAuth(c *gin.Context) {
	// 取 cookie 的 session
	cookie, err := c.Cookie("session")
	if err != nil {
		apiErr := errorcode.GetAPIError("AUTH_VAILDATE_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	// 檢查帳號權限
	adminBus := business.AdminBus{}
	_, apiErr := adminBus.CheckPermission(cookie)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	c.Next()
}
