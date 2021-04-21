package helper

import (
	"encoding/json"
	"goformat/app/global"
	"goformat/app/global/structs"
	"goformat/library/errorcode"
	"math/rand"
	"os"
	"time"

	"git.cchntek.com/rd3-pkg/teamgo"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

// Success 回傳成功API
func Success(result interface{}) *structs.APIResult {
	res := &structs.APIResult{
		ErrorCode:   1,
		ErrorMsg:    "SUCCESS",
		LogIDentity: "",
		Result:      result,
	}

	return res
}

// Fail 回傳失敗API
func Fail(err errorcode.Error) *structs.APIResult {
	res := &structs.APIResult{}

	res.ErrorCode = err.GetErrorCode()
	res.ErrorMsg = err.GetErrorText()
	res.LogIDentity = err.GetLogID()
	res.Result = map[string]string{}

	return res
}

// RanderStr 亂數產生字串
func RanderStr(length int) string {

	// 定義規則
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// ParseTime 轉換時間格式(string ---> time.Time)
func ParseTime(myTime string) (t time.Time, apiErr errorcode.Error) {
	var err error

	if myTime == "0000-00-00 00:00:00" {
		return
	}

	// 指定時區
	local, err := time.LoadLocation("Local")
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, "GET_TIME_ZONE_ERROR", err.Error())
		return
	}

	t, err = time.ParseInLocation("2006-01-02 15:04:05", myTime, local)
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, "PARSE_TIME_ERROR", err.Error())
		return
	}

	return
}

// StructToMap struct型態 轉 map型態 (For DB 使用)
func StructToMap(myStruct interface{}) (myMap map[string]interface{}, apiErr errorcode.Error) {

	// 轉形成map，才可以處理空字串,0,false值
	myMap = make(map[string]interface{})

	// 資料轉型
	byteData, err := json.Marshal(myStruct)
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, "JSON_MARSHAL_ERROR", err.Error())
		return
	}

	if err := json.Unmarshal(byteData, &myMap); err != nil {
		apiErr = ErrorHandle(global.WarnLog, "JSON_UNMARSHAL_ERROR", err.Error())
		return
	}

	return
}

// GrpcServerConnect grpc 連線建立
func GrpcServerConnect(address string) (conn *grpc.ClientConn, apiErr errorcode.Error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		apiErr := ErrorHandle(global.WarnLog, "CONNECT_GRPC_SERVER_FAILED", err.Error())
		return nil, apiErr
	}
	return

}

// TeamPlus 發送訊息至 teamplus
func TeamPlus(org string, content interface{}) {

	// 初始化
	teamInfo := teamgo.New()

	// 塞入資料
	teamInfo.TeamPlusURL = global.TeamPlusURL
	teamInfo.Account = global.TeamPlusAccount
	teamInfo.APIKey = global.TeamPlusAPIKey
	teamInfo.ChatSn = global.TeamPlusChatSn
	teamInfo.ENV = os.Getenv("ENV") + os.Getenv("PROJECT_NAME")
	teamInfo.Org = org

	byteData, _ := jsoniter.Marshal(content)
	teamInfo.Content = string(byteData)

	// 發送訊息
	if err := teamInfo.SendToGroup(); err != nil {
		_ = ErrorHandle(global.WarnLog, "SEND_MESSAGE_TO_TEAM_PLUS_ERROR", err)
	}
}
