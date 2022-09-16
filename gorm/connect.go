package gorm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup(dsn string) {
	DB = InitDB(dsn)
}

func InitDB(dsn string) *gorm.DB {
	var err error
	var db *gorm.DB
	if dsn == "" {
		panic("get mysql config error")
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect db error: %#v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	return db
}
