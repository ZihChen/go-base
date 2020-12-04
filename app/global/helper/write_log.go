package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goformat/app/global"
	"goformat/library/errorcode"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AccessLogFormat 紀錄Log格式
type AccessLogFormat struct {
	Level       string      `json:"level"`        // Log 層級
	LogTime     string      `json:"logTime"`      // Log 當前時間
	ClientIP    string      `json:"clientIP"`     // 用戶IP
	Path        string      `json:"path"`         // 當前路徑
	Status      int         `json:"status"`       // 狀態碼
	Method      string      `json:"method"`       // GET,POST,PUT,DELETE
	Params      interface{} `json:"parameter"`    // 用戶帶入的參數
	HTTPReferer string      `json:"http_referer"` // 來源網址
}

// ErrorLogFormat 紀錄Log格式
type ErrorLogFormat struct {
	Level       string      `json:"level"`       // Log 層級
	LogIDentity string      `json:"logIDentity"` // Log 識別證
	LogTime     string      `json:"logTime"`     // Log 當前時間
	Path        string      `json:"path"`        // 當前路徑
	FuncName    string      `json:"funcname"`    // 發生錯誤的func名稱
	Params      interface{} `json:"parameter"`   // 錯誤發生時參數
	Result      interface{} `json:"reslut"`      // 錯誤訊息
}

// 宣告預設寫log路徑 + 格式至各環境 other.yaml 查詢
var fileName = ""
var filePath = ""

// ErrorHandle 取錯誤代碼 + 寫錯誤 Log
func ErrorHandle(errorType, errorCode string, errMsg interface{}, param ...interface{}) (apiErr errorcode.Error) {
	var logID string

	// New 一個 Error Interface
	apiErr = errorcode.NewError()

	// 塞入 Error 對應清單
	apiErr.SetErrorCode(errorCode)

	switch errorType {
	case global.WarnLog:
		logID = warnLog(fmt.Sprintf(errorCode+": %v", errMsg), param)
	default:
		logID = fatalLog(fmt.Sprintf(errorCode+": %v", errMsg), param)
	}

	// 存入 Log 識別證
	apiErr.SetLogID(logID)

	return
}

// AccessLog access.log
func AccessLog(c *gin.Context) {
	// 初始化
	content := AccessLogFormat{
		Level:       "[💚 START💚 ]",
		LogTime:     time.Now().Format("2006-01-02 15:04:05 -07:00"),
		ClientIP:    c.ClientIP(),
		Path:        c.Request.URL.Path,
		Status:      c.Writer.Status(),
		Method:      c.Request.Method,
		Params:      []string{},
		HTTPReferer: c.GetHeader("Referer"),
	}

	// 取檔案位置
	fileName = global.Config.Log.AccessLog
	filePath = global.Config.Log.LogDir

	// 檢查網址後方式否有帶入參數
	raw := c.Request.URL.RawQuery
	if raw != "" {
		content.Path = c.Request.URL.Path + "?" + c.Request.URL.RawQuery
	}

	// 檢查是否有method

	if c.Request.Method == "GET" {
		content.Params = c.Request.URL.RawQuery
	} else {
		_ = c.Request.ParseMultipartForm(1000)

		content.Params = c.Request.PostForm

		// 以 application/json 傳遞參數需用 GetRawData 接才有
		if len(c.Request.PostForm) < 1 {
			rd, _ := c.GetRawData()
			srd := string(rd)
			srd = strings.Replace(srd, " ", "", -1)
			srd = strings.Replace(srd, "\n", "", -1)
			srd = strings.Replace(srd, "\t", "", -1)
			content.Params = srd
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rd))
		}

		// 若參數有帶入密碼，將密碼換成「*」號
		if c.Request.PostForm.Get("pwd") != "" || c.Request.PostForm.Get("password") != "" {
			c.Request.PostForm.Set("pwd", "******")
			content.Params = c.Request.PostForm
		}
	}

	// 檢查路徑是否存在
	_ = CheckFileIsExist(filePath, fileName, global.DirPermission)

	// 型態轉換
	byteData, _ := json.Marshal(content)

	// 寫Log
	writeLog(byteData)
}

// fatalLog 組合error log內容，不可預期的錯誤才可使用
func fatalLog(err interface{}, param interface{}) string {
	content := &ErrorLogFormat{
		Level:       "[❌ Fatal❌ ]",
		LogIDentity: Md5EncryptionWithTime("identity"),
		LogTime:     time.Now().Format("2006-01-02 15:04:05 -07:00"),
		FuncName:    "",
		Path:        "",
		Params:      "",
		Result:      fmt.Sprintf("%v", err),
	}

	// 檢查是否需要紀錄帶入的參數
	content.Params = fmt.Sprintf("%v", param)

	// 取檔案位置
	fileName = global.Config.Log.ErrorLog
	filePath = global.Config.Log.LogDir

	// 檢查路徑是否存在
	_ = CheckFileIsExist(filePath, fileName, global.DirPermission)

	// 紀錄檔案名稱 + 行數 + func名稱
	getFilePath(6, content)

	// 型態轉換
	byteData, _ := json.Marshal(content)

	// 寫Log
	writeLog(byteData)

	return content.LogIDentity
}

// warnLog 組合warn log內容，可預期的錯誤才可使用
func warnLog(err interface{}, param interface{}) string {
	content := &ErrorLogFormat{
		Level:       "[⚠️ Warn ⚠️ ]",
		LogIDentity: Md5EncryptionWithTime(RanderStr(6)),
		LogTime:     time.Now().Format("2006-01-02 15:04:05 -07:00"),
		FuncName:    "",
		Path:        "",
		Params:      "",
		Result:      fmt.Sprintf("%v", err),
	}

	// 檢查是否需要紀錄帶入的參數
	content.Params = fmt.Sprintf("%v", param)

	// 取檔案位置
	fileName = global.Config.Log.ErrorLog
	filePath = global.Config.Log.LogDir

	// 檢查路徑是否存在
	_ = CheckFileIsExist(filePath, fileName, global.DirPermission)

	// 紀錄檔案名稱 + 行數 + func名稱
	getFilePath(3, content)

	// 型態轉換
	byteData, _ := json.Marshal(content)

	// 寫Log
	writeLog(byteData)

	return content.LogIDentity
}

// getFilePath 取檔案名稱 + 行數 + func名稱
func getFilePath(n int, content *ErrorLogFormat) {

	// 取檔案行數 + 路徑
	funcName, file, line, _ := runtime.Caller(n)
	f := runtime.FuncForPC(funcName)

	// 	組合資料
	content.FuncName = fmt.Sprintf("%v", f.Name())
	content.Path = fmt.Sprintf("%v:%d", file, line)

}

// writeLog 寫Log
func writeLog(logTxt []byte) {

	// 開啟檔案
	logFile, err := os.OpenFile(filePath+fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, global.FilePermission)
	if err != nil {
		log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n", err)
	}
	defer logFile.Close()

	// 寫入Log
	_, err = logFile.WriteString(string(logTxt) + "\n")
	if err != nil {
		log.Printf("❌ WriteLog: 寫檔案錯誤 [%v] ❌ \n", err)
	}
}
