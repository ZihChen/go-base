package business

import (
	"GoFormat/app/global/errorcode"
	"GoFormat/app/repository"
	"sync"
)

// DBBus 管理者Business
type DBBus struct {
}

var dbSingleton *DBBus
var dbOnce sync.Once

// DBIns 獲得單例對象
func DBIns() *DBBus {
	dbOnce.Do(func() {
		dbSingleton = &DBBus{}
	})
	return dbSingleton
}

// PingDBOnce 存值進入redis
func (a *DBBus) PingDBOnce() (err errorcode.Error) {
	db := repository.DBIns()
	err = db.PingDBOnce()
	if err != nil {
		return
	}

	return
}

// PingDBSecond 存值進入redis
func (a *DBBus) PingDBSecond() (err errorcode.Error) {
	db := repository.DBIns()
	err = db.PingDBSecond()
	if err != nil {
		return
	}

	return
}
