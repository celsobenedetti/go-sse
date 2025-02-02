package server

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisPubSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{
		rdb: client,
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
