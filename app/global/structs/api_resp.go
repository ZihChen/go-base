package structs

// CyRankTopGame 熱門遊戲排行榜
type CyRankTopGame struct {
	Data   cyRankTopGameData `json:"data"`
	Status cyStatus          `json:"status"`
}

// 熱門遊戲排行榜Hits
type cyRankTopGameHits struct {
	Rank     int               `json:"rank"`
	GameName map[string]string `json:"gameName"`
	GameCode string            `json:"gameCode"`
	Bets     float64           `json:"bets"`
}

// 熱門遊戲排行榜資料
type cyRankTopGameData struct {
	Hits           []cyRankTopGameHits `json:"hits"`
	TotalBets      float64             `json:"totalBets"`
	StartTimestamp int                 `json:"startTimestamp"`
	EndTimestamp   int                 `json:"endTimestamp"`
	NbHits         int                 `json:"nbHits"`
	IsDesc         bool                `json:"isDesc"`
}

// CyRankTopReward 遊戲大獎次數排行榜
type CyRankTopReward struct {
	Data   cyRankTopRewardData `json:"data"`
	Status cyStatus            `json:"status"`
}

// 遊戲大獎次數排行榜Hits
type cyRankTopRewardHits struct {
	Rank     int               `json:"rank"`
	GameName map[string]string `json:"gameName"`
	GameCode string            `json:"gameCode"`
	Count    int               `json:"count"`
}

// 遊戲大獎次數排行榜資料
type cyRankTopRewardData struct {
	Hits           []cyRankTopRewardHits `json:"hits"`
	X              int                   `json:"x"`
	StartTimestamp int                   `json:"startTimestamp"`
	EndTimestamp   int                   `json:"endTimestamp"`
	NbHits         int                   `json:"nbHits"`
	IsDesc         bool                  `json:"isDesc"`
}

// CyRankTopLucky 玩家倍數排行榜
type CyRankTopLucky struct {
	Data   cyRankTopLuckyData `json:"data"`
	Status cyStatus           `json:"status"`
}

// 玩家倍數排行榜Hits
type cyRankTopLuckyHits struct {
	Rank     int               `json:"rank"`
	GameName map[string]string `json:"gameName"`
	GameCode string            `json:"gameCode"`
	Account  string            `json:"account"`
	RoundID  string            `json:"roundID"`
	WinRate  float64           `json:"winRate"`
}

// 玩家倍數排行榜資料
type cyRankTopLuckyData struct {
	Hits           []cyRankTopLuckyHits `json:"hits"`
	StartTimestamp int                  `json:"startTimestamp"`
	EndTimestamp   int                  `json:"endTimestamp"`
	NbHits         int                  `json:"nbHits"`
	IsDesc         bool                 `json:"isDesc"`
}

// CyRankTopWin 玩家派彩排行榜
type CyRankTopWin struct {
	Data   cyRankTopWinData `json:"data"`
	Status cyStatus         `json:"status"`
}

// 玩家派彩排行榜Hits
type cyRankTopWinHits struct {
	Rank     int               `json:"rank"`
	GameName map[string]string `json:"gameName"`
	GameCode string            `json:"gameCode"`
	Account  string            `json:"account"`
	RoundID  string            `json:"roundID"`
	Win      float64           `json:"win"`
}

// 玩家派彩排行榜資料
type cyRankTopWinData struct {
	Hits           []cyRankTopWinHits `json:"hits"`
	StartTimestamp int                `json:"startTimestamp"`
	EndTimestamp   int                `json:"endTimestamp"`
	NbHits         int                `json:"nbHits"`
	IsDesc         bool               `json:"isDesc"`
}

// CyGameNameList 遊戲名稱清單
type CyGameNameList struct {
	Data   gameNameList `json:"data"`
	Status cyStatus     `json:"status"`
}

// CyMarqueeSet 跑馬燈贏分條件與倍率條件
type CyMarqueeSet struct {
	Data   interface{} `json:"data"`
	Status cyStatus    `json:"status"`
}

// CyMarqueeGame 跑馬燈遊戲設定
type CyMarqueeGame struct {
	Data   []string `json:"data"`
	Status cyStatus `json:"status"`
}

// CyCDMarqueeGame 跑馬燈遊戲設定(新增刪除)
type CyCDMarqueeGame struct {
	Data   string   `json:"data"`
	Status cyStatus `json:"status"`
}

// 回傳狀態
type cyStatus struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	TraceCode string `json:"traceCode"`
}

// 遊戲名稱清單
type gameNameList struct {
	Hits   []gameNameListHit `json:"hits"`
	NbHits int               `json:"nbHits"`
}

// 遊戲名稱清單資料
type gameNameListHit struct {
	GameCode string    `json:"gameCode"`
	GameHall string    `json:"gameHall"`
	GameType string    `json:"gameType"`
	NameSet  []nameSet `json:"nameSet"`
}

// 遊戲名稱
type nameSet struct {
	Name string `json:"name"`
	Lang string `json:"lang"`
}

// CyAutoMarqueeList 前台自動跑馬燈清單
type CyAutoMarqueeList struct {
	Data   []cyAmData `json:"data"`
	Status cyAmStatus `json:"status"`
}

// AmData 自動跑馬燈清單回傳資料
type cyAmData struct {
	Timestamp int     `json:"timestamp"` // 時間戳記，到奈秒
	Type      string  `json:"type"`      // 跑馬燈種類，win / winrate
	Account   string  `json:"account"`   // 帳號
	Roundid   string  `json:"roundid"`   // 單號
	GameType  string  `json:"gametype"`  // 遊戲類型
	Gamecode  string  `json:"gamecode"`  // 遊戲代號
	Wins      float64 `json:"wins"`      // 贏分
	Winrate   float64 `json:"winrate"`   // 贏分率
}

// CY API回傳狀態
type cyAmStatus struct {
	Code      string `json:"code"`      // 錯誤代碼
	Message   string `json:"message"`   // 錯誤訊息
	Datetime  string `json:"datetime"`  // 錯誤時間
	TraceCode string `json:"traceCode"` // 錯誤log id
}

// CyAutoMarqueeGame 自動跑馬燈遊戲
type CyAutoMarqueeGame struct {
	Data   []string `json:"data"`
	Status cyStatus `json:"status"`
}

// // Rd1SlotGameOrderDetail RD1 slot game細單資訊
// type Rd1SlotGameOrderDetail struct {
// 	Data   rd1OrderDetailData `json:"data"` // 細單資料
// 	Status rd1APIStatus       `json:"status"`
// }

// // rd1OrderDetailData 細單資料
// type rd1OrderDetailData struct {
// 	Account    string          `json:"-"`          // 帳戶
// 	Parentacc  string          `json:"-"`          // 代理帳號
// 	Actionlist []rd1Actionlist `json:"actionlist"` // 下注贏分清單列表
// 	Detail     rd1DetailWager  `json:"detail"`     // 注單資料
// }

// // rd1DetailWager 注單資料
// type rd1DetailWager struct {
// 	Wager rd1Wager `json:"wager"` // 注單資料
// }

// // rd1Actionlist 下注贏分清單列
// type rd1Actionlist struct {
// 	Action    string  `json:"action"`    // 狀態
// 	Amount    float64 `json:"amount"`    // 金額
// 	Eventtime string  `json:"eventtime"` // 事件時間
// }

// // rd1Wager 母單
// type rd1Wager struct {
// 	SeqNo            int            `json:"seq_no"`              // 場次/序號
// 	EndTime          string         `json:"end_time"`            // 結算時間
// 	UserID           string         `json:"user_id"`             // 玩家ID
// 	GameID           int            `json:"game_id"`             // 遊戲代號
// 	PlayDenom        string         `json:"play_denom"`          // 押注倍率
// 	Currency         string         `json:"currency"`            // 幣別代號
// 	StartTime        string         `json:"start_time"`          // 下注時間
// 	ServerIP         string         `json:"server_ip"`           // 伺服器IP
// 	ClientIP         string         `json:"client_ip"`           // 客端IP
// 	PlayBet          string         `json:"play_bet"`            // 押注金額
// 	Platform         string         `json:"platform"`            // 押注平台
// 	BetMultiple      string         `json:"bet_multiple"`        // 下注？
// 	Rng              []int          `json:"rng"`                 // RNG
// 	Multiple         string         `json:"multiple"`            // 倍數
// 	BaseGameWin      string         `json:"base_game_win"`       // basic game 贏分
// 	WinOverLimitLock int            `json:"win_over_limit_lock"` // 是否超過可贏額(0:否、1:是)
// 	GameType         int            `json:"game_type"`           // 遊戲型態
// 	WinType          int            `json:"win_type"`            // 贏分方式
// 	SettleType       int            `json:"settle_type"`         // 結算方式
// 	WagerType        int            `json:"wager_type"`          // 注單型態
// 	TotalWin         string         `json:"total_win"`           // 總贏分
// 	WinLineCount     int            `json:"win_line_count"`      // 贏分線數
// 	BetTid           string         `json:"bet_tid"`             // 錢包扣款交易ID
// 	WinTid           string         `json:"win_tid"`             // 贏分交易ID
// 	Proof            rd1WagerProof  `json:"proof"`               // 母單贏分畫面
// 	Sub              []rd1WagerSub  `json:"sub"`                 // 子單資料
// 	Pick             []rd1WagerPick `json:"pick"`                // bonus game資料
// }

// // rd1WagerProof 母單贏分畫面
// type rd1WagerProof struct {
// 	WinLineData         []rd1WagerProofWinLineData `json:"win_line_data"` // 贏分線資訊
// 	SymbolData          []string                   `json:"symbol_data"`   // 圖標位置
// 	ExtraData           []int                      `json:"extra_data"`    // 額外畫面資訊
// 	LockPosition        []interface{}              `json:"lock_position"` // 圖標鎖定位置
// 	ReelPosChg          []int                      `json:"reel_pos_chg"`
// 	ReelLenChange       []interface{}              `json:"reel_len_change"`
// 	ReelPay             []interface{}              `json:"reel_pay"`
// 	RespinReels         []int                      `json:"respin_reels"`
// 	BonusType           int                        `json:"bonus_type"`
// 	SpecialAward        int                        `json:"special_award"`
// 	SpecialSymbol       int                        `json:"special_symbol"`
// 	IsRespin            bool                       `json:"is_respin"`
// 	FgTimes             int                        `json:"fg_times"`
// 	FgRounds            int                        `json:"fg_rounds"`
// 	NextSTable          int                        `json:"next_s_table"`
// 	ExtendFeatureByGame []interface{}              `json:"extend_feature_by_game"`
// }

// // rd1WagerProofWinLineData 母單贏分畫面贏分線資訊
// type rd1WagerProofWinLineData struct {
// 	LineExtraData  []int   `json:"line_extra_data"`
// 	LineMultiplier int     `json:"line_multiplier"`
// 	LinePrize      int     `json:"line_prize"`
// 	LineType       int     `json:"line_type"`
// 	SymbolID       string  `json:"symbol_id"`
// 	SymbolCount    int     `json:"symbol_count"`
// 	NumOfKind      int     `json:"num_of_kind"`
// 	WinLineNo      int     `json:"win_line_no"`
// 	WinPosition    [][]int `json:"win_position"`
// }

// // rd1WagerSub 子單資料
// type rd1WagerSub struct {
// 	SubNo        int              `json:"sub_no"`         // 子單流水號
// 	GameType     int              `json:"game_type"`      // 遊戲型態
// 	Rng          []int            `json:"rng"`            // RNG
// 	Multiple     string           `json:"multiple"`       // 倍數
// 	Win          string           `json:"win"`            // 贏分
// 	WinLineCount int              `json:"win_line_count"` // 贏分線數
// 	WinType      int              `json:"win_type"`       // 贏分方式
// 	Proof        rd1WagerSubProof `json:"proof"`          // 子單贏分畫面
// }

// // rd1WagerSubProof 子單贏分畫面
// type rd1WagerSubProof struct {
// 	WinLineData         []interface{} `json:"win_line_data"` // 贏分線資訊
// 	SymbolData          []string      `json:"symbol_data"`   // 圖標位置
// 	ExtraData           []int         `json:"extra_data"`    // 額外畫面資訊
// 	LockPosition        []interface{} `json:"lock_position"` // 圖標鎖定位置
// 	ReelPosChg          []int         `json:"reel_pos_chg"`
// 	ReelLenChange       []interface{} `json:"reel_len_change"`
// 	ReelPay             []interface{} `json:"reel_pay"`
// 	RespinReels         []int         `json:"respin_reels"`
// 	BonusType           int           `json:"bonus_type"`
// 	SpecialAward        int           `json:"special_award"`
// 	SpecialSymbol       int           `json:"special_symbol"`
// 	IsRespin            bool          `json:"is_respin"`
// 	FgTimes             int           `json:"fg_times"`
// 	FgRounds            int           `json:"fg_rounds"`
// 	NextSTable          int           `json:"next_s_table"`
// 	ExtendFeatureByGame []interface{} `json:"extend_feature_by_game"`
// }

// // rd1WagerPick bonus game資料
// type rd1WagerPick struct {
// 	PickNo   int               `json:"pick_no"` // bonus game 流水號
// 	Multiple string            `json:"multiple"`
// 	GameType int               `json:"game_type"`
// 	Win      string            `json:"win"` // 贏分
// 	Proof    rd1WagerPickProof `json:"proof"`
// }

// // rd1WagerPickProof Bonus Game贏分畫面
// type rd1WagerPickProof struct {
// 	SpinTimesOptions []int         `json:"spin_times_options"` // 免費遊戲次數選項
// 	MultipleOptions  []int         `json:"multiple_options"`   // 倍率選項
// 	WinOptions       []interface{} `json:"win_options"`        // 贏分選項
// 	ExtraOptions     []interface{} `json:"extra_options"`      // 額外選項
// 	PlayerSelected   []int         `json:"player_selected"`    // 使用者選擇項目
// 	FgTimes          int           `json:"fg_times"`
// 	FgRounds         int           `json:"fg_rounds"`
// 	NextSTable       int           `json:"next_s_table"`
// 	JpItemSelected   []interface{} `json:"jp_item_selected"`
// 	JpItemLevel      []interface{} `json:"jp_item_level"`
// }

// Rd1OrderDetail RD1 slot game細單資訊
type Rd1OrderDetail struct {
	Data   interface{}  `json:"data"` // 細單資料
	Status rd1APIStatus `json:"status"`
}

// rd1APIStatus RD1 API 狀態
type rd1APIStatus struct {
	Code     string `json:"code"`     // 錯誤代碼
	Message  string `json:"message"`  // 錯誤訊息
	Datetime string `json:"datetime"` // 錯誤日期
}
