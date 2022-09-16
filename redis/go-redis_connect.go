package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client *Redis

type Redis struct {
	prefix string
	Client redis.UniversalClient
}

func Setup(uri, username, password, prefix string, db int) {
	var err error
	Client, err = NewClient(uri, username, password, prefix, db)
	if err != nil {
		panic(err)
	}
}

func NewClient(uri, username, password, prefix string, db int) (client *Redis, err error) {
	addrs := strings.Split(uri, ",")
	opt := &redis.UniversalOptions{
		Addrs:    addrs,
		DB:       db,
		Username: username,
		Password: password,
	}
	c := redis.NewUniversalClient(opt)
	return &Redis{prefix: prefix, Client: c}, c.Ping(context.TODO()).Err()
}

func (i *Redis) BuildKey(key string) string {
	if len(i.prefix) > 0 {
		return i.prefix + ":" + key
	}
	return key
}

// GetLock 获取锁(redis)
func (i *Redis) GetLock(ctx context.Context, key string, expiration time.Duration) bool {
	result, err := i.Client.SetNX(ctx, i.BuildKey(key), 1, expiration).Result()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

//ReleaseLock 释放锁(redis)
func (i *Redis) ReleaseLock(ctx context.Context, key string) bool {
	return i.Client.Del(ctx, i.BuildKey(key)).Val() > 0
}

func (i *Redis) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	return i.Client.Do(ctx, args...)
}

func (i *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return i.Client.Set(ctx, i.BuildKey(key), value, expiration)
}

func (i *Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return i.Client.Get(ctx, i.BuildKey(key))
}

func (i *Redis) Del(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.Del(ctx, i.BuildKey(key))
}

func (i *Redis) Exists(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.Exists(ctx, i.BuildKey(key))
}

func (i *Redis) Incr(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.Incr(ctx, i.BuildKey(key))
}

func (i *Redis) Decr(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.Decr(ctx, i.BuildKey(key))
}

func (i *Redis) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return i.Client.IncrBy(ctx, i.BuildKey(key), value)
}

func (i *Redis) DecrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	return i.Client.DecrBy(ctx, i.BuildKey(key), value)
}

func (i *Redis) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	return i.Client.IncrByFloat(ctx, i.BuildKey(key), value)
}

func (i *Redis) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return i.Client.Expire(ctx, i.BuildKey(key), expiration)
}

func (i *Redis) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	return i.Client.ExpireAt(ctx, i.BuildKey(key), tm)
}

func (i *Redis) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return i.Client.Keys(ctx, i.BuildKey(pattern))
}

func (i *Redis) Pipeline() redis.Pipeliner {
	return i.Client.Pipeline()
}

func (i *Redis) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return i.Client.Pipelined(ctx, fn)
}

func (i *Redis) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return i.Client.TxPipelined(ctx, fn)
}

func (i *Redis) TxPipeline() redis.Pipeliner {
	return i.Client.TxPipeline()
}

func (i *Redis) Command(ctx context.Context) *redis.CommandsInfoCmd {
	return i.Client.Command(ctx)
}

func (i *Redis) ClientGetName(ctx context.Context) *redis.StringCmd {
	return i.Client.ClientGetName(ctx)
}

func (i *Redis) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	return i.Client.Echo(ctx, message)
}

func (i *Redis) Ping(ctx context.Context) *redis.StatusCmd {
	return i.Client.Ping(ctx)
}

func (i *Redis) Quit(ctx context.Context) *redis.StatusCmd {
	return i.Client.Quit(ctx)
}

func (i *Redis) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	return i.Client.SetBit(ctx, i.BuildKey(key), offset, value)
}

func (i *Redis) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	return i.Client.GetBit(ctx, i.BuildKey(key), offset)
}

func (i *Redis) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	return i.Client.BitCount(ctx, i.BuildKey(key), bitCount)
}

func (i *Redis) HSet(ctx context.Context, key, field string, value interface{}) *redis.IntCmd {
	return i.Client.HSet(ctx, i.BuildKey(key), field, value)
}

func (i *Redis) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	return i.Client.HGet(ctx, i.BuildKey(key), field)
}

func (i *Redis) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	return i.Client.HGetAll(ctx, i.BuildKey(key))
}

func (i *Redis) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	return i.Client.HExists(ctx, i.BuildKey(key), field)
}

func (i *Redis) HDel(ctx context.Context, key, field string) *redis.IntCmd {
	return i.Client.HDel(ctx, i.BuildKey(key), field)
}

func (i *Redis) HLen(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.HLen(ctx, i.BuildKey(key))
}

func (i *Redis) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	return i.Client.HKeys(ctx, i.BuildKey(key))
}

func (i *Redis) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	return i.Client.HVals(ctx, i.BuildKey(key))
}

func (i *Redis) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	newKeys := make([]string, len(keys))
	for d, key := range keys {
		newKeys[d] = i.BuildKey(key)
	}
	return i.Client.BLPop(ctx, timeout, newKeys...)
}

func (i *Redis) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	newKeys := make([]string, len(keys))
	for d, key := range keys {
		newKeys[d] = i.BuildKey(key)
	}
	return i.Client.BRPop(ctx, timeout, newKeys...)
}

func (i *Redis) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	return i.Client.BRPopLPush(ctx, i.BuildKey(source), i.BuildKey(destination), timeout)
}

func (i *Redis) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	return i.Client.LIndex(ctx, i.BuildKey(key), index)
}

func (i *Redis) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	return i.Client.LInsert(ctx, i.BuildKey(key), op, pivot, value)
}

func (i *Redis) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return i.Client.LInsertBefore(ctx, i.BuildKey(key), pivot, value)
}

func (i *Redis) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return i.Client.LInsertAfter(ctx, i.BuildKey(key), pivot, value)
}

func (i *Redis) LLen(ctx context.Context, key string) *redis.IntCmd {
	return i.Client.LLen(ctx, i.BuildKey(key))
}

func (i *Redis) LPop(ctx context.Context, key string) *redis.StringCmd {
	return i.Client.LPop(ctx, i.BuildKey(key))
}

func (i *Redis) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	return i.Client.LPopCount(ctx, i.BuildKey(key), count)
}

// TODO set 集合相关操作
// 添加
func (i *Redis) SAdd(ctx context.Context, key string, value ...interface{}) *redis.IntCmd {
	return i.Client.SAdd(ctx, i.BuildKey(key), value...)
}

// 获取容器key中所有元素
func (i *Redis) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	return i.Client.SMembers(ctx, i.BuildKey(key))
}

// 检查VALUE是否是SET容器中的成员
func (i *Redis) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	return i.Client.SIsMember(ctx, i.BuildKey(key), member)
}
