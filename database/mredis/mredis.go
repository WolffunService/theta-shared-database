package mredis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/WolffunGame/theta-shared-database/database/mredis/thetanlock"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var client goredislib.UniversalClient

func NewPool() redis.Pool {
	return goredis.NewPool(client)
}

func ConnectRedisUniversal(opts *UniversalConfig) {
	client = goredislib.NewUniversalClient(opts)
	ping := client.Ping(context.Background())
	if err := ping.Err(); err != nil {
		panic(fmt.Sprintf("connect redis: %v", err))
	}

	thetanlock.InitPool(NewPool())
}

func ConnectRedisV2(config *RedisConnectionConfig) {
	ops := &goredislib.FailoverOptions{}
	if len(config.MasterName) > 0 {
		ops.MasterName = config.MasterName
	}
	if len(config.Password) > 0 {
		ops.Password = config.Password
	}
	if len(config.SentinelAddrs) > 0 {
		ops.SentinelAddrs = config.SentinelAddrs
	}
	if len(config.SentinelPassword) > 0 {
		ops.SentinelPassword = config.SentinelPassword
	}

	client = goredislib.NewFailoverClient(ops)
	client.Ping(context.Background())
	thetanlock.InitPool(NewPool())
}

func ConnectRedisV1(config *RedisConnectionConfig) {
	client = goredislib.NewClient(&goredislib.Options{
		Addr:     config.Addr,
		Username: config.UserName,
		Password: config.Password,
	})
	thetanlock.InitPool(NewPool())
}

func GetClient() goredislib.UniversalClient {
	return client
}

func Exists(ctx context.Context, key string) (bool, error) {
	res, err := client.Exists(ctx, key).Result()
	return res == 1, err
}

// Set struct to redis
// EX used: err = util.SetObject(ctx, "key", userModel, 86400)
func SetObject(ctx context.Context, key string, value interface{}, expirationSecond int) error {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, jsonStr, time.Duration(expirationSecond)*time.Second).Err()
}

// Get struct from redis
// EX used: existsFlag, err := util.GetObject(ctx, "key", &userModel)
func GetObject(ctx context.Context, key string, refObj interface{}) (bool, error) {
	bytes, err := client.Get(ctx, key).Bytes()

	if err != nil {
		if err == goredislib.Nil {
			// Key not exists
			return false, nil
		}

		return false, err
	}

	err = json.Unmarshal(bytes, &refObj)
	if err != nil {
		return false, err
	}
	return true, nil
}
