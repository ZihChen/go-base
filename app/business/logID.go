package business

import (
	"goformat/app/global"
	"goformat/app/global/errorcode"
	"goformat/app/global/helper"
	"fmt"
	"sync"
)

// ErrBus Business邏輯
type ErrBus struct {
}

var errSingleton *ErrBus
var errOnce sync.Once

// ErrIns 獲得單例對象
func ErrIns() *ErrBus {
	errOnce.Do(func() {
		errSingleton = &ErrBus{}
	})
	return errSingleton
}

// GetErrorLog 取錯誤代碼
func (e *ErrBus) GetErrorLog() (apiErr errorcode.Error) {
	apiErr = helper.ErrorHandle(global.WarnLog, "PERMISSION_DENIE", "")

	// fmt.Println("====>", apiErr.GetErrorCode())
	// fmt.Println("====>", apiErr.GetErrorText())
	fmt.Println("====>", apiErr.GetLogID())
	return
}
