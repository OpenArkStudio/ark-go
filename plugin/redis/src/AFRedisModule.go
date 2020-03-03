package redisSrc

import (
	"errors"
	"fmt"
	logInterface "github.com/ArkNX/ark-go/plugin/log/interface"
	"log"
	"time"

	"github.com/go-redis/redis"

	"github.com/ArkNX/ark-go/interface"
	redisInterface "github.com/ArkNX/ark-go/plugin/redis/interface"
	"github.com/ArkNX/ark-go/util"
)

var (
	redisModuleType   = util.GetType((*AFCRedisModule)(nil))
	redisModuleName   = util.GetName((*AFCRedisModule)(nil))
	redisModuleUpdate = fmt.Sprintf("%p", (&AFCRedisModule{}).Update) != fmt.Sprintf("%p", (&ark.AFCModule{}).Update)
)

var (
	// ErrInvalidRedisAddr describes error of invalid redis address
	ErrInvalidRedisAddr = errors.New("invalid redis address")
)

func init() {
	redisInterface.AFIRedisModuleName = util.GetName((*AFCRedisModule)(nil))
}

type AFCRedisModule struct {
	ark.AFCModule
	// other module
	log logInterface.AFILogModule
	// other data
	conn redis.Cmdable
}

func (redisModule *AFCRedisModule) Init() error {
	m := redisModule.GetPluginManager().FindModule(logInterface.AFILogModuleName)
	logModule, ok := m.(logInterface.AFILogModule)
	if !ok {
		log.Fatal("failed to get log module in httpServer module")
	}
	redisModule.log = logModule
	return nil
}

func (redisModule *AFCRedisModule) Connect(addrs []string, password string, poolSize int) error {
	if len(addrs) == 0 {
		return ErrInvalidRedisAddr
	}

	var conn redis.Cmdable
	if len(addrs) > 1 {
		conn = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    addrs,
			Password: password,
			PoolSize: poolSize,
		})
	} else {
		conn = redis.NewClient(&redis.Options{
			Addr:     addrs[0],
			Password: password,
			PoolSize: poolSize,
		})
	}

	if _, err := conn.Ping().Result(); err != nil {
		redisModule.log.GetLogger().WithFields(map[string]interface{}{
			"redisAddr": addrs,
		}).Error("failed to ping redis during connection")
		return err
	}

	redisModule.conn = conn

	return nil

}

func (redisModule *AFCRedisModule) GetConn() redis.Cmdable {
	return redisModule.conn
}

// --------------- some basic cmd ---------------
func (redisModule *AFCRedisModule) Get(key string) (string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return "", err
	}

	return redisModule.conn.Get(key).Result()
}

func (redisModule *AFCRedisModule) Set(key string, value interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	return redisModule.conn.Set(key, value, expiration).Err()
}

func (redisModule *AFCRedisModule) INCR(key string) (int64, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return 0, err
	}

	return redisModule.conn.Incr(key).Result()
}

func (redisModule *AFCRedisModule) INCRBy(key string, value int64) (int64, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return 0, err
	}

	return redisModule.conn.IncrBy(key, value).Result()
}

func (redisModule *AFCRedisModule) HSet(key, field string, value interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	if err := redisModule.conn.HSet(key, field, value).Err(); err != nil {
		return err
	}

	if err := redisModule.conn.Expire(key, expiration).Err(); err != nil {
		redisModule.conn.Del(key)
		return err
	}

	return nil
}

func (redisModule *AFCRedisModule) HMSet(key string, fields map[string]interface{}, expiration time.Duration) error {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return err
	}

	if err := redisModule.conn.HMSet(key, fields).Err(); err != nil {
		return err
	}

	if err := redisModule.conn.Expire(key, expiration).Err(); err != nil {
		redisModule.conn.Del(key)
		return err
	}

	return nil
}

func (redisModule *AFCRedisModule) HGet(key, field string) (string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return "", err
	}

	return redisModule.conn.HGet(key, field).Result()
}

func (redisModule *AFCRedisModule) HGetAll(key string) (map[string]string, error) {
	if err := redisModule.conn.Ping().Err(); err != nil {
		return nil, err
	}

	return redisModule.conn.HGetAll(key).Result()
}

func (redisModule *AFCRedisModule) Del(keys ...string) {
	redisModule.conn.Del(keys...)
}
