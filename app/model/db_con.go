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

// DBConnection 建立Master連線
func DBConnection(mode string) (*gorm.DB, errorcode.Error) {
	var err error
	fmt.Println("尚未連線憶體位址：", DBPool)
	if DBPool != nil {
		fmt.Println("當DBPool不是nil,記憶體位址：", DBPool)
		fmt.Println("當DBPool不是nil,Stats狀態：", DBPool.DB().Stats())
		return DBPool, nil
	}

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

	fmt.Println("第一次連線成功,Stats狀態：", DBPool.DB().Stats())
	fmt.Println("第一次連線成功記憶體位址：", DBPool)
	return DBPool, nil
}

// func DBConnection(mode string) (db *gorm.DB, apiErr errorcode.Error) {
// 	connString := composeString(mode)
// 	db, err := gorm.Open("mysql", connString)
// 	if err != nil {
// 		go helper.FatalLog(fmt.Sprintf("DB_CONNECT_ERROR: %v", err))
// 		apiErr = errorcode.GetAPIError("DB_CONNECT_ERROR")
// 		return
// 	}

// 	// 限制最大開啟的連線數
// 	db.DB().SetMaxIdleConns(100)
// 	// 限制最大閒置連線數
// 	db.DB().SetMaxOpenConns(2000)
// 	// 空閒連線 timeout 時間
// 	db.DB().SetConnMaxLifetime(300 * time.Second)

// 	log.Println(db.DB().Stats())

// 	// 全局禁用表名复数
// 	// db.SingularTable(true)
// 	// 開啟SQL Debug模式
// 	db.LogMode(global.Config.DB.Debug)

// 	return
// }

// DBConnectTest 檢查DB機器是否可以連線
func DBConnectTest() {
	// 檢查 Master 連線
	_, apiErr := DBConnection(global.GoFormatMa)
	if apiErr != nil {

		panic("DB MASTER CONNECT ERROR")
	}
	// defer dbM.Close()

	// 檢查M Slave 連線
	// dbS, apiErr := DBConnection(global.GoFormatSl)
	// if apiErr != nil {
	// 	panic("DB SLAVE CONNECT ERROR")
	// }
	// defer dbS.Close()

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
