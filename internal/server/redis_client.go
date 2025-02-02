package server

import "github.com/redis/go-redis/v9"

type RedisClient struct {
	redis  *redis.Client
	Pubsub *RedisPubSub
	KV     *RedisKV
}

func NewRedisClient(addr string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisClient{
		Pubsub: NewRedisPubSub(rdb),
		KV:     NewRedisKV(rdb),
	}
}
