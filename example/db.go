package example

import (
	"fmt"

	"github.com/zhtfirst/go-packages/example/internal/dao"

	"github.com/zhtfirst/go-packages/config"
	"github.com/zhtfirst/go-packages/gorm"
)

func DbConnect() {
	config.Setup("") // 初始化配置

	// 连接数据库
	gorm.Setup(config.GetString("mysql"))
	fmt.Println("连接数据库:", gorm.DB)
	dao.DB = gorm.DB
	if config.GetBoole("log_debug") {
		dao.DB = gorm.DB.Debug() // 开启debug模式
	}
	// sql Demo: db = GetDB(ctx).WithContext(ctx).Model(&sqlmodel.OfferDetail{}).Where(m, conds...)

	//// 读写分离连接数据库
	//gorm.ReadWriteSetup(config.GetString("mysql_read"), config.GetString("mysql"))
	//fmt.Println("读写分离连接数据库:", gorm.RWDB) // *gorm.DB
	//dao.DB = gorm.RWDB
	//if config.GetBoole("log_debug") {
	//	dao.DB = gorm.RWDB.Debug() // 开启debug模式
	//}
	// sql Demo: db = GetDB(ctx).Clauses(dbresolver.Read).WithContext(ctx).Model(&sqlmodel.OfferDetail{}).Where(m, conds...)

}
