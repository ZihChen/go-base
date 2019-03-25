package model

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"fmt"

	"github.com/jinzhu/gorm"
)

// dbCon DB連線資料
type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// DBConnection 建立Master連線
func DBConnection(mode string) (db *gorm.DB, apiErr errorcode.Error) {
	connString := composeString(mode)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		go helper.FatalLog(fmt.Sprintf("DB_CONNECT_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("DB_CONNECT_ERROR")
		return
	}

	// 全局禁用表名复数
	// db.SingularTable(true)
	// 開啟SQL Debug模式
	db.LogMode(global.Config.DB.Debug)

	return
}

// DBConnectTest 檢查DB機器是否可以連線
func DBConnectTest() {
	// 檢查 Master 連線
	dbM, apiErr := DBConnection(global.GoFormatMa)
	if apiErr != nil {
		panic("DB MASTER CONNECT ERROR")
	}
	defer dbM.Close()

	// 檢查M Slave 連線
	dbS, apiErr := DBConnection(global.GoFormatSl)
	if apiErr != nil {
		panic("DB SLAVE CONNECT ERROR")
	}
	defer dbS.Close()

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
