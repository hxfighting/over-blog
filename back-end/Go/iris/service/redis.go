package service

import (
	"blog/config"
	"github.com/go-redis/redis"
)

var (
	Redis = NewInstance()
)

func NewInstance() *redis.Client {
	addr := config.GetConfig("redis.host").(string)
	port := config.GetConfig("redis.port").(string)
	pass := config.GetConfig("redis.password").(string)
	db := int(config.GetConfig("redis.db").(int64))
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
