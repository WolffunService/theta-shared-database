package mredis

import goredislib "github.com/redis/go-redis/v9"

type UniversalConfig = goredislib.UniversalOptions

type RedisConnectionConfig struct {
	Addr             string
	UserName         string
	Password         string
	MasterName       string
	SentinelAddrs    []string
	SentinelPassword string
}
