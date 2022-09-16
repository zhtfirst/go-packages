package example

import (
	"context"
	"fmt"
	"time"

	"github.com/zhtfirst/go-packages/config"
	"github.com/zhtfirst/go-packages/redis"
)

func RedisConnect() {
	config.Setup("") // 初始化配置

	// go-redis 连接redis
	redis.Setup(
		config.GetString("redis_conf", "hosts"),
		config.GetString("redis_conf", "username"),
		config.GetString("redis_conf", "password"),
		config.GetString("redis_conf", "prefix"),
		int(config.GetInt64("redis_conf", "db")),
	)

	// 使用redis:
	//fmt.Println("10秒内禁止重建：", redis.Client.GetLock(context.Background(), "one", 10*time.Second))
	//fmt.Println("释放锁：", redis.Client.ReleaseLock(context.Background(), "one"))

	result, err := redis.Client.Set(context.Background(), "key", "value", 100*time.Second).Result() // 设置key
	fmt.Printf("添加结果为：%v， err:%v", result, err)

}
