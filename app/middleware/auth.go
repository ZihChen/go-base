package middleware

import (
	"GoFormat/app/business"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"GoFormat/app/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth 登入驗證
func Auth(c *gin.Context) {
	var account string
	var apiErr errorcode.Error
	// 取 cookie 的 session
	cookie, err := c.Cookie("session")
	if err != nil {
		apiErr = errorcode.GetAPIError("AUTH_VAILDATE_FAIL")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	// 檢查是否存在redis
	redis := &repository.Redis{}
	key := fmt.Sprintf("GoFormat:admin:%v", cookie)
	account, apiErr = redis.Get(key)
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	// 若redis沒有值，取DB
	if account == "" {
		adminBus := business.AdminBus{}
		account, apiErr = adminBus.CheckSessionExist(cookie)
		if apiErr != nil {
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			c.Abort()
			return
		}

	}

	c.Set("account", account)
	c.Set("session", cookie)
	c.Next()
}
