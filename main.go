package main

import "github.com/zhtfirst/go-packages/example"

func main() {
	//// 初始化日志
	//example.Logging()

	//// JWT
	//example.Jwt()

	//// 获取配置
	//example.GetConfig()

	//// Mysql 数据库连接
	//example.DbConnect()

	//// Redis 数据库连接
	//example.RedisConnect()

	//// snowflake 分布式ID生成器
	//example.Snowflake()

	//// MongoDB 数据库连接
	//example.MongoDbConnect()

	// handler 助手方法
	example.HandlerTest()

}
