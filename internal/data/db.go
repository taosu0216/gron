package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gron/internal/conf"
	"time"
)

func NewDB(conf *conf.Data) *gorm.DB {
	addr := conf.Database.Addr
	user := conf.Database.User
	password := conf.Database.Password
	database := conf.Database.Database
	maxIdleConn := conf.Database.MaxIdleConn
	maxOpenConn := conf.Database.MaxOpenConn
	maxIdleTime := conf.Database.MaxIdleTime
	//slow_threshold_millisecond := conf.Database.SlowThresholdMillisecond
	dsn := user + ":" + password + "@tcp(" + addr + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(int(maxIdleConn))
	sqlDB.SetMaxOpenConns(int(maxOpenConn))
	sqlDB.SetConnMaxIdleTime(time.Duration(maxIdleTime))

	return db
}
