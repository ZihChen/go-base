package helper

import (
	"GoFormat/app/global"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LogFormat 紀錄Log格式
type LogFormat struct {
	Level    string        `json:"level"`
	LogTime  string        `json:"logTime"`
	ClientIP string        `json:"clientIP"`
	Path     string        `json:"path"`
	FileName string        `json:"filename"`
	Status   int           `json:"status"`
	Method   string        `json:"method"`
	Params   interface{}   `json:"params"`
	Result   interface{}   `json:"reslut"`
	ExecTime time.Duration `json:"execTime"`
}

// 宣告預設寫log路徑 + 格式
var fileName = "goformat_access.log"
var filePath = "/home/log/"
var wLog = &LogFormat{
	Level:    "Default",
	LogTime:  time.Now().Format("2006-01-02 15:04:05 -07:00"),
	ClientIP: "127.0.0.1",
	Path:     "",
	FileName: "",
	Status:   0,
	Method:   "",
	Params:   []string{},
	Result:   []string{},
}

// AccessLog access.log
func AccessLog() {
	// 取檔案位置
	fileName = global.Config.Log.AccessLog
	filePath = global.Config.Log.LogDir

	wLog.Level = "[💚 START💚 ]"

	// 檢查路徑是否存在
	CheckFileIsExist(filePath, fileName, 0755)

	// 寫Log
	writeLog()
}

// FatalLog 組合error log內容
func FatalLog(err interface{}) {

	// 取檔案位置
	fileName = global.Config.Log.ErrorLog
	filePath = global.Config.Log.LogDir

	// 組合Log內容
	wLog.Level = "[❌ Fatal❌ ]"

	// 檢查是否有回傳結果
	wLog.Result = fmt.Sprintf("%v", err)

	// 檢查路徑是否存在
	CheckFileIsExist(filePath, fileName, 0755)

	// 寫Log
	writeLog()
}

// PidLog 組合起服務 log 內容
func PidLog(pid int) {

	// 取檔案位置
	fileName = "pid.log"
	filePath = global.Config.Log.LogDir

	// 組合Log內容
	wLog.Level = "[💚 START💚 ]"

	// 檢查是否有回傳結果
	wLog.Result = fmt.Sprintf("Service pid is %v", pid)

	// 檢查路徑是否存在
	CheckFileIsExist(filePath, fileName, 0755)

	// 寫Log
	writeLog()
}

// WarnLog 組合warn log內容
func WarnLog(err interface{}) {
	// 取檔案位置
	fileName = global.Config.Log.ErrorLog
	filePath = global.Config.Log.LogDir

	// 組合Log內容
	wLog.Level = "[⚠️ Warn ⚠️ ]"
	wLog.Result = fmt.Sprintf("%v", err)

	// 檢查路徑是否存在
	CheckFileIsExist(filePath, fileName, 0755)

	// 寫Log
	writeLog()
}

// ComposeLog 組合Log內容
func ComposeLog(c *gin.Context) {

	wLog.LogTime = time.Now().Format("2006-01-02 15:04:05 -07:00")

	// 檢查是否有來源IP
	if c.ClientIP() != "" {
		wLog.ClientIP = c.ClientIP()
	}

	// 檢查是否有router路徑
	if c.Request.URL.Path != "" {
		wLog.Path = c.Request.URL.Path

		// 檢查網址後方式否有帶入參數
		raw := c.Request.URL.RawQuery
		if raw != "" {
			wLog.Path = c.Request.URL.Path + "?" + c.Request.URL.RawQuery
		}
	}

	// 檢查狀態碼
	if c.Writer.Status() != 0 {
		wLog.Status = c.Writer.Status()
	}

	// 檢查是否有method
	if c.Request.Method != "" {
		wLog.Method = c.Request.Method

		if c.Request.Method == "GET" {
			wLog.Params = c.Request.URL.RawQuery
		} else {
			c.Request.ParseMultipartForm(1000)

			wLog.Params = c.Request.PostForm

			// 若參數有帶入密碼，將密碼換成「*」號
			if c.Request.PostForm.Get("pwd") != "" || c.Request.PostForm.Get("password") != "" {
				c.Request.PostForm.Set("pwd", "******")
				wLog.Params = c.Request.PostForm
			}
		}
	}

}

// writeLog 寫Log
func writeLog() error {

	logTxt, err := json.Marshal(wLog)

	// 開啟檔案
	logFile, err := os.OpenFile(filePath+fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n", err)
		return nil
	}
	defer logFile.Close()

	// 寫入Log
	_, err = logFile.WriteString(string(logTxt) + "\n")
	if err != nil {
		log.Printf("❌ WriteLog: 寫檔案錯誤 [%v] ❌ \n", err)
		return nil
	}

	return nil
}
