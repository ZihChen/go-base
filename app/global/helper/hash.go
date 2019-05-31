package helper

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Md5Encryption md5加密
func Md5Encryption(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)

	return md5Str
}

// Md5EncryptionWithTime md5 加密（加上奈秒時間）
func Md5EncryptionWithTime(str string) string {
	naTime := time.Now().UnixNano()
	data := str + strconv.FormatInt(naTime, 10)
	key := []byte(data)

	token := md5.Sum(key)
	md5Str := fmt.Sprintf("%x", token)

	return md5Str
}

// HashPassword 密碼加密(註冊管理者使用)
func HashPassword(password string) (value string, apiErr errorcode.Error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, "CRYPTION_ERROR", err.Error())
		return string(bytes), apiErr
	}

	return string(bytes), apiErr
}

// CheckPasswordHash 檢查檢查(登入使用))
func CheckPasswordHash(password, dbPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}
