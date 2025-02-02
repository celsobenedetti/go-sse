package server

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisKV struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewRedisKV(client *redis.Client) *RedisKV {
	return &RedisKV{
		rdb: client,
		ttl: 10 * time.Minute, // TODO: C-32 get from env
	}
}

// WIP: check the "any" behavior
func (r *RedisKV) Set(k string, v any) error {
	ctx := context.TODO()

	slog.Info("kv saved to redis", "k", k, "v", v)
	return r.rdb.Set(ctx, k, v, r.ttl).Err()
}

func (r *RedisKV) Get(k string) (any, error) {
	ctx := context.TODO()

	cmd := r.rdb.Get(ctx, k)

	return cmd.String(), cmd.Err()
}
