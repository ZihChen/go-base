package helper

import (
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/structs"
	"encoding/json"
	"math/rand"
	"time"
)

// Success 回傳成功API
func Success(result interface{}) *structs.APIResult {
	res := &structs.APIResult{
		ErrorCode:   1,
		ErrorMsg:    "SUCCESS",
		LogIDentity: "",
		Result:      []string{},
	}

	if wLog.LogIDentity == "" {
		panic("LOG ID NOT EXIST")
	}

	res.LogIDentity = wLog.LogIDentity

	if result != "" && result != nil {
		res.Result = result
		return res
	}

	return res
}

// Fail 回傳失敗API
func Fail(err errorcode.Error) *structs.APIResult {
	res := &structs.APIResult{}

	res.ErrorCode = err.GetErrorCode()
	res.ErrorMsg = err.GetErrorText()
	res.LogIDentity = wLog.LogIDentity
	res.Result = []string{}

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

	t, err = time.Parse("2006-01-02 15:04:05", myTime)
	if err != nil {
		apiErr = errorcode.GetAPIError("Parse_Time_ERROR")
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
		apiErr = errorcode.GetAPIError("JSON_MARSHAL_ERROR")
		return
	}

	if err := json.Unmarshal(byteData, &myMap); err != nil {
		apiErr = errorcode.GetAPIError("JSON_UNMARSHAL_ERROR")
		return
	}

	return
}
