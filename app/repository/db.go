package repository

import (
	"GoFormat/app/global/errorcode"
	"GoFormat/app/model"
	"fmt"
	"sync"
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
	db, apiErr := model.SlaveConnect()
	if apiErr != nil {
		return
	}

	admin := model.Admin{}

	go func() {

		if err := db.Where("account = ?", "user3").Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("account = ?", "user4").Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("account = ?", "user1").Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("account = ?", "user2").Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	fmt.Println(admin)
	return
}

// PingDBSecond ping db 測試
func (*DB) PingDBSecond() (apiErr errorcode.Error) {
	db, apiErr := model.MasterConnect()
	if apiErr != nil {
		return
	}

	admin := model.Admin{}

	go func() {

		if err := db.Where("id = ?", 1).Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("id = ?", 2).Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("id = ?", 3).Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	go func() {

		if err := db.Where("id = ?", 4).Find(&admin).Error; err != nil {
			fmt.Println("Find Error:", err.Error())
			return
		}
		fmt.Println("======>", db.DB().Stats())
	}()

	fmt.Println(admin)

	return
}
