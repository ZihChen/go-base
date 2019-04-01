package model

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// dbCon DB連線資料
type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// MasterPool 存放 db Master 連線池的全域變數
var MasterPool *gorm.DB

// SlavePool 存放 db Slave 連線池的全域變數
var SlavePool *gorm.DB

// MasterConnect 建立 Master Pool 連線
func MasterConnect() (*gorm.DB, errorcode.Error) {
	if MasterPool != nil {
		return MasterPool, nil
	}

	connString := composeString(global.GoFormatMa)
	MasterPool, err := gorm.Open("mysql", connString)
	if err != nil {
		go helper.FatalLog(fmt.Sprintf("DB_CONNECT_ERROR: %v", err.Error()))
		apiErr := errorcode.GetAPIError("DB_CONNECT_ERROR")
		return nil, apiErr
	}

	// 限制最大開啟的連線數
	MasterPool.DB().SetMaxIdleConns(100)
	// 限制最大閒置連線數
	MasterPool.DB().SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	MasterPool.DB().SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// MasterPool.SingularTable(true)
	// 開啟SQL Debug模式
	MasterPool.LogMode(global.Config.DB.Debug)

	return MasterPool, nil
}

// SlaveConnect 建立 Slave Pool 連線
func SlaveConnect() (*gorm.DB, errorcode.Error) {
	if SlavePool != nil {
		return SlavePool, nil
	}

	connString := composeString(global.GoFormatSl)
	SlavePool, err := gorm.Open("mysql", connString)
	if err != nil {
		go helper.FatalLog(fmt.Sprintf("DB_CONNECT_ERROR: %v", err.Error()))
		apiErr := errorcode.GetAPIError("DB_CONNECT_ERROR")
		return nil, apiErr
	}

	// 限制最大開啟的連線數
	SlavePool.DB().SetMaxIdleConns(100)
	// 限制最大閒置連線數
	SlavePool.DB().SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	SlavePool.DB().SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// SlavePool.SingularTable(true)
	// 開啟SQL Debug模式
	SlavePool.LogMode(global.Config.DB.Debug)

	return SlavePool, nil
}

// DBPing 檢查DB是否啟動
func DBPing() {
	// 檢查 master db
	masterPool, apiErr := MasterConnect()
	if apiErr != nil {
		panic("MASTER DB CONNECT ERROR")
	}

	err := masterPool.DB().Ping()
	if err != nil {
		panic("PING MASTER DB ERROR:" + err.Error())
	}

	// 檢查 slave db
	slavePool, apiErr := SlaveConnect()
	if apiErr != nil {
		panic("SLAVE DB CONNECT ERROR")
	}

	err = slavePool.DB().Ping()
	if err != nil {
		panic("PING SLAVE DB ERROR:" + err.Error())
	}
}

// composeString 組合DB連線前的字串資料
func composeString(mode string) string {
	db := dbCon{}

	switch mode {
	case global.GoFormatMa:
		db.Host = global.Config.DBMaster.Host
		db.Username = global.Config.DBMaster.Username
		db.Password = global.Config.DBMaster.Password
		db.Database = global.Config.DBMaster.Database
	case global.GoFormatSl:
		db.Host = global.Config.DbSlave.Host
		db.Username = global.Config.DbSlave.Username
		db.Password = global.Config.DbSlave.Password
		db.Database = global.Config.DbSlave.Database
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Database)
}
