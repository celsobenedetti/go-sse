package containers

import (
	"context"
	"log/slog"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

func Redis() (*redis.RedisContainer, func(), error) {
	ctx := context.Background()
	redisContainer, err := redis.Run(ctx,
		"redis:7.4.2", // TODO: CE-32 get this from env
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
		// redis.WithConfigFile(filepath.Join("testdata", "redis7.confg")),
	)

	close := func() {
		if err := testcontainers.TerminateContainer(redisContainer); err != nil {
			slog.Error("failed to terminate container", "err", err.Error())
		}
	}
	if err != nil {
		slog.Error("failed to start container", "err", err.Error())
		return nil, nil, err
	}
	return redisContainer, close, nil
}
