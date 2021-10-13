package util

import (
	"context"
	"encoding/json"
	"time"

	"github.com/WolffunGame/theta-shared-database/database/mredis"
	goredislib "github.com/go-redis/redis/v8"
)

// Set struct to redis
// EX used: err = util.SetObject(ctx, "key", userModel, 86400)
func SetObject(ctx context.Context, key string, value interface{}, expirationSecond int) error {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return mredis.GetClient().Set(ctx, key, jsonStr, time.Duration(expirationSecond)*time.Second).Err()
}

// Get struct from redis
// EX used: existsFlag, err := util.GetObject(ctx, "key", &userModel)
func GetObject(ctx context.Context, key string, refObj interface{}) (bool, error) {
	bytes, err := mredis.GetClient().Get(ctx, key).Bytes()
	if err != nil {
		return false, err
	}
	if err == goredislib.Nil {
		// Key not exists
		return false, nil
	}
	err = json.Unmarshal(bytes, &refObj)
	if err != nil {
		return false, err
	}
	return true, nil
}
