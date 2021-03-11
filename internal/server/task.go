package server

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"

	"github.com/robfig/cron/v3"
)

// Schedule 背景服務
func Schedule() {

	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			_ = helper.ErrorHandle(global.FatalLog, fmt.Sprintf("[❌ Fatal❌ ]: %v", err), "")
			// 錯誤時重新執行背景
			Schedule()
		}
	}()

	c := cron.New(cron.WithSeconds())

	// 給物件增加定時任務
	// _, err := c.AddFunc("* * * * * *", func() {
	// 	log.Println("hello world")
	// })

	// if err != nil {
	// 	_ = helper.ErrorHandle(global.WarnLog, "", err.Error())
	// }

	c.Start()

	// hook
	select {}
}
