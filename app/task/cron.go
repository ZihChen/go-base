package task

import (
	"fmt"
	"goformat/app/global/helper"
	"goformat/library/errorcode"
	"time"
)

func HelloWorld() (apiErr errorcode.Error) {
	str := helper.Md5EncryptionWithTime("hdsk")
	fmt.Println(str + " Hello Custom Job")
	time.Sleep(time.Second * 30)
	fmt.Println(str + " After 30 Second Later in Hello Custom Job")
	return
}

func SayHi() (apiErr errorcode.Error) {
	str := helper.Md5EncryptionWithTime("hdsk")
	fmt.Println("now ===>", time.Now())
	fmt.Println(str + " HI Custom Job")
	time.Sleep(time.Second * 10)
	fmt.Println(str + " After 10 Second Later in HI Custom Job")
	return
}
