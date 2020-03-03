package redisInterface

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/go-redis/redis"
	"time"
)

var AFIRedisModuleName string

type AFIRedisModule interface {
	ark.AFIModule
	Connect(addr []string, authKey string, poolSize int) error
	GetConn() redis.Cmdable
	// some basic command
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	INCR(key string) (int64, error)
	INCRBy(key string, value int64) (int64, error)
	HSet(key, field string, value interface{}, expiration time.Duration) error
	HMSet(key string, fields map[string]interface{}, expiration time.Duration) error
	HGet(key, field string) (string, error)
	HGetAll(key string) (map[string]string, error)
	Del(keys ...string)
}
