package schedule

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/app/task"
	"goformat/library/errorcode"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type function_name func() (apiErr errorcode.Error)
type CronJob struct {
	Name     string          `json:"name"`      // 背景名稱
	Spec     string          `json:"spec"`      // 執行週期
	FuncName function_name   `json:"func_name"` // 函式名稱
	isRetry  bool            `json:"is_retry"`  // 重複執行
	EntryID  cron.EntryID    `json:"entry_id"`  // EntryID
	running  bool            `json:"running"`   // running
	mux      *sync.RWMutex   `json:"mux"`       // 讀寫鎖
	wg       *sync.WaitGroup `json:"wg"`        // 等待通道
}

var Singleton *CronJob
var Once sync.Once

// SeriesIns 獲得單例對象
func SeriesIns() *CronJob {
	Once.Do(func() {
		Singleton = &CronJob{}
	})
	return Singleton
}

// Run 自定義 Crob Job 接口
func (c *CronJob) Run() {
	// 加鎖，檢查是否可以執行背景
	c.mux.RLock()
	isRetry := c.isRetry
	running := c.running
	// 解鎖
	c.mux.RUnlock()

	// 如果不可重複則跳過
	if !isRetry {
		return
	}

	// 如果還在執行，則跳過
	if running {
		msg := fmt.Sprintf("%v 還在執行中", c.Name)
		_ = helper.ErrorHandle(global.WarnLog, "CRON_JOB_STILL_WORKING", msg)
		return
	}

	// todo 背景開關功能，撈 db 檢查

	// 開始前，基本設定
	c.wg.Add(1)

	// 將執行狀態改為 true
	c.mux.Lock()
	c.running = true
	c.mux.Unlock()
	startTime := time.Now()

	// 開始執行
	apiErr := c.Exec()
	// 執行後，基本設定
	endtime := time.Now()
	c.mux.Lock()
	c.running = false
	c.mux.Unlock()
	c.wg.Done()

	// 紀錄執行時間
	c.RecordJobStatus(c, startTime, endtime, apiErr)

}

// Init 初始化
func (c *CronJob) Init() {
	c.wg = new(sync.WaitGroup)
	c.mux = new(sync.RWMutex)
}

// Wait 等待 wg 結束
func (c *CronJob) Wait() {
	c.wg.Wait()
}

// Exec 開始執行背景
func (c *CronJob) Exec() (apiErr errorcode.Error) {
	if c.FuncName == nil {
		_ = helper.ErrorHandle(global.WarnLog, "FUNC_NOT_EXIST", c.FuncName)
	}

	return c.FuncName()
}

// SetEntryID 設定 pid
func (c *CronJob) SetEntryID(entryID cron.EntryID) {
	c.EntryID = entryID
}

// RecordJobStatus 背景執行完畢，會呼叫這個Func來紀錄執行狀態
func (c *CronJob) RecordJobStatus(job *CronJob, startTime, endTime time.Time, apiErr errorcode.Error) {

	msg := ""
	execTime := endTime.Sub(startTime)
	if apiErr != nil {
		msg = fmt.Sprintf("%v error, error reason %v , and totally spent %v", c.Name, apiErr.Error(), execTime)
		_ = helper.ErrorHandle(global.WarnLog, "CRON_JOB_ERROR", msg)
		return
	}

	msg = fmt.Sprintf("%v execute success, and totally spent %v", c.Name, execTime)
	_ = helper.ErrorHandle(global.SuccessLog, "CRON_JOB_SUCCESS_EXECUTE", msg)
}

// LoadSchedule 載入所有排程
func (c *CronJob) LoadSchedule() (jobs []*CronJob) {

	// 載入所有排程
	jobs = []*CronJob{
		// 範例
		// {
		// 	Name:     "印出 hello world", // 排程名稱
		// 	Spec:     "@every 10s",     // 排程時間
		// 	FuncName: task.HelloWorld,  // 對應的 func 名稱
		// 	isRetry:  true,             // 是否可重複執行
		// },
		{
			Name:     "印出 hi",
			Spec:     "@every 2s",
			FuncName: task.SayHi,
			isRetry:  true,
		},
	}

	return
}
