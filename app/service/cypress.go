package service

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
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
