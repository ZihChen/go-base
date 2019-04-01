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

// DBPool 存放 db 連線池的全域變數
var DBPool *gorm.DB

// NewConn 取connection pool 新連線
func NewConn(mode string) (*gorm.DB, errorcode.Error) {
	fmt.Println("尚未連線====>", DBPool)
	if DBPool != nil {
		return DBPool.New(), nil
	}

	var apiErr errorcode.Error
	DBPool, apiErr = DBConnection(mode)
	if apiErr != nil {
		return nil, apiErr
	}

	fmt.Println("連線後====>", DBPool)
	return DBPool.New(), nil
}

// DBConnection 建立 Pool 連線
func DBConnection(mode string) (*gorm.DB, errorcode.Error) {
	var err error
	connString := composeString(mode)
	DBPool, err = gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("當DB連線錯誤：", err.Error())
		go helper.FatalLog(fmt.Sprintf("DB_CONNECT_ERROR: %v", err.Error()))
		apiErr := errorcode.GetAPIError("DB_CONNECT_ERROR")
		return nil, apiErr
	}

	// 限制最大開啟的連線數
	DBPool.DB().SetMaxIdleConns(100)
	// 限制最大閒置連線數
	DBPool.DB().SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	DBPool.DB().SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// DBPool.SingularTable(true)
	// 開啟SQL Debug模式
	DBPool.LogMode(global.Config.DB.Debug)

	return DBPool, nil
}

// DBConnectTest 檢查DB機器是否可以連線
// func DBConnectTest() {
// 	// 檢查 Master 連線
// 	_, apiErr := DBConnection(global.GoFormatMa)
// 	if apiErr != nil {

// 		panic("DB MASTER CONNECT ERROR")
// 	}
// 	// defer dbM.Close()

// 	// 檢查M Slave 連線
// 	// dbS, apiErr := DBConnection(global.GoFormatSl)
// 	// if apiErr != nil {
// 	// 	panic("DB SLAVE CONNECT ERROR")
// 	// }
// 	// defer dbS.Close()
// }

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
