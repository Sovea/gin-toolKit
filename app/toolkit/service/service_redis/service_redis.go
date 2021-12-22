package service_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedistPool *redis.Pool

func init() {
	RedistPool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 0,
		Dial: func() (redis.Conn, error) {
			c, error := redis.Dial("tcp", "localhost:6379")
			c.Do("AUTH", "")
			return c, error
		},
	}
}

func SetString(key string, value string, timeout int) bool {
	conn := RedistPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set error:", err)
		return false
	}
	_, err = conn.Do("expire", key, timeout)
	if err != nil {
		fmt.Println("set expire error: ", err)
		return false
	}
	return true
}

func GetString(key string) (string, error) {
	conn := RedistPool.Get()
	defer conn.Close()
	code, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get error:", err)
		return "", err
	}
	return code, nil
}

func DelString(key string) error {
	conn := RedistPool.Get()
	defer conn.Close()
	_, err := conn.Do("Del", key)
	if err != nil {
		fmt.Println("redis del error:", err)
		return err
	}
	return nil
}
