package service

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"sync"
)

// RD1Ser RD1 API 專用
type RD1Ser struct {
}

var rd1Singleton *RD1Ser
var rd1Once sync.Once

// RD1Ins 獲得Rotate對象
func RD1Ins() *RD1Ser {
	rd1Once.Do(func() {
		rd1Singleton = &RD1Ser{}
	})
	return rd1Singleton
}

// GetOrderDetail 取細單(SlotGame + 街機)
func (*RD1Ser) GetOrderDetail(roundID string) (content []byte, apiErr errorcode.Error) {
	//API位置
	url := global.Config.API.RD1URL + "/promote/v1/db/wager/" + roundID

	//組 Header
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
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
