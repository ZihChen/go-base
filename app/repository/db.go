package repository

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/model"
	"fmt"
	"sync"
	"time"
)

// DB 存取值
type DB struct{}

var dbSingleton *DB
var dbOnce sync.Once

// DBIns 獲得單例對象
func DBIns() *DB {
	dbOnce.Do(func() {
		dbSingleton = &DB{}
	})
	return dbSingleton
}

// PingDBOnce ping db 測試
func (*DB) PingDBOnce() (apiErr errorcode.Error) {
	db, apiErr := model.NewConn(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	// defer db.Close()

	time.Sleep(10 * time.Second)
	admin := model.Admin{}

	if err := db.Where("account = ?", "user3").Find(&admin).Error; err != nil {
		fmt.Println("Ping Error:", err.Error())
		db.DB().Close()
		return
	}

	fmt.Println(admin)
	return
}

// PingDBSecond ping db 測試
func (*DB) PingDBSecond() (apiErr errorcode.Error) {
	db, apiErr := model.NewConn(global.GoFormatSl)
	if apiErr != nil {
		return
	}
	// defer db.Close()

	time.Sleep(10 * time.Second)
	admin := model.Admin{}

	if err := db.Where("id = ?", 1).Find(&admin).Error; err != nil {
		fmt.Println("Ping Error:", err.Error())
		db.DB().Close()
		return
	}

	fmt.Println(admin)

	return
}
