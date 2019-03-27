package service

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/structs"
	"sync"
)

// CypressSer Cypress API 專用
type CypressSer struct {
}

var cySingleton *CypressSer
var cyOnce sync.Once

// CyIns 獲得Rotate對象
func CyIns() *CypressSer {
	cyOnce.Do(func() {
		cySingleton = &CypressSer{}
	})
	return cySingleton
}

//GetTopGames 熱門遊戲排行榜
func (*CypressSer) GetTopGames() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/v1/ranking/topGames"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{
		"nbHits": 20,
	}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetTopReward 遊戲大獎次數排行榜
func (*CypressSer) GetTopReward(x int) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/v1/ranking/topReward"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{
		"x":      x,
		"nbHits": 20,
	}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetTopLucky 玩家倍數排行榜
func (*CypressSer) GetTopLucky() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/v1/ranking/topLucky"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{
		"nbHits": 20,
	}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetTopWin 玩家派彩排行榜
func (*CypressSer) GetTopWin() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/v1/ranking/topWin"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{
		"nbHits": 20,
	}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetGameList 取遊戲名稱列表
func (*CypressSer) GetGameList() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/v1/util/gameList"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetMarquee 跑馬燈清單
func (*CypressSer) GetMarquee(autoMarquee structs.MarqueeAPI) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/data"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{
		"startTimestamp": autoMarquee.StartTimestamp,
		"endTimestamp":   autoMarquee.EndTimestamp,
		"limit":          autoMarquee.Limit,
	}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//GetMarqueeSet 取得跑馬燈贏分條件與倍率條件
func (*CypressSer) GetMarqueeSet() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/set"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//SetMarqueeSet 設定跑馬燈贏分條件與倍率條件
func (*CypressSer) SetMarqueeSet(param map[string]interface{}) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/set"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//執行
	content, err := sendPut(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

// GetMarqueeGame 取跑馬燈遊戲
func (*CypressSer) GetMarqueeGame() (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/game"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//組參數
	param := map[string]interface{}{}

	//執行
	content, err := sendGet(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//CreateMarqueeGame 新增跑馬燈遊戲
func (*CypressSer) CreateMarqueeGame(param map[string]interface{}) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/game"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//執行
	content, err := sendPost(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}

//DeleteMarqueeGame 刪除跑馬燈遊戲
func (*CypressSer) DeleteMarqueeGame(param map[string]interface{}) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.CypressURL + "/promoweb/marquee/game/delete"

	//組 Header
	header := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": global.Config.API.CypressToken,
	}

	//執行
	content, err := sendPost(url, header, param)
	if err != nil {
		return nil, err
	}

	return content, apiErr
}
