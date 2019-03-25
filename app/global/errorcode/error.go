package errorcode

import (
	"fmt"
)

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

// Error API錯誤訊息
func (e *APIError) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorCode, e.ErrorMsg)
}
