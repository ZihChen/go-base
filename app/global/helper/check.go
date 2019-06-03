package helper

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"syscall"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// CheckDirIsExist 檢查檔案路徑是否存在
func CheckDirIsExist(filePath string, perm os.FileMode) (apiErr errorcode.Error) {
	// 重新設置umask
	// syscall.Umask(0)

	// 檢查檔案路徑是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 建制資料夾
		if err := os.MkdirAll(filePath, perm); err != nil {
			apiErr = ErrorHandle(global.FatalLog, "CREATE_DIR_ERROR", err.Error())
			return
		}
	}

	return
}

// CheckFileIsExist 檢查檔案 + 路徑是否存在
func CheckFileIsExist(filePath, fileName string, perm os.FileMode) error {
	// 重新設置umask
	syscall.Umask(0)

	// 檢查檔案路徑是否存在
	if _, err := os.Stat(filePath + fileName); os.IsNotExist(err) {
		// 建制資料夾
		if err := os.MkdirAll(filePath, perm); err != nil {
			log.Printf("❌ WriteLog: 建立資料夾錯誤 [%v] ❌ \n", err.Error())
			return nil
		}

		//  建制檔案
		_, err := os.Create(filePath + fileName)
		if err != nil {
			log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n", err.Error())
			return nil
		}
	}

	return nil
}

// ValidateStruct 驗證struct規則
func ValidateStruct(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

// ValidateRegex 驗證字串正則規則
func ValidateRegex(req string, regex string) bool {
	if ok, _ := regexp.MatchString(regex, req); !ok {
		return false
	}

	return true
}

// InArray 檢測值是否在陣列內
func InArray(val string, array []string) (exists bool) {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

// CatchError 回傳不可預期的錯誤
func CatchError(c *gin.Context) {
	if err := recover(); err != nil {
		// 回傳不可預期的錯誤
		apiErr := ErrorHandle(global.FatalLog, fmt.Sprintf("%v", err), "")
		c.JSON(http.StatusBadRequest, Fail(apiErr))
	}
}
