package mredis

import goredislib "github.com/go-redis/redis/v9"

type UniversalConfig = goredislib.UniversalOptions

type RedisConnectionConfig struct {
	Addr             string
	UserName         string
	Password         string
	MasterName       string
	SentinelAddrs    []string
	SentinelPassword string
}
