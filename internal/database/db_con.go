package database

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/app/model"
	"goformat/library/errorcode"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// dbCon DB連線資料
type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// masterPool 存放 db Master 連線池的全域變數
var masterPool *gorm.DB

// slavePool 存放 db Slave 連線池的全域變數
var slavePool *gorm.DB

type IDber interface {
	MasterConnect() (*gorm.DB, errorcode.Error)
	SlaveConnect() (*gorm.DB, errorcode.Error)
	DBPing()
	CheckTableIsExist()
}

func NewDbConnection() IDber {
	return &dbCon{}
}

// MasterConnect 建立 Master Pool 連線
func (d *dbCon) MasterConnect() (*gorm.DB, errorcode.Error) {
	var err error

	if masterPool != nil {
		if global.Config.DB.Debug {
			return masterPool.Debug(), nil
		}
		return masterPool, nil
	}

	connString := d.composeString(global.DBMaster)
	masterPool, err = gorm.Open(mysql.Open(connString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		apiErr := helper.ErrorHandle(global.FatalLog, "DB_CONNECT_ERROR", err.Error())

		return nil, apiErr
	}

	sqlPool, err := masterPool.DB()

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// masterPool.SingularTable(true)

	if global.Config.DB.Debug {
		return masterPool.Debug(), nil
	}
	return masterPool, nil
}

// SlaveConnect 建立 Slave Pool 連線
func (d *dbCon) SlaveConnect() (*gorm.DB, errorcode.Error) {
	var err error

	if slavePool != nil {
		if global.Config.DB.Debug {
			return slavePool.Debug(), nil
		}
		return slavePool, nil
	}

	connString := d.composeString(global.DBSlaver)
	slavePool, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		apiErr := helper.ErrorHandle(global.FatalLog, "DB_CONNECT_ERROR", err.Error())
		return nil, apiErr
	}

	sqlPool, err := slavePool.DB()

	// 限制最大開啟的連線數
	sqlPool.SetMaxIdleConns(100)
	// 限制最大閒置連線數
	sqlPool.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlPool.SetConnMaxLifetime(15 * time.Second)

	// 全局禁用表名复数
	// slavePool.SingularTable(true)

	if global.Config.DB.Debug {
		return slavePool.Debug(), nil
	}
	return slavePool, nil
}

// DBPing 檢查DB是否啟動
func (d *dbCon) DBPing() {
	// 檢查 master db
	masterPool, apiErr := d.MasterConnect()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	masterDB, err := masterPool.DB()
	if err != nil {
		log.Fatalf("🔔🔔🔔 CONNECT MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}
	err = masterDB.Ping()
	if err != nil {
		log.Fatalf("🔔🔔🔔 PING MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

	// 檢查 slave db
	slavePool, apiErr := d.SlaveConnect()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 SLAVE DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBSlave.Host)
	}
	slaveDB, err := slavePool.DB()
	if err != nil {
		log.Fatalf("🔔🔔🔔 CONNECT SLAVE DB ERROR: %v 🔔🔔🔔", err.Error())
	}
	err = slaveDB.Ping()
	if err != nil {
		log.Fatalf("🔔🔔🔔 PING SLAVE DB ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// CheckTableIsExist 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func (d *dbCon) CheckTableIsExist() {
	db, apiErr := d.MasterConnect()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	// 會自動建置 DB Table
	err := db.Set("gorm:table_options", "comment '細單規則'").AutoMigrate(&model.Admin{})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		_ = helper.ErrorHandle(global.FatalLog, "DB_TABLE_NOT_EXIST", fmt.Sprintf("❌ 設置DB錯誤： %v ❌", err.Error()))
		log.Fatalf("🔔🔔🔔 MIGRATE MASTER TABLE ERROR: %v 🔔🔔🔔", err.Error())
	}

}

// composeString 組合DB連線前的字串資料
func (d *dbCon) composeString(mode string) string {
	db := dbCon{}

	switch mode {
	case global.DBMaster:
		db.Host = global.Config.DBMaster.Host
		db.Username = global.Config.DBMaster.Username
		db.Password = global.Config.DBMaster.Password
		db.Database = global.Config.DBMaster.Database
	case global.DBSlaver:
		db.Host = global.Config.DBSlave.Host
		db.Username = global.Config.DBSlave.Username
		db.Password = global.Config.DBSlave.Password
		db.Database = global.Config.DBSlave.Database
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Database)
}
