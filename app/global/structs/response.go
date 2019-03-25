package structs

// VideoFront [前台] 宣傳影片
type VideoFront struct {
	Name map[string]string `json:"name"` // 詳細內容
	URL  string            `json:"url"`  // 影片連結
}

// VideoBack [後台] 宣傳影片
type VideoBack struct {
	ID        int               `json:"id"`         // 影片流水號
	Name      map[string]string `json:"name"`       // 詳細內容
	URL       string            `json:"url"`        // 影片連結
	Status    bool              `json:"status"`     // 啟用狀態
	EditedBy  string            `json:"edited_by"`  // 最後編輯人員
	CreatedAt string            `json:"created_at"` // 資料建立時間
	UpdatedAt string            `json:"updated_at"` // 資料更新時間
}

// GetNewsFront [前台]取最新報導
type GetNewsFront struct {
	Category []CategoryType  `json:"category"`  // 類別清單
	NewsList []NewsDataFront `json:"news_list"` // 清單資料
}

// NewsDataFront [前台]最新報導資料
type NewsDataFront struct {
	ID        int                 `json:"id"`         // 流水號
	Type      CategoryType        `json:"type"`       // 類別
	Data      map[string]NewsData `json:"Data"`       // 內容
	StartTime string              `json:"start_time"` // 起始時間
	EndTime   string              `json:"end_time"`   // 結束時間
	CreatedAt string              `json:"created_at"` // 資料建立時間
	UpdatedAt string              `json:"updated_at"` // 資料更新時間
}

// MarqueeListFront [前台]跑馬燈清單
type MarqueeListFront struct {
	Category    []CategoryType         `json:"category"`     // 類別清單
	MarqueeList []MarqueeListFrontData `json:"marquee_list"` // 清單資料
}

// MarqueeListFrontData [前台]跑馬燈清單資料
type MarqueeListFrontData struct {
	Data      map[string]MarqueeContent `json:"data"`       // 內容
	Type      CategoryType              `json:"type"`       // 類型
	URL       string                    `json:"url"`        // 影片連結
	StartTime string                    `json:"start_time"` // 起始時間
	EndTime   string                    `json:"end_time"`   // 結束時間
}

// CategoryType 類別清單
type CategoryType struct {
	TagID int               `json:"tag_id"` // 類別ID
	Name  map[string]string `json:"name"`   // 類別名稱
}

// BackendCategoryType 類別清單
type BackendCategoryType struct {
	Marquee CategoryDetail `json:"marquee"` // 跑馬燈
	News    CategoryDetail `json:"news"`    // 最新報導
	Style   CategoryDetail `json:"style"`   // 風格
	Games   CategoryDetail `json:"games"`   // 遊戲攻略
}

// CategoryDetail 類別內容
type CategoryDetail struct {
	CategoryID   int               `json:"category_id"`   // 類別ID
	CategoryName map[string]string `json:"category_name"` // 類別名稱
	TagInfo      []TagInfo         `json:"tag"`           // 標籤資訊
}

// TagInfo 標籤資訊
type TagInfo struct {
	TagID int               `json:"tag_id"`
	Name  map[string]string `json:"name"`
}

// MarqueeListBackData [後台]跑馬燈清單資料
type MarqueeListBackData struct {
	ID        int                       `json:"id"`         // ID
	TypeID    int                       `json:"type_id"`    // 類型
	Data      map[string]MarqueeContent `json:"data"`       // 內容
	URL       string                    `json:"url"`        // 影片連結
	Status    bool                      `json:"status"`     // 啟用狀態
	StartTime string                    `json:"start_time"` // 起始時間
	EndTime   string                    `json:"end_time"`   // 結束時間
}

// MarqueeListBackDetail [後台]跑馬燈詳細資料
type MarqueeListBackDetail struct {
	ID        int                       `json:"id"`         // ID
	Data      map[string]MarqueeContent `json:"data"`       // 內容
	URL       string                    `json:"url"`        // 連結
	TypeID    int                       `json:"type_id"`    // 類型
	Status    bool                      `json:"status"`     // 啟用狀態
	StartTime string                    `json:"start_time"` // 起始時間
	EndTime   string                    `json:"end_time"`   // 結束時間
}

// NewsDataBack 後台取清單資料
type NewsDataBack struct {
	ID        int                 `json:"id"`         // 報導ID
	TagID     int                 `json:"tag_id"`     // 標籤ID
	Title     string              `json:"title"`      // 報導標題(後台專用)
	Content   map[string]NewsData `json:"content"`    // 報導內容(各語系名稱,各語系內容)
	Site      string              `json:"site"`       // 系統商
	Status    bool                `json:"status"`     // 狀態
	StartTime string              `json:"start_time"` // 報導起始時間
	EndTime   string              `json:"end_time"`   // 報導結束時間
	EditedBy  string              `json:"edited_by"`  // 最後編輯人員
	CreatedAt string              `json:"created_at"` // 報導文件創立時間
	UpdatedAt string              `json:"updated_at"` // 報導文件最後更新時間
}

// GetActiveListBack [後台] 取活動清單
type GetActiveListBack struct {
	ID        int               `json:"id"`         // 活動流水號ID
	Title     string            `json:"title"`      // 活動標題(僅供後台顯示使用)
	URL       string            `json:"url"`        // 活動網址
	Pic       map[string]string `json:"pic"`        // 各語系圖片
	Status    bool              `json:"status"`     // 活動狀態 flase:停用,true:啟用
	EditedBy  string            `json:"edited_by"`  // 最後更改人員
	CreatedAt string            `json:"created_at"` // 活動建立時間
	UpdatedAt string            `json:"updated_at"` // 活動最後更新時間
}

// GetActiveListFront [前台] 取活動清單
type GetActiveListFront struct {
	ID  int               `json:"id"`  // 活動流水號ID
	URL string            `json:"url"` // 活動網址
	Pic map[string]string `json:"pic"` // 各語系圖片
}

// GetStyleListFront [前台] 風格個清單
type GetStyleListFront struct {
	ID      int    `json:"id"`      // 風格流水號ID
	Content string `json:"content"` // 風格參數
}

// GetStyleListBack [後台] 風格個清單
type GetStyleListBack struct {
	ID        int    `json:"id"`         // 風格流水號ID
	Title     string `json:"title"`      // 風格標題(僅供後台顯示使用)
	Content   string `json:"content"`    // 風格參數
	Status    bool   `json:"status"`     // 風格狀態 flase:停用,true:啟用
	EditedBy  string `json:"edited_by"`  // 最後更改人員
	CreatedAt string `json:"created_at"` // 風格建立時間
	UpdatedAt string `json:"updated_at"` // 風格最後更新時間
}

// GetCategoryList [後台]取類別清單
type GetCategoryList struct {
	ID       int               `json:"id"`        // 分類ID
	Name     map[string]string `json:"name"`      // 類別名稱
	GroupID  int               `json:"group_id"`  // 管理權限
	EditedBy string            `json:"edited_by"` // 最後更改人員
}

// LoginInfo 登入後資訊
type LoginInfo struct {
	ID         int      `json:"id"`         // 帳號流水號ID
	Account    string   `json:"account"`    // 帳號
	Permission []string `json:"permission"` // 權限
}

// AutoMarqueeSet 跑馬燈贏分條件與倍率條件
type AutoMarqueeSet struct {
	Winrate float64 `json:"winrate"` // 跑馬燈贏分率
	Win     float64 `json:"win"`     // 跑馬燈贏分
}

// AutoMarqueeList 前台自動跑馬燈清單
type AutoMarqueeList struct {
	Type     string            `json:"type"`      // 跑馬燈類型
	Account  string            `json:"account"`   // 玩家帳號
	Gamecode string            `json:"gamecode"`  // 遊戲 ID
	GameName map[string]string `json:"game_name"` // 遊戲名稱
	Time     string            `json:"time"`      // 時間
	GameType string            `json:"game_type"` // 遊戲類型
	Wins     float64           `json:"wins"`      // 贏分
	Winrate  float64           `json:"winrate"`   // 贏分率
	URL      string            `json:"url"`       // 細單連結
}
