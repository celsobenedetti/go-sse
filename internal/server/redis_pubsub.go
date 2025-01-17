package server

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisPubSub(addr string) *RedisPubSub {
	return &RedisPubSub{
		rdb: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

type RedisPubSub struct {
	rdb *redis.Client
}

func (r *RedisPubSub) Health() (string, error) {
	ctx := context.Background()
	return r.rdb.Ping(ctx).Result()
}

func (r *RedisPubSub) Publish(msg Message) error {
	ctx := context.Background()

	cmd := r.rdb.Publish(ctx, fmt.Sprintf("rooms:%s", msg.RoomId), msg)

	return cmd.Err()
}

func (r *RedisPubSub) Subscribe(roomId string) *redis.PubSub {
	ctx := context.Background()
	return r.rdb.Subscribe(ctx, fmt.Sprintf("rooms:%s", roomId))
}
