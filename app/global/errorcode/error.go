package errorcode

import (
	"fmt"
)

// Error 自定義錯誤
type Error interface {
	Error() string
	GetErrorCode() int
	GetErrorText() string
}

// APIError API錯誤格式
type APIError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// GetAPIError 由錯誤碼取得API回傳
func GetAPIError(code string) APIError {
	if code == "" {
		return APIError{}
	}

	api, ok := errorCode[code]
	if !ok {
		return APIError{9999, fmt.Sprintf("Undefined Error (%s)", code)}
	}
	return APIError{api.ErrorCode, fmt.Sprintf(api.ErrorMsg+"(%d)", api.ErrorCode)}
}

// GetErrorCode 錯誤代碼
func (e APIError) GetErrorCode() int {
	return e.ErrorCode
}

// GetErrorText 錯誤訊息
func (e APIError) GetErrorText() string {
	return e.ErrorMsg
}

// Error API錯誤訊息
func (e APIError) Error() string {
	return fmt.Sprintf("%d: %v", e.ErrorCode, e.ErrorMsg)
}
