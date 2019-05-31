package task

import (
	"GoFormat/app/global"
	"GoFormat/app/global/helper"
	"fmt"

	"github.com/robfig/cron"
)

// Schedule 背景服務
func Schedule() {

	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.WarnLog, fmt.Sprintf("[❌ Fatal❌ ]: %v", err), "")
			// 錯誤時重新執行背景
			Schedule()
		}
	}()

	c := cron.New()

	// 刪除過期session
	// delSidTime := "0 0 3 */1 * *"
	// c.AddFunc(delSidTime, func() {
	// 	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// 	aBus := business.AdminBus{}
	// 	aBus.ClearExpiredSession(timeStr)
	// })

	// 刪除多餘圖片
	// delImgTime := "0 0 4 * */1 *"
	// c.AddFunc(delImgTime, func() {
	// 	aBus := business.DeleteImage{}
	// 	aBus.Delete()
	// })

	c.Start()
}
