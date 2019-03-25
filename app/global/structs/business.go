package structs

import "time"

// NewsData 最新報導內容
type NewsData struct {
	Name    string `json:"name"`    // 各語系標題
	Desc    string `json:"desc"`    // 各語系敘述
	Content string `json:"content"` // 各語系內容
	Pic     string `json:"pic"`     // 各語系圖片
}

// MarqueeListBack [後台]跑馬燈清單
type MarqueeListBack struct {
	MarqueeList []MarqueeListBackData `json:"marquee_list"` // 清單資料
}

// MarqueeListDetail [後台]跑馬燈詳細內容
type MarqueeListDetail struct {
	ID        int                       `json:"id"`         // ID
	Title     string                    `json:"title"`      // 標題
	Data      map[string]MarqueeContent `json:"data"`       // 內容
	Type      int                       `json:"type"`       // 類型
	Status    bool                      `json:"status"`     // 啟用狀態
	StartTime string                    `json:"start_time"` // 起始時間
	EndTime   string                    `json:"end_time"`   // 結束時間
}

// CUMarquee [後台]新增/修改跑馬燈
type CUMarquee struct {
	Title     string                    `json:"title"`      // 標題
	Data      map[string]MarqueeContent `json:"data"`       // 內容
	TypeID    int                       `json:"type_id"`    // 類型ID
	Status    bool                      `json:"status"`     // 啟用狀態
	StartTime string                    `json:"start_time"` // 起始時間
	EndTime   string                    `json:"end_time"`   // 結束時間
}

// MarqueeContent 跑馬燈詳細內容
type MarqueeContent struct {
	Name string `json:"name"` // 各語系標題
}

// VideoBusiness 影片商業邏輯專用
type VideoBusiness struct {
	ID        int               `json:"id"`
	Name      map[string]string `json:"name"`
	URL       string            `json:"url"`
	Status    bool              `json:"status"`
	Site      string            `json:"site"`
	EditedBy  string            `json:"edited_by"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// TopGames 熱門遊戲排行榜
type TopGames struct {
	Hits           []TopGamesHit `json:"hits"`            // 熱門遊戲排行榜資料
	TotalBets      float64       `json:"total_bets"`      // 所有遊戲壓碼量總和
	StartTimestamp int           `json:"start_timestamp"` // 排行榜開始區間
	EndTimestamp   int           `json:"end_timestamp"`   // 排行榜結束區間
	NbHits         int           `json:"nb_hits"`         // 排行榜資料總數
	IsDesc         bool          `json:"is_desc"`         // 使用降冪排列
}

// TopGamesHit 熱門遊戲排行榜資料
type TopGamesHit struct {
	Rank     int               `json:"rank"`      // 排名
	GameName map[string]string `json:"game_name"` // 遊戲名稱
	GameCode string            `json:"game_code"` // 遊戲代碼
	Bets     float64           `json:"bets"`      // 遊戲壓碼量
}

// TopReward 遊戲大獎次數排行榜
type TopReward struct {
	Hits           []TopRewardHit `json:"hits"`            // 遊戲大獎次數排行榜資料
	X              int            `json:"x"`               // 排行榜倍數
	StartTimestamp int            `json:"start_timestamp"` // 排行榜開始區間
	EndTimestamp   int            `json:"end_timestamp"`   // 排行榜結束區間
	NbHits         int            `json:"nb_hits"`         // 排行榜資料總數
	IsDesc         bool           `json:"is_desc"`         // 使用降冪排列
}

// TopRewardHit 遊戲大獎次數排行榜資料
type TopRewardHit struct {
	Rank     int               `json:"rank"`      // 排名
	GameName map[string]string `json:"game_name"` // 遊戲名稱
	GameCode string            `json:"game_code"` // 遊戲代碼
	Count    int               `json:"count"`     // 出現倍率次數
}

// TopLucky 玩家倍數排行榜
type TopLucky struct {
	Hits           []TopLuckyHit `json:"hits"`            // 玩家倍數排行榜資料
	StartTimestamp int           `json:"start_timestamp"` // 排行榜開始區間
	EndTimestamp   int           `json:"end_timestamp"`   // 排行榜結束區間
	NbHits         int           `json:"nb_hits"`         // 排行榜資料總數
	IsDesc         bool          `json:"is_desc"`         // 使用降冪排列
}

// TopLuckyHit 玩家倍數排行榜資料
type TopLuckyHit struct {
	Rank     int               `json:"rank"`      // 排名
	GameName map[string]string `json:"game_name"` // 遊戲名稱
	GameCode string            `json:"game_code"` // 遊戲代碼
	Account  string            `json:"account"`   // 遊戲玩家帳號
	RoundID  string            `json:"round_id"`  // 局號
	WinRate  float64           `json:"win_rate"`  // 嬴分率 (win/bet)
}

// TopWin 玩家派彩排行榜
type TopWin struct {
	Hits           []TopWinHit `json:"hits"`            // 玩家派彩排行榜資料
	StartTimestamp int         `json:"start_timestamp"` // 排行榜開始區間
	EndTimestamp   int         `json:"end_timestamp"`   // 排行榜結束區間
	NbHits         int         `json:"nb_hits"`         // 排行榜資料總數
	IsDesc         bool        `json:"is_desc"`         // 使用降冪排列
}

// TopWinHit 玩家派彩排行榜資料
type TopWinHit struct {
	Rank     int               `json:"rank"`      // 排名
	GameName map[string]string `json:"game_name"` // 遊戲名稱
	GameCode string            `json:"game_code"` // 遊戲代碼
	Account  string            `json:"account"`   // 遊戲玩家帳號
	RoundID  string            `json:"round_id"`  // 局號
	Win      float64           `json:"win"`       // 贏分
}

// MenuList 後台功能選單
type MenuList struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
}

// AdminRepo 取管理者相關資訊
type AdminRepo struct {
	ID      int    `json:"id"`
	Account string `json:"account"`
}

// UpdatePwdRepo 回傳修改密碼（僅限制預設密碼可回傳）
type UpdatePwdRepo struct {
	Password string `json:"password"`
}

// AdminList 管理者清單
type AdminList struct {
	ID        int    `json:"id"`
	Account   string `json:"account"`
	LoginAt   string `json:"login_at"`
	Status    bool   `json:"status"`
	GroupID   int    `json:"group_id"`
	EditedBy  string `json:"edited_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// MarqueeList 跑馬燈清單
type MarqueeList struct {
	ID        int                       `json:"id"`
	Title     string                    `json:"title"`
	Content   map[string]MarqueeContent `json:"content"`
	URL       string                    `json:"url"`
	TypeID    int                       `json:"type_id"`
	Status    bool                      `json:"status"`
	Site      string                    `json:"site"`
	StartTime string                    `json:"start_time"`
	EndTime   string                    `json:"end_time"`
	EditedBy  string                    `json:"edited_by"`
	CreatedAt string                    `json:"created_at"`
	UpdatedAt string                    `json:"updated_at"`
}

// NewsList 報導清單
type NewsList struct {
	ID        int                 `json:"id"`
	Title     string              `json:"title"`
	Content   map[string]NewsData `json:"content"`
	TagID     int                 `json:"tag_id"`
	Status    bool                `json:"status"`
	Site      string              `json:"site"`
	StartTime string              `json:"start_time"`
	EndTime   string              `json:"end_time"`
	EditedBy  string              `json:"edited_by"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
}

// Category 種類清單
type Category struct {
	CategoryID   int               `json:"category_id"`
	CategoryName map[string]string `json:"category_name"`
	GroupID      int               `json:"group_id"`
	EditedBy     string            `json:"edited_by"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	TagDetail    []TagDetail       `json:"tag_detail"`
}

// TagDetail 標籤詳細資料
type TagDetail struct {
	TagID     int               `json:"tag_id"`
	TagName   map[string]string `json:"tag_name"`
	Status    bool              `json:"status"`
	EditedBy  string            `json:"edited_by"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// Active 管理者帳號
type Active struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	URL       string            `json:"url"`
	Pic       map[string]string `json:"pic"`
	Status    bool              `json:"status"`
	Site      string            `json:"site"`
	EditedBy  string            `json:"edited_by"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// CategoryTagBus 標籤資料
type CategoryTagBus struct {
	TagID  int    `json:"tag_id"` // 標籤ID
	Name   string `json:"name"`   // 標籤各語系名稱
	Status *bool  `json:"status"` // 標籤狀態
}

// CUNewsBus 後台新增修改報導
type CUNewsBus struct {
	NewsID    int       `json:"news_id"`    // 報導ID
	TagID     int       `json:"tag_id"`     // 標籤ID
	Title     string    `json:"title"`      // 報導標題
	Content   string    `json:"content"`    // 各語系內容
	Site      string    `json:"site"`       // 系統商
	Status    *bool     `json:"status"`     // 報導開關 false: 停用, true:啟用
	StartTime time.Time `json:"start_time"` // 報導開始時間
	EndTime   time.Time `json:"end_time"`   // 報導結束時間
	EditedBy  string    `json:"edited_by"`  // 最後編輯人員
}

// CUActiveBus 後台新增修改活動
type CUActiveBus struct {
	ID       int    `json:"id"`        // 活動ID
	Title    string `json:"title"`     // 活動標題
	URL      string `json:"url"`       // 活動網址
	Pic      string `json:"pic"`       // 各語系圖片
	Status   *bool  `json:"status"`    // 活動開關 false: 停用, true:啟用
	Site     string `json:"site"`      // 系統商
	EditedBy string `json:"edited_by"` // 最後編輯人員
}

// GetRotateBack [後台]輪播圖清單
type GetRotateBack struct {
	ID        int                          `json:"id"`         // 輪播圖流水號ID
	Title     string                       `json:"title"`      // 輪播圖標題(僅供後台顯示使用)
	Status    bool                         `json:"status"`     // 輪播圖狀態 flase:停用,true:啟用
	Sort      int                          `json:"sort"`       // 輪播圖排序
	Pic       map[string]map[string]string `json:"pic"`        // 輪播圖圖片名稱
	EditedBy  string                       `json:"edited_by"`  // 最後更改人員
	StartTime string                       `json:"start_time"` // 輪播圖起始時間
	EndTime   string                       `json:"end_time"`   // 輪播圖結束時間
	CreatedAt string                       `json:"created_at"` // 輪播圖建立時間
	UpdatedAt string                       `json:"updated_at"` // 輪播圖最後更新時間
}

// StyleBus 風格資料
type StyleBus struct {
	ID        int    `json:"id" gorm:"column:id"`
	TagID     int    `json:"tag_id" gorm:"column:tag_id"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content"`
	Status    *bool  `json:"status" gorm:"column:status"`
	EditedBy  string `json:"edited_by" gorm:"column:edited_by"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

// Admin 管理者資料
type Admin struct {
	ID        int    `json:"id"`
	Account   string `json:"account"`
	LoginIP   string `json:"login_ip"`
	LoginAt   string `json:"login_at"`
	Status    bool   `json:"status"`
	GroupID   int    `json:"group_id"`
	EditedBy  string `json:"edited_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CategoryFront 前台取類別清單
type CategoryFront struct {
	TagID     int               `json:"tag_id"`
	Category  CategoryKind      `json:"category"`
	Name      map[string]string `json:"name"`
	EditedBy  string            `json:"edited_by"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// CategoryKind 類別內容
type CategoryKind struct {
	CategoryID int               `json:"category_id"`
	Name       map[string]string `json:"name"`
}

// MarqueeAPI 取自動跑馬燈 API 需提供的參數
type MarqueeAPI struct {
	StartTimestamp int64 `json:"startTimestamp"`
	EndTimestamp   int64 `json:"endTimestamp"`
	Limit          int   `json:"limit"`
}
