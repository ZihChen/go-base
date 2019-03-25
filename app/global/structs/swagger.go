package structs

type orderGetArcadeGameOrderDetail struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string           `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	Result   arcadeGameResult `json:"result"`
}

type orderGetSlotGameOrderDetail struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string         `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	Result   slotGameResult `json:"result"`
}

type marqueeGetAutoMarqueeList struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []autoMarqueeList `json:"result"`
}

type adminLoginBody struct {
	Account  string `json:"account"  example:"admin"`
	Password string `json:"pwd"  example:"1234"`
}
type adminLogin struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []loginInfo `json:"result"`
}

type styleChooseStyleBody struct {
	// 風格ID
	StyleID int `json:"style_id" example:"1"`
}
type styleCreateStyleBody struct {
	// 風格標題
	Title string `json:"title" example:"風格標題一"`
	// 風格參數
	Content string `json:"content" example:"{\"SKIN_COLORS\": {\"PRIMARY_COLOR\": \"#102859\",\"SECONDARY_COLOR\": \"#2548a4\",\"TERTIARY_COLOR\": \"#147BDE\",\"QUATERNARY_COLOR\": \"#0b0e57\"},\"MAIN_COLORS\": {\"PRIMARY_COLOR\": \"#E57C00\",\"SECONDARY_COLOR\": \"#FFDE1C\",\"TERTIARY_COLOR\": \"#FDCB00\",},\"GRADIENT_COLORS\": {\"PRIMARY_COLOR\": {\"DARK\": \"SKIN_COLORS.PRIMARY_COLOR\",\"LIGHT\": \"SKIN_COLORS.TERTIARY_COLOR\"},\"SECONDARY_COLOR\": {    \"DARK\": \"SKIN_COLORS.SECONDARY_COLOR\",\"LIGHT\": \"SKIN_COLORS.QUATERNARY_COLOR\"}\"TERTIARY_COLOR\": {\"DARK\": \"SKIN_COLORS.TERTIARY_COLOR\",\"LIGHT\": \"SKIN_COLORS.SECONDARY_COLOR\"},\"QUATERNARY_COLOR\": {\"DARK\": \"MAIN_COLORS.PRIMARY_COLOR\",\"LIGHT\": \" MAIN_COLORS.SECONDARY_COLOR\"}},\"GRAY_COLORS\": {\"PRIMARY_COLOR\": \"#000000\",\"SECONDARY_COLOR\": \"#666666\",\"TERTIARY_COLOR\": \"#aaaaaa\",\"QUATERNARY_COLOR\": \"#d8d8d8\",\"QUINARY_COLOR\": \"#e9e9e9\",\"SENARY_COLOR\": \"#f2f2f2\",},\"SHADOW\": {\"PRIMARY_COLOR\": \"0 14px 28px 0 rgba(0, 0, 0, 0.19), 0 10px 10px 0 rgba(0, 0, 0, 0.26)\",\"SECONDARY_COLOR\": \"0 19px 38px 0 rgba(0, 0, 0, 0.3), 0 15px 12px 0 rgba(0, 0, 0, 0.22)\",\"TERTIARY_COLOR\": \"0 2px 4px 0 rgba(0, 0, 0, 0.4)\",},}"`
	// 風格類別(可從風格清單取類別ID)
	TagID int `json:"tag_id"  example:"2"`
}

type styleUpdateStyleSwitchBody struct {
	// 風格類別(可從風格清單取類別ID)
	TagID int `json:"tag_id"  example:"2"`
	// 風格ID
	StyleID int `json:"style_id" example:"1"`
}

type styleDeleteStyleBody struct {
	// 風格ID
	StyleID []int `json:"style_id" example:"1,2"`
}

type styleGetStyleListBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result getStyleListBack `json:"result"`
}

type styleGetStyleListFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []getStyleListFront `json:"result"`
}

type activeDeleteActiveBody struct {
	// 報導ID
	ActiveID []int `json:"active_id" example:"1,2"`
}
type activeCreateActiveBody struct {
	// 活動標題
	Title string `json:"title" example:"活動標題一(後台顯示用)"`
	// 活動網址
	URL string `json:"url" example:"https://youtu.be/7L06_HW_HcA"`
	// 活動各語系圖片
	Pic langString `json:"pic"`
}

type activeUpdateActiveBody struct {
	// 活動ID
	ID int `json:"id" example:"1"`
	// 活動標題
	Title string `json:"title" example:"活動標題一(後台顯示用)"`
	// 活動網址
	URL string `json:"url" example:"https://youtu.be/7L06_HW_HcA"`
	// 活動各語系圖片
	Pic langString `json:"pic"`
	// 活動開關 false: 停用, true:啟用
	Status bool `json:"status" example:"false"`
	// 系統商
	Site string `json:"site" example:"系統商A"`
}

type activeGetActiveListBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []getActiveListBack `json:"result"`
}

type activeGetActiveListFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []getActiveListFront `json:"result"`
}

type newsCreateNewsBody struct {
	// 標籤ID
	TagID int `json:"tag_id" example:"1"`
	// 標題
	Title string `json:"title" example:"標題一"`
	// 各語系內容
	Content newsDataLang `json:"content"`
	// 系統商
	Site string `json:"site" example:"系統商A"`
	// 報導開關 false: 停用, true:啟用(預設false)
	Status bool `json:"status" example:"true"`
	// 報導開始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 報導結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type newsUpdateNewsBody struct {
	// 報導ID
	NewsID int `json:"news_id" example:"1"`
	// 標籤ID
	TagID int `json:"tag_id" example:"1"`
	// 標題
	Title string `json:"title" example:"標題一"`
	// 各語系內容
	Content newsDataLang `json:"content"`
	// 系統商
	Site string `json:"site" example:"系統商A"`
	// 報導開關 false: 停用, true:啟用(預設false)
	Status bool `json:"status" example:"false"`
	// 報導開始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 報導結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type newsDeleteNewsBody struct {
	// 報導ID
	NewsID []int `json:"news_id" example:"1,2"`
}
type categoryDeleteTagBody struct {
	// 標籤ID
	TagID []int `json:"tag_id" example:"1,2"`
}
type categoryCreateTagBody struct {
	// 類別ID
	CategoryID int         `json:"category_id" example:"1"`
	Tag        []createTag `json:"tag"`
}

type categoryUpdateTagBody struct {
	// 類別ID
	CategoryID int         `json:"category_id" example:"1"`
	Tag        []updateTag `json:"tag"`
}

type categoryCreateCategoryBody struct {
	// 各語系名稱
	Name langString `json:"name"`
	// 管理權限
	GroupID int `json:"group_id"`
}

type categoryGetCategoryList struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []categoryList `json:"result"`
}

type marqueeGetAutoMarqueeSet struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []marqueeSet `json:"result"`
}

type marqueeSetAutoMarqueeSetBody struct {
	// 跑馬燈贏分值
	Winrate float64 `json:"winrate" example:"400.00"`
	// 跑馬燈贏分率
	Win float64 `json:"win" example:"10000"`
}

type marqueeCreateAutoMarqueeGameBody struct {
	// 遊戲gamecode列表
	Gamecode []string `json:"gamecode" example:"AB1,CQ9,192"`
}

type marqueeDeleteAutoMarqueeGameBody struct {
	// 遊戲gamecode列表
	Gamecode []string `json:"gamecode" example:"AB1,CQ9,192"`
}

type marqueeGetAutoMarqueeGame struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []string `json:"result" example:"[AB1,CQ9,158]"`
}

type categoryUpdateCategoryBody struct {
	// 類別ID
	ID int `json:"id"`
	// 各語系名稱
	Name langString `json:"name"`
	// 管理權限
	GroupID int `json:"group_id"`
}

type videoStatusVideoBody struct {
	// 宣傳片ID
	ID int `json:"id" example:"1"`
}

type categoryGetTagList struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []backendCategoryType `json:"result"`
}

type rotateGetRotateFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result rotateLang `json:"result"`
}

type rotateGetRotateBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []rotateBack `json:"result"`
}

type rotateCreateRotateBody struct {
	// 輪播圖標題(僅供後台顯示使用)
	Title string `json:"title" example:"輪播圖標題"`
	// 輪播圖狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 輪播圖排序
	Sort int `json:"sort" example:"5"`
	// 輪播圖圖片名稱
	Pic rotateLang `json:"pic"`
	// 輪播圖起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 輪播圖結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type rotateUpdateRotateBody struct {
	// 輪播圖標題(僅供後台顯示使用)
	Title string `json:"title" example:"輪播圖標題"`
	// 輪播圖狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 輪播圖排序
	Sort int `json:"sort" example:"5"`
	// 輪播圖圖片名稱
	Pic rotateLang `json:"pic"`
	// 輪播圖起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 輪播圖結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

// rotateSortRotateBody 輪播圖排序
type rotateSortRotateBody struct {
	// 輪播圖ID
	ID []int `json:"id" example:"5,8,1,9"`
}

type newsGetNewsFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []getNewsFront `json:"result"`
}

type newsGetNewsBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []newsDataBack `json:"result"`
}

type marqueeGetMarqueeIndex struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []marqueeIndexData `json:"result"`
}

type marqueeGetMarqueeListFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []marqueeListFront `json:"result"`
}

type marqueeCreateMarqueeBody struct {
	// 內容
	Data marqueeMarqueeData `json:"data"`
	// 連結
	URL string `json:"url" example:"https://youtu.be/7L06_HW_HcA"`
	// 類型ID
	TypeID int `json:"type_id" example:"1"`
	// 啟用狀態
	Status bool `json:"status" example:"true"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type marqueeUpdateMarqueeBody struct {
	// 內容
	Data marqueeMarqueeData `json:"data"`
	// 連結
	URL string `json:"url" example:"https://youtu.be/zGrYK1VTIjs"`
	// 類型ID
	TypeID int `json:"type_id" example:"1"`
	// 啟用狀態
	Status bool `json:"status" example:"true"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type marqueeGetMarqueeListBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []marqueeListBackData `json:"result"`
}

type marqueeGetMarqueeDetail struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []marqueeGetMarqueeDetailData `json:"result"`
}

type videoGetVideoFront struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []videoFront `json:"result"`
}

type videoGetVideoBack struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []videoBack `json:"result"`
}

type videoCreateVideoBody struct {
	// 名稱
	Name langString `json:"name"`
	// 網址
	URL string `json:"url"`
}

type videoUpdateVideoBody struct {
	// 名稱
	Name langString `json:"name"`
	// 網址
	URL string `json:"url"`
}

type rankGetTopGames struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result topGames `json:"result"`
}

type rankGetTopReward struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result topReward `json:"result"`
}

type rankGetTopLucky struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result topLucky `json:"result"`
}

type rankGetTopWin struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result topWin `json:"result"`
}

type adminIndex struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result adminList `json:"result"`
}

type adminRegisterBody struct {
	// 帳號
	Account string `json:"account" example:"user02"`
	// 密碼
	Password string `json:"pwd" example:"qwe123"`
	// 再次確認密碼
	PasswordConfirmation string `json:"pwd_confirmation" example:"qwe123"`
	// 帳號權限 0:一般,1:超級使用者
	GroupID int `json:"group_id" example:"1"`
}

type adminCategoryMenu struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result []menuList `json:"result"`
}

type updatePwdRepoBody struct {
	// 密碼
	Password string `json:"password" example:"qwe123"`
	// 再次確認密碼
	PasswordConfirmation string `json:"pwd_confirmation" example:"qwe123"`
}

type editAdminBody struct {
	// 帳號狀態 false:停用,true:啟用
	Status bool `json:"status" example:"false"`
	// 帳號權限 0:一般,1:超級使用者
	GroupID int `json:"group_id" example:"0"`
}

// adminEditAdmin 編輯管理者權限
type categoryCreateTag struct {
	// 錯誤代碼(成功為0)
	ErrorCode int `json:"error_code" example:"2000"`
	// 錯誤訊息(預設空字串)
	ErrorMsg string `json:"error_msg" example:"DB CONNECT ERRORR(2000)"`
	// 回傳資料
	Result createTagOption `json:"result"`
}

/*********************** 可共用struct ***********************/

type marqueeListFront struct {
	// 類別清單
	Category []categoryType `json:"category"`
	// 清單資料
	MarqueeList []marqueeListFrontData `json:"marquee_list"`
}

type categoryType struct {
	// 標籤ID
	TagID int `json:"tag_id" example:"1"`
	// 標籤名稱
	Name langString `json:"name"`
}

// CUCategoryOption 新增/修改 標籤
type createTagOption struct {
	// 類別ID
	CategoryID int `json:"category_id" example:"1"`
	// 標籤名稱
	Name langString
}

type langString struct {
	// 簡體
	CN string `example:"简体内容"`
	// 英文
	EN string `example:"English Content"`
	// 繁體
	TW string `example:"繁體內容"`
}

type cylangString struct {
	// 簡體
	CN string `json:"zh-cn" example:"简体内容"`
	// 英文
	EN string `json:"en" example:"English Content"`
	// 繁體
	TW string `json:"zh-tw" example:"繁體內容"`
	// 泰文
	TH string `json:"th" example:"繁體內容"`
}

type marqueeListFrontData struct {
	// 跑馬燈內容
	Data marqueeMarqueeData `json:"data"`
	// 跑馬燈類型
	Type categoryType `json:"type"`
	// 影片連結
	URL string `json:"url" example:"https://youtu.be/qZU63nWWS6M"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type marqueeContent struct {
	// 各語系標題
	Name string `json:"name" example:"各語系標題"`
}

type marqueeMarqueeData struct {
	// 簡體
	CN marqueeContent
	// 英文
	EN marqueeContent
	// 繁體
	TW marqueeContent
}

type marqueeListBackData struct {
	// ID
	ID int `json:"id" example:"2"`
	// 類型
	TypeID int `json:"type_id" example:"1"`
	// 啟用狀態
	Status bool `json:"status" example:"false"`
	// 跑馬燈內容
	Data marqueeMarqueeData `json:"data"`
	// 影片連結
	URL string `json:"url" example:"https://youtu.be/qZU63nWWS6M"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type marqueeListBackDetail struct {
	// ID
	ID int `json:"id" example:"1"`
	// 標題
	Title string `json:"title" example:"標題為後台清單顯示用"`
	// 內容
	Data map[string]MarqueeContent `json:"data"`
	// 類型
	Type int `json:"type" example:"2"`
	// 啟用狀態
	Status bool `json:"status" example:"false"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type getNewsFront struct {
	// 類別清單
	Category []categoryType `json:"category"`
	// 清單資料
	NewsList []newsDataFront `json:"news_list"`
}
type newsDataFront struct {
	// 流水號
	ID int `json:"id" example:"2"`
	// 標籤
	TagID categoryType `json:"tag_id"`
	// 內容
	Data newsDataLang `json:"Data"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type newsDataLang struct {
	// 簡體
	CN newsData
	// 英文
	EN newsData
	// 繁體
	TW newsData
}
type newsData struct {
	// 各語系標題
	Name string `json:"name" example:"報導名稱一"`
	// 各語系敘述
	Desc string `json:"desc" example:"報導描述...."`
	// 各語系內容
	Content string `json:"content" example:"報導html內容"`
	// 各語系圖片
	Pic string `json:"pic" example:"圖片連結"`
}

type topGames struct {
	// 熱門遊戲排行榜資料
	Hits []topGamesHit `json:"hits"`
	// 所有遊戲壓碼量總和
	TotalBets float64 `json:"total_bets"`
	// 排行榜開始區間
	StartTimestamp int `json:"start_timestamp"`
	// 排行榜結束區間
	EndTimestamp int `json:"end_timestamp"`
	// 排行榜資料總數
	NbHits int `json:"nb_hits"`
	// 使用降冪排列
	IsDesc bool `json:"is_desc"`
}

type topGamesHit struct {
	// 排名
	Rank int `json:"rank"`
	// 遊戲名稱
	GameName cylangString `json:"game_name"`
	// 遊戲代碼
	GameCode string `json:"game_code"`
	// 遊戲壓碼量
	Bets float64 `json:"bets"`
}

type topReward struct {
	// 遊戲大獎次數排行榜資料
	Hits []topRewardHit `json:"hits"`
	// 排行榜倍數
	X int `json:"x"`
	// 排行榜開始區間
	StartTimestamp int `json:"start_timestamp"`
	// 排行榜結束區間
	EndTimestamp int `json:"end_timestamp"`
	// 排行榜資料總數
	NbHits int `json:"nb_hits"`
	// 使用降冪排列
	IsDesc bool `json:"is_desc"`
}

type topRewardHit struct {
	// 排名
	Rank int `json:"rank"`
	// 遊戲名稱
	GameName cylangString `json:"game_name"`
	// 遊戲代碼
	GameCode string `json:"game_code"`
	// 出現倍率次數
	Count int `json:"count"`
}

type topLucky struct {
	// 玩家倍數排行榜資料
	Hits []topLuckyHit `json:"hits"`
	// 排行榜開始區間
	StartTimestamp int `json:"start_timestamp"`
	// 排行榜結束區間
	EndTimestamp int `json:"end_timestamp"`
	// 排行榜資料總數
	NbHits int `json:"nb_hits"`
	// 使用降冪排列
	IsDesc bool `json:"is_desc"`
}

type topLuckyHit struct {
	// 排名
	Rank int `json:"rank"`
	// 遊戲名稱
	GameName cylangString `json:"game_name"`
	// 遊戲代碼
	GameCode string `json:"game_code"`
	// 遊戲玩家帳號
	Account string `json:"account"`
	// 局號
	RoundID string `json:"round_id"`
	// 嬴分率 (win/bet)
	WinRate float64 `json:"win_rate"`
}

type topWin struct {
	// 玩家派彩排行榜資料
	Hits []topWinHit `json:"hits"`
	// 排行榜開始區間
	StartTimestamp int `json:"start_timestamp"`
	// 排行榜結束區間
	EndTimestamp int `json:"end_timestamp"`
	// 排行榜資料總數
	NbHits int `json:"nb_hits"`
	// 使用降冪排列
	IsDesc bool `json:"is_desc"`
}

type topWinHit struct {
	// 排名
	Rank int `json:"rank"`
	// 遊戲名稱
	GameName cylangString `json:"game_name"`
	// 遊戲代碼
	GameCode string `json:"game_code"`
	// 遊戲玩家帳號
	Account string `json:"account"`
	// 局號
	RoundID string `json:"round_id"`
	// 贏分
	Win float64 `json:"win"`
}

type rotateLang struct {
	// 電腦版
	PC langString
	// 手機版
	Mobile langString
	// 平板
	Tablet langString
}

type videoFront struct {
	// 詳細內容
	Name langString `json:"name"`
	// 影片連結
	URL string `json:"url" example:"https://youtu.be/kJD_cx7qQR8"`
}

type videoBack struct {
	// 影片流水號
	ID int `json:"id" example:"1"`
	// 詳細內容
	Name langString `json:"name"`
	// 影片連結
	URL string `json:"url" example:"https://youtu.be/kJD_cx7qQR8"`
	// 啟用狀態
	Status bool `json:"status" example:"false"`
	// 最後編輯人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 資料建立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 資料更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type adminList struct {
	// 流水號
	ID int `json:"id" example:"2"`
	// 帳號
	Account string `json:"account" example:"user01"`
	// 最後登入時間
	LoginAt string `json:"login_at" example:"2019-02-25 00:00:00"`
	// 帳號狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 帳號權限 0:一般,1:超級使用者
	GroupID int `json:"group_id" example:"1"`
	// 最後更改帳號人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 帳號建立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 帳號最後更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type menuList struct {
	// 功能ID
	ID string `json:"id" example:"1"`
	// 功能名稱
	Name string `json:"name" example:"跑馬燈"`
	// 功能代號
	Route string `json:"route" example:"marquee"`
}

type backendCategoryType struct {
	CategoryID   int               `json:"category_id"`   // 類別ID
	CategoryName map[string]string `json:"category_name"` // 類別名稱
	TagInfo      []tagInfo         `json:"tag"`           // 標籤資訊
}

type tagInfo struct {
	TagID int        `json:"tag_id"`
	Name  langString `json:"name"`
}

type marqueeGetMarqueeDetailData struct {
	// ID
	ID int `json:"id" example:"1"`
	// 內容
	Data marqueeMarqueeData `json:"data"`
	// 連結
	URL string `json:"url" example:"https://youtu.be/NJS1zCXp3Lo"`
	// 類型
	TypeID int `json:"type_id" example:"1"`
	// 啟用狀態
	Status bool `json:"status" example:"true"`
	// 起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
}

type marqueeIndexData struct {
	Data      marqueeMarqueeData `json:"data"`       // 內容
	Type      categoryType       `json:"type"`       // 類型
	URL       string             `json:"url"`        // 影片連結
	StartTime string             `json:"start_time"` // 起始時間
	EndTime   string             `json:"end_time"`   // 結束時間
}

type createTag struct {
	// 標籤名稱
	Name langString `json:"name"`
}

type updateTag struct {
	// 標籤ID
	TagID int `json:"tag_id" example:"1"`
	// 標籤名稱
	Name langString `json:"name"`
	// 標籤狀態
	Status bool `json:"status" example:"true"`
}

type newsDataBack struct {
	// 報導ID
	ID int `json:"id" example:"1"`
	// 標籤ID
	TagID int `json:"tag_id" example:"2"`
	// 報導標題(後台專用)
	Title string `json:"title" example:"報導標題，僅後台清單顯示"`
	// 報導內容(各語系名稱,各語系內容)
	Content string `json:"content"`
	// 系統商
	Site string `json:"site" example:"系統商A"`
	// 狀態
	Status bool `json:"status" example:"true"`
	// 報導起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 報導結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
	// 最後編輯人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 報導文件創立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 報導文件最後更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type getActiveListFront struct {
	// 活動流水號ID
	ID int `json:"id" example:"1"`
	// 活動網址
	URL string `json:"url" example:"http://chip-king.cq9promo.com/"`
	// 各語系圖片
	Pic langString `json:"pic"`
}

type getActiveListBack struct {
	// 活動流水號ID
	ID int `json:"id" example:"1"`
	// 活動標題(僅供後台顯示使用)
	Title string `json:"title" example:"活動標題"`
	// 活動網址
	URL string `json:"url" example:"http://chip-king.cq9promo.com/"`
	// 各語系圖片
	Pic langString `json:"pic"`
	// 活動狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 最後更改人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 活動建立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 活動最後更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type rotateBack struct {
	// 輪播圖流水號ID
	ID int `json:"id" example:"1"`
	// 輪播圖標題(僅供後台顯示使用)
	Title string `json:"title" example:"活動標題"`
	// 輪播圖狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 輪播圖排序
	Sort int `json:"sort" example:"5"`
	// 輪播圖圖片名稱
	Pic rotateLang `json:"pic"`
	// 最後更改人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 輪播圖起始時間
	StartTime string `json:"start_time" example:"2019-02-25 00:00:00"`
	// 輪播圖結束時間
	EndTime string `json:"end_time" example:"2019-02-25 00:00:00"`
	// 輪播圖建立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 輪播圖最後更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type getStyleListFront struct {
	// 風格流水號ID
	ID int `json:"id" example:"1"`
	// 風格參數
	Content string `json:"content"`
}

type getStyleListBack struct {
	// 風格1
	TagID1 []getStyleListBackData `json:"10"`
	// 風格2
	TagID2 []getStyleListBackData `json:"18"`
}

type getStyleListBackData struct {
	// 風格流水號ID
	ID int `json:"id" example:"1"`
	// 風格標題(僅供後台顯示使用)
	Title string `json:"title" example:"風格標題一"`
	// 風格參數
	Content string `json:"content"`
	// 風格狀態 flase:停用,true:啟用
	Status bool `json:"status" example:"true"`
	// 最後更改人員
	EditedBy string `json:"edited_by" example:"admin"`
	// 風格建立時間
	CreatedAt string `json:"created_at" example:"2019-02-25 00:00:00"`
	// 風格最後更新時間
	UpdatedAt string `json:"updated_at" example:"2019-02-25 00:00:00"`
}

type categoryList struct {
	// 分類ID
	ID int `json:"id" example:"1"`
	// 類別名稱
	Name langString `json:"name"`
	// 管理權限
	GroupID int `json:"group_id" example:"1"`
	// 最後更改人員
	EditedBy string `json:"edited_by" example:"admin"`
}

type loginInfo struct {
	// 帳號流水號ID
	ID int `json:"id" example:"1"`
	// 帳號
	Account string `json:"account" example:"admin"`
	// 權限
	Permission []string `json:"permission" example:"PROMOTION_ADMIN"`
}

type marqueeSet struct {
	// 跑馬燈贏分值
	Winrate float64 `json:"winrate" example:"400.00"`
	// 跑馬燈贏分率
	Win float64 `json:"win" example:"10000"`
}

type autoMarqueeList struct {
	// 跑馬燈類型
	Type string `json:"type" example:"win / winrate"`
	// 玩家帳號
	Account string `json:"account" example:"xb******7"`
	// 遊戲 ID
	Gamecode string `json:"gamecode" example:"5"`
	// 遊戲名稱
	GameName cylangString `json:"game_name"`
	// 時間
	Time string `json:"time"  example:"2019-03-15 14:57:00"`
	// 贏分
	Win float64 `json:"win" example:"9000"`
	// 贏分率
	Winrate float64 `json:"winrate" example:"100"`
	// 細單連結
	URL string `json:"url" example:"/orderDetail/slot?token=40dbafae68d8271bcdbdb43b83b35698"`
}

// rd1OrderDetailData 細單資料
type slotGameOrderDetail struct {
	// 下注贏分清單列表
	Actionlist []actionlist `json:"actionlist"`
	// 注單資料
	Detail detailWager `json:"detail"`
}

type detailWager struct {
	// 注單資料
	Wager wager `json:"wager"`
}

type actionlist struct {
	// 狀態
	Action string `json:"action" example:"bet/win"`
	// 金額
	Amount float64 `json:"amount" example:"0.9"`
	// 事件時間
	Eventtime string `json:"eventtime" example:"2017-05-12T06:53:53-04:00"`
}

type wager struct {
	// 場次/序號
	SeqNo int `json:"seq_no" example:"100114310"`
	// 結算時間
	EndTime string `json:"end_time" example:"2017-05-12T06:53:58-04:00"`
	// 玩家ID
	UserID string `json:"user_id" example:"58f7683c87f5150001abc810"`
	// 遊戲代號
	GameID int `json:"game_id" example:"1"`
	// 押注倍率
	PlayDenom string `json:"play_denom" example:"10"`
	// 幣別代號
	Currency string `json:"currency" example:"CNY"`
	// 下注時間
	StartTime string `json:"start_time" example:"2017-05-12T06:53:53-04:00"`
	// 伺服器IP
	ServerIP string `json:"server_ip" example:"192.168.3.35"`
	// 客端IP
	ClientIP string `json:"client_ip" example:"192.168.3.41"`
	// 押注金額
	PlayBet string `json:"play_bet" example:"0.9"`
	// 押注平台
	Platform string `json:"platform" example:"web"`
	// 下注量
	BetMultiple string `json:"bet_multiple" example:"50"`
	// RNG
	Rng []int `json:"rng" example:"3,44,39,35,37"`
	// 倍數
	Multiple string `json:"multiple" example:"0"`
	// basic game 贏分
	BaseGameWin string `json:"base_game_win" example:"0"`
	// 是否超過可贏額(0:否、1:是)
	WinOverLimitLock int `json:"win_over_limit_lock" example:"0"`
	// 遊戲型態
	GameType int `json:"game_type" example:"0"`
	// 贏分方式
	WinType int `json:"win_type" example:"0"`
	// 結算方式
	SettleType int `json:"settle_type" example:"0"`
	// 注單型態
	WagerType int `json:"wager_type"  example:"0"`
	// 總贏分
	TotalWin string `json:"total_win" example:"0"`
	// 贏分線數
	WinLineCount int `json:"win_line_count" example:"0"`
	// 錢包扣款交易ID
	BetTid string `json:"bet_tid" example:"rel-bet-100114310"`
	// 贏分交易ID
	WinTid string `json:"win_tid" example:"rel-win-100114310"`
	// 母單贏分畫面
	Proof wagerProof `json:"proof"`
	// 子單資料
	Sub []wagerSub `json:"sub"`
	// bonus game資料
	Pick []wagerPick `json:"pick"`
}

// wagerProof 母單贏分畫面
type wagerProof struct {
	// 贏分線資訊
	WinLineData []wagerProofWinLineData `json:"win_line_data"`
	// 圖標位置
	SymbolData []string `json:"symbol_data" example:"14 12 3 13 2,F 5 12 15 14,11 15 12 W 15"`
	// 額外畫面資訊
	ExtraData []int `json:"extra_data" example:"0"`
	// 圖標鎖定位置
	LockPosition        interface{} `json:"lock_position" example:""`
	ReelPosChg          []int       `json:"reel_pos_chg" example:"0"`
	ReelLenChange       interface{} `json:"reel_len_change"`
	ReelPay             interface{} `json:"reel_pay"`
	RespinReels         []int       `json:"respin_reels" example:"0,0,0,0,0"`
	BonusType           int         `json:"bonus_type" example:"0"`
	SpecialAward        int         `json:"special_award" example:"0"`
	SpecialSymbol       int         `json:"special_symbol" example:"0"`
	IsRespin            bool        `json:"is_respin" example:"false"`
	FgTimes             int         `json:"fg_times" example:"0"`
	FgRounds            int         `json:"fg_rounds" example:"0"`
	NextSTable          int         `json:"next_s_table" example:"0"`
	ExtendFeatureByGame interface{} `json:"extend_feature_by_game"`
}

// wagerSub 子單資料
type wagerSub struct {
	SubNo        int           `json:"sub_no" example:"10"`          // 子單流水號
	GameType     int           `json:"game_type" example:"50"`       // 遊戲型態
	Rng          []int         `json:"rng" example:"19,77,97,16,95"` // RNG
	Multiple     string        `json:"multiple" example:"6"`         // 倍數
	Win          string        `json:"win" example:"0"`              // 贏分
	WinLineCount int           `json:"win_line_count" example:"0"`   // 贏分線數
	WinType      int           `json:"win_type" example:"0"`         // 贏分方式
	Proof        wagerSubProof `json:"proof"`                        // 子單贏分畫面
}

// wagerSubProof 子單贏分畫面
type wagerSubProof struct {
	// 贏分線資訊
	WinLineData interface{} `json:"win_line_data"`
	// 圖標位置
	SymbolData []string `json:"symbol_data" example:"15 2 15 15 14,13 14 12 SC 11,3 13 2 13 4"`
	// 額外畫面資訊
	ExtraData []int `json:"extra_data" example:"0"`
	// 圖標鎖定位置
	LockPosition        interface{} `json:"lock_position"`
	ReelPosChg          []int       `json:"reel_pos_chg" example:"0"`
	ReelLenChange       interface{} `json:"reel_len_change"`
	ReelPay             interface{} `json:"reel_pay"`
	RespinReels         []int       `json:"respin_reels" example:"0,0,0,0,0"`
	BonusType           int         `json:"bonus_type" example:"0"`
	SpecialAward        int         `json:"special_award" example:"0"`
	SpecialSymbol       int         `json:"special_symbol" example:"0"`
	IsRespin            bool        `json:"is_respin" example:"false"`
	FgTimes             int         `json:"fg_times" example:"0"`
	FgRounds            int         `json:"fg_rounds" example:"0"`
	NextSTable          int         `json:"next_s_table" example:"0"`
	ExtendFeatureByGame interface{} `json:"extend_feature_by_game"`
}

// wagerPick bonus game資料
type wagerPick struct {
	PickNo   int            `json:"pick_no" example:"2"` // bonus game 流水號
	Multiple string         `json:"multiple" example:"6"`
	GameType int            `json:"game_type" example:"777"`
	Win      string         `json:"win" example:"0"` // 贏分
	Proof    wagerPickProof `json:"proof"`
}

// wagerPickProof Bonus Game贏分畫面
type wagerPickProof struct {
	SpinTimesOptions []int       `json:"spin_times_options" example:"10,11,9,14,18"` // 免費遊戲次數選項
	MultipleOptions  []int       `json:"multiple_options" example:"6,2,3,5,4"`       // 倍率選項
	WinOptions       interface{} `json:"win_options"`                                // 贏分選項
	ExtraOptions     interface{} `json:"extra_options"`                              // 額外選項
	PlayerSelected   []int       `json:"player_selected" example:"4,3"`              // 使用者選擇項目
	FgTimes          int         `json:"fg_times" example:"0"`
	FgRounds         int         `json:"fg_rounds" example:"0"`
	NextSTable       int         `json:"next_s_table" example:"0"`
	JpItemSelected   interface{} `json:"jp_item_selected"`
	JpItemLevel      interface{} `json:"jp_item_level"`
}

type arcadeGameOrderDetail struct {
	// 下注贏分清單列表
	Actionlist []actionlist `json:"actionlist"`
	// 注單資料
	Detail arcadeDetail `json:"detail"`
}

type arcadeDetail struct {
	// 注單資料
	Wager arcadeWager `json:"wager"`
}

type arcadeWager struct {
	// 場次/序號
	SeqNo int `json:"seq_no" example:"100114310"`
	// 結算時間
	EndTime string `json:"end_time" example:"2017-05-12T06:53:58-04:00"`
	// 玩家ID
	UserID string `json:"user_id" example:"58f7683c87f5150001abc810"`
	// 遊戲代號
	GameID int `json:"game_id" example:"1"`
	// 押注倍率
	PlayDenom string `json:"play_denom" example:"10"`
	// 幣別代號
	Currency string `json:"currency" example:"CNY"`
	// 下注時間
	StartTime string `json:"start_time" example:"2017-05-12T06:53:53-04:00"`
	// 伺服器IP
	ServerIP string `json:"server_ip" example:"192.168.3.35"`
	// 客端IP
	ClientIP string `json:"client_ip" example:"192.168.3.41"`
	// 押注金額
	PlayBet string `json:"play_bet" example:"0.9"`
	// 押注平台
	Platform string `json:"platform" example:"web"`
	// 下注量
	BetMultiple string `json:"bet_multiple" example:"50"`
	// RNG
	Rng []int `json:"rng" example:"3,44,39,35,37"`
	// 倍數
	Multiple string `json:"multiple" example:"0"`
	// basic game 贏分
	BaseGameWin string `json:"base_game_win" example:"0"`
	// 是否超過可贏額(0:否、1:是)
	WinOverLimitLock int `json:"win_over_limit_lock" example:"0"`
	// 遊戲型態
	GameType int `json:"game_type" example:"0"`
	// 贏分方式
	WinType int `json:"win_type" example:"0"`
	// 結算方式
	SettleType int `json:"settle_type" example:"0"`
	// 總贏分
	TotalWin string `json:"total_win" example:"0"`
	// 贏分線數
	WinLineCount int `json:"win_line_count" example:"0"`
	// 錢包扣款交易ID
	BetTid string `json:"bet_tid" example:"rel-bet-100114310"`
	// 贏分交易ID
	WinTid string           `json:"win_tid" example:"rel-win-100114310"`
	Proof  arcadeWagerProof `json:"proof"`
}

type arcadeWagerProof struct {
	BonusType int `json:"bonus_type" example:"0"`
	// 額外畫面資訊
	ExtraData          []int `json:"extra_data" example:"0"`
	PlayerBetMultiples []int `json:"player_bet_multiples"`
	SpecialAward       int   `json:"special_award" example:"0"`
	// 圖標位置
	SymbolData []string `json:"symbol_data" example:"15 2 15 15 14,13 14 12 SC 11,3 13 2 13 4"`
	// 圖標位置
	WinLineData []wagerProofWinLineData `json:"win_line_data"`
}

// wagerProofWinLineData 母單贏分畫面贏分線資訊
type wagerProofWinLineData struct {
	// 額外畫面資訊
	LineExtraData []int `json:"line_extra_data"  example:"0"`
	// 倍數
	LineMultiplier int `json:"line_multiplier" example:"6"`
	// 	金額
	LinePrize int `json:"line_prize" example:"7500"`
	// 贏分線型態
	LineType int `json:"line_type" example:"0"`
	// 贏分線型態
	NumOfKind int `json:"num_of_kind" example:"4"`
	// 圖標數
	SymbolCount int    `json:"symbol_count" example:"4"`
	SymbolID    string `json:"symbol_id" example:"13"`
	// 贏分線編號
	WinLineNo int `json:"win_line_no" example:"0"`
	// 贏分位置
	WinPosition [][]int `json:"win_position"`
}

type slotGameResult struct {
	Data   slotGameOrderDetail  `json:"data"`
	Status orderDetailAPIStatus `json:"status"`
}

type orderDetailAPIStatus struct {
	Code     string `json:"code"  example:"0"`                            // 錯誤代碼
	Message  string `json:"message" example:"Success"`                    // 錯誤訊息
	Datetime string `json:"datetime" example:"2019-03-12T02:47:03-04:00"` // 錯誤日期
}

type arcadeGameResult struct {
	Data   arcadeGameOrderDetail `json:"data"`
	Status orderDetailAPIStatus  `json:"status"`
}
