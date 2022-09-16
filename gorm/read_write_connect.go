package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"time"

	"gorm.io/plugin/dbresolver"
)

var RWDB *gorm.DB

func ReadWriteSetup(connRead, connWrite string) {
	RWDB = InitRWDB(connRead, connWrite)
}

func InitRWDB(connRead, connWrite string) *gorm.DB {
	var (
		err error
		db  *gorm.DB
	)

	if connRead == "" || connWrite == "" {
		panic("get mysql config error")
	}
	db, err = gorm.Open(mysql.Open(connWrite), &gorm.Config{}) // `connWrite` 作为 sources（DB 的默认连接）
	if err != nil {
		log.Fatalf("connect db error: %#v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(500)                 // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 连接最大时间

	_ = db.Use(dbresolver.
		Register(dbresolver.Config{
			// `connWrite` 作为 sources; `connRead`、`connWrite` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                       // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connWrite)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                     // sources/replicas 负载均衡策略
		}))
	return db
}
