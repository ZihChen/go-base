package errorcode

import (
	"fmt"
)

// Error 自定義錯誤
type Error interface {
	Error() string
	GetErrorCode() int
	GetErrorText() string
	GetLogID() string
	SetLogID(logID string)
	SetErrorCode(code string)
}

// newError API錯誤格式
type newError struct {
	ErrorCode       int    `json:"error_code"`
	ErrorMsg        string `json:"error_msg"`
	LogILogIDentity string `json:"logIDentity"`
}

// NewError 由錯誤碼取得API回傳
func NewError() Error {
	return &newError{}
}

// SetLogID 塞入 Log 識別證
func (e *newError) SetLogID(logID string) {
	e.LogILogIDentity = logID
}

// SetErrorCode 設定 errorcode
func (e *newError) SetErrorCode(code string) {

	api, ok := errorCode[code]
	if !ok {
		e.ErrorCode = 9999
		e.ErrorMsg = fmt.Sprintf("Undefined Error (%s)", code)
		e.LogILogIDentity = ""
	} else {
		e.ErrorCode = api.ErrorCode
		e.ErrorMsg = fmt.Sprintf(api.ErrorMsg+"(%d)", api.ErrorCode)
		e.LogILogIDentity = ""
	}
}

// GetErrorCode 錯誤代碼
func (e *newError) GetErrorCode() int {
	return e.ErrorCode
}

// GetErrorText 錯誤訊息
func (e *newError) GetErrorText() string {
	return e.ErrorMsg
}

// GetLogID Log身份
func (e *newError) GetLogID() string {
	return e.LogILogIDentity
}

// Error API錯誤訊息
func (e *newError) Error() string {
	return fmt.Sprintf("%d: %v", e.ErrorCode, e.ErrorMsg)
}
