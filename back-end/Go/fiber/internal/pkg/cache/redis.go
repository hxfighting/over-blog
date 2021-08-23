package cache

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/tools/log"
)

var (
	client *redis.Client
	once   sync.Once
)

type LockCtx func() (context.Context, *sync.Mutex)

func New() (err error) {
	cfg := configs.Config.Redis
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:         cfg.Host + ":" + cfg.Port,
			Password:     cfg.Password, // no password set
			DB:           cfg.DB,       // use default DB
			DialTimeout:  cfg.DialTimeout,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		})
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err = client.Ping(ctx).Result()
		if err != nil {
			return
		}
	})
	return
}

func Close() {
	if err := client.Close(); err != nil {
		log.Log.Err(err).Send()
	}
}

func Remember(f LockCtx, key string, ttl time.Duration, callback func() (interface{}, error)) (interface{}, error) {
	ctx, lock := f()
	exist := client.Get(ctx, key)
	if exist.Err() != nil {
		lock.Lock()
		defer lock.Unlock()
		exist = client.Get(ctx, key)
		if r := exist.Val(); r != "" {
			return r, nil
		}
		val, err := callback()
		if val == nil || err != nil {
			client.Set(ctx, key, "", time.Minute)
			return nil, err
		}
		res := client.Set(ctx, key, val, ttl)
		if res.Err() != nil {
			return nil, res.Err()
		}
		return val, nil
	}
	return exist.Val(), nil
}

func Remove(f LockCtx, key string) error {
	ctx, _ := f()
	return client.Del(ctx, key).Err()
}
