package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"bilibili/config"
)

var redisClient *redis.Pool
var RedisDb redis.Conn
func init() {
	dbConfig := config.GetDbConfig()
	var maxIdle int
	var maxActive int
	var MaxIdleTimeout time.Duration
	var timeout time.Duration
	if v, ok := dbConfig["MAX_IDLE"]; ok {
		maxIdle = v.(int)
	}
	if v, ok := dbConfig["MAX_ACTIVE"]; ok {
		maxActive = v.(int)
	}
	if v, ok := dbConfig["MAX_IDLE_TIMEOUT"]; ok {
		MaxIdleTimeout = time.Duration(v.(int))
	}
	if v, ok := dbConfig["TIMEOUT"]; ok {
		timeout = time.Duration(v.(int))
	}
	redisHost := dbConfig["REDIS_HOST"]
	redisPort := dbConfig["REDIS_PORT"]
	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: MaxIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp",fmt.Sprintf("%s:%d", redisHost, redisPort),
				redis.DialPassword(dbConfig["REDIS_PWD"].(string)),
				redis.DialDatabase(dbConfig["REDIS_SELECT"].(int)),
				redis.DialConnectTimeout(timeout*time.Second),
				redis.DialReadTimeout(timeout*time.Second),
				redis.DialWriteTimeout(timeout*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	RedisDb = redisClient.Get()

}

