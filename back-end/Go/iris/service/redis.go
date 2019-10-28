package service

import (
	"blog/config"
	"github.com/go-redis/redis"
	"log"
)

var (
	Redis *redis.Client
)

func NewRedis() {
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
		log.Fatalln(err.Error())
	}
	Redis = client
}
