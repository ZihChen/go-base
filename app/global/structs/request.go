package structs

// MarqueeOption 跑馬燈清單
type MarqueeOption struct {
	ID        int    `json:"id"`                                         // 跑馬燈樓水號ID(後台專用)
	Platform  string `json:"platform" validate:"oneof=backend frontend"` // 前台或後台的篩選條件
	Page      int    `json:"page" form:"page"`                           // 當前頁數，0表示全部比數
	Count     int    `json:"count" form:"count"`                         // 每頁比數，0表示全部比數
	Type      int    `json:"type" form:"type"`                           // 顯示類型，0表示全部類型
	StartTime string `json:"start_time" form:"start_time"`               // 跑馬燈活動開始時間
	EndTime   string `json:"end_time" form:"end_time"`                   // 跑馬燈活動結束時間
}

// Login [後台]登入
type Login struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"pwd" validate:"required"`
	LoginIP  string `json:"-" validate:"required"`
	Session  string `json:"-"`
}

// ADRegister [後台]註冊新管理者
// Password 需與 PasswordConfirmation 相同
type ADRegister struct {
	Account              string `json:"account" validate:"required,min=2,max=30"`
	Password             string `json:"pwd" validate:"required,min=6,max=20,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"pwd_confirmation" validate:"required,min=6,max=20"`
	Status               bool   `json:"status" validate:"required"`
	GroupID              int    `json:"group_id" validate:"oneof=0 1"`
	EditedBy             string `json:"edited_by"`
}

// UpdatePassword [後台]變更密碼
type UpdatePassword struct {
	ID                   int    `json:"id"`
	Account              string `json:"account"`
	Password             string `json:"pwd" validate:"required,min=6,max=20,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"pwd_confirmation" validate:"required,min=6,max=20"`
	EditedBy             string `json:"edited_by"`
}

// EditAdmin [後台]編輯管理者
type EditAdmin struct {
	ID      int    `json:"id" validate:"required"`
	Status  bool   `json:"status"`
	GroupID int    `json:"group_id" validate:"oneof=0 1"`
	EditBy  string `json:"edited_by"`
}

// VideoOption 取影片的條件參數
type VideoOption struct {
	Sort string `json:"sort" form:"sort" validate:"oneof=id created_at updated_at"` // 資料排序方式
}

// NewsOption 報導清單
type NewsOption struct {
	Platform  string `json:"platform" validate:"oneof=backend frontend"` // 前台或後台的篩選條件
	Page      int    `json:"page" form:"page"`                           // 當前頁數，0表示全部比數
	Count     int    `json:"count" form:"count"`                         // 每頁比數，0表示全部比數
	TagID     int    `json:"tag_id" form:"tag_id"`                       // 顯示類型，0表示全部類型
	StartTime string `json:"start_time" form:"start_time"`               // 報導開始時間
	EndTime   string `json:"end_time" form:"end_time"`                   // 報導結束時間
}

// CategoryOption 類別清單
type CategoryOption struct {
	TagID      int `json:"tag_id" form:"tag_id"`           // Tag ID
	CategoryID int `json:"category_id" form:"category_id"` // 種類ID
}

// CUMarqueeOption 新增/修改跑馬燈
type CUMarqueeOption struct {
	ID        int                       `json:"id"`                             // 流水號
	Data      map[string]MarqueeContent `json:"data" validate:"required"`       // 名稱
	URL       string                    `json:"url"`                            // 網址
	StartTime string                    `json:"start_time" validate:"required"` // 起始時間
	EndTime   string                    `json:"end_time" validate:"required"`   // 結束時間
	Status    *bool                     `json:"status" validate:"required"`     // 啟用狀態
	TypeID    int                       `json:"type_id" validate:"required"`    // 分類
	EditedBy  string                    `json:"edit_by" validate:"required"`    // 最後編輯者
}

// CUVideoOption 新增/修改宣傳影片
type CUVideoOption struct {
	ID       int               `json:"id"`                             // 流水號
	Name     map[string]string `json:"name" validate:"required"`       // 名稱
	URL      string            `json:"url" validate:"required"`        // 網址
	EditedBy string            `json:"edited_by"  validate:"required"` // 最後編輯者
}

// CUTagOption 新增/修改標籤
type CUTagOption struct {
	CategoryID int           `json:"category_id" validate:"required"` // 種類ID
	Tag        []CategoryTag `json:"tag"`                             // 標籤資訊
	EditedBy   string        `json:"edited_by" validate:"required"`   // 最後編輯人員
}

// CategoryTag 標籤資料
type CategoryTag struct {
	TagID  int               `json:"tag_id"`                   // 標籤ID
	Name   map[string]string `json:"name" validate:"required"` // 標籤各語系名稱
	Status *bool             `json:"status"`                   // 標籤狀態
}

// DeleteTag 刪除標籤
type DeleteTag struct {
	TagID []int `json:"tag_id" validate:"required"` // 標籤ID
}

// BackendDeleteNews [後台]刪除報導
type BackendDeleteNews struct {
	NewsID []int `json:"news_id" validate:"required"` // 報導ID
}

// BackendCUNews [後台]新增報導
type BackendCUNews struct {
	NewsID int    `json:"news_id"`                    // 報導ID
	TagID  int    `json:"tag_id" validate:"required"` // 標籤ID
	Title  string `json:"title" validate:"required"`  // 標題
	// Content   map[string]map[string]string `json:"content" validate:"required"`    // 各語系內容
	Content   NewsDataLang `json:"content" validate:"required"`    // 各語系內容
	Site      string       `json:"site"`                           // 系統商
	Status    *bool        `json:"status"`                         // 報導開關 false: 停用, true:啟用
	StartTime string       `json:"start_time" validate:"required"` // 報導開始時間
	EndTime   string       `json:"end_time" validate:"required"`   // 報導結束時間
	EditedBy  string       `json:"edited_by" validate:"required"`  // 最後編輯人員
}

// NewsDataLang 報導各語系
type NewsDataLang struct {
	CN NewsDataReq `json:"cn"  validate:"required"` // 簡體
	EN NewsDataReq `json:"en"  validate:"required"` // 英文
	TW NewsDataReq `json:"tw"  validate:"required"` // 繁體
}

// NewsDataReq 報導各語系內容
type NewsDataReq struct {
	Name    string `json:"name" validate:"required"`    // 各語系標題
	Desc    string `json:"desc"`                        // 各語系敘述
	Content string `json:"content" validate:"required"` // 各語系內容
	Pic     string `json:"pic"`                         // 各語系圖片
}

// ActiveOption [後台]活動清單
type ActiveOption struct {
	ActiveID int    `json:"active_id" form:"active_id"`                                          // 活動ID
	Sort     string `json:"sort" form:"sort" validate:"required,oneof=id created_at updated_at"` // 活動排序方式
	Platform string `json:"platform" validate:"required,oneof=backend frontend"`                 // 前台或後台的篩選條件
}

// BackendDeleteActive [後台]刪除活動
type BackendDeleteActive struct {
	ActiveID []int `json:"active_id" validate:"required"` // 活動ID
}

// CURotateOption [後台]新增/更新輪播圖清單
type CURotateOption struct {
	ID        int                          `json:"id"`         // 輪播圖流水號ID
	Title     string                       `json:"title"`      // 輪播圖標題(僅供後台顯示使用)
	Status    *bool                        `json:"status"`     // 輪播圖狀態 flase:停用,true:啟用
	Sort      int                          `json:"sort"`       // 輪播圖排序
	Pic       map[string]map[string]string `json:"pic"`        // 輪播圖圖片名稱
	EditedBy  string                       `json:"edited_by"`  // 最後更改人員
	StartTime string                       `json:"start_time"` // 輪播圖起始時間
	EndTime   string                       `json:"end_time"`   // 輪播圖結束時間
}

// BackendCUActive [後台] 新增修改活動
type BackendCUActive struct {
	ID       int               `json:"id"`                            // 活動ID
	Title    string            `json:"title" validate:"required"`     // 活動標題
	URL      string            `json:"url" validate:"required"`       // 活動網址
	Pic      map[string]string `json:"pic" validate:"required"`       // 活動各語系圖片
	Status   *bool             `json:"status"`                        // 活動開關 false: 停用, true:啟用
	Site     string            `json:"site"`                          // 系統商
	EditedBy string            `json:"edited_by" validate:"required"` // 最後編輯人員
}

// BackendDeleteStyle 後台刪除風格
type BackendDeleteStyle struct {
	StyleID []int `json:"style_id" validate:"required"` // 風格ID
}

// BackendCreateStyle [後台] 新增修改風格
type BackendCreateStyle struct {
	TagID    int    `json:"tag_id" validate:"required"`    // 風格類別(可從風格清單取類別ID)
	Title    string `json:"title" validate:"required"`     // 風格標題
	Content  string `json:"content" validate:"required"`   // 風格參數
	Status   *bool  `json:"status"`                        // 風格開關 false: 停用, true:啟用
	EditedBy string `json:"edited_by" validate:"required"` // 最後編輯人員
}

// UpdateStyleSwitch [後台] 更改風格開關
type UpdateStyleSwitch struct {
	TagID    int    `json:"tag_id" validate:"required"`    // 風格類別(可從風格清單取類別ID)
	StyleID  int    `json:"style_id" validate:"required"`  // 風格ID
	EditedBy string `json:"edited_by" validate:"required"` // 最後編輯人員
}

// BackendChooseStyle [後台] 選擇風格
type BackendChooseStyle struct {
	StyleID  int    `json:"style_id" validate:"required"`  // 風格ID
	EditedBy string `json:"edited_by" validate:"required"` // 最後編輯人員
}

// VideoStatus [後台] 修改宣傳影片開關
type VideoStatus struct {
	ID       int    `json:"id" validate:"required"`        // 影片ID
	EditedBy string `json:"edited_by" validate:"required"` // 最後編輯人員
}

// CUCategory [後台] 新增/修改類別
type CUCategory struct {
	ID       int               `json:"id"`                            // 類別ID
	GroupID  *int              `json:"group_id" validate:"oneof=0 1"` // 權限種類
	Name     map[string]string `json:"name" validate:"required"`      // 各語系名稱
	EditedBy string            `json:"edited_by" validate:"required"` // 最後編輯人員
}

// MarqueeSetOption 跑馬燈贏分條件與倍率條件
type MarqueeSetOption struct {
	Winrate float64 `json:"winrate" validate:"required"` // 跑馬燈贏分率
	Win     float64 `json:"win" validate:"required"`     // 跑馬燈贏分
}

// MarqueeGameOption 跑馬燈遊戲設定
type MarqueeGameOption struct {
	Gamecode []string `json:"gamecode" validate:"required"` // 跑馬燈贏分率
}

// SortRotateOption 輪播圖排序
type SortRotateOption struct {
	// 輪播圖ID
	ID []int `json:"id" validate:"required"`
}
