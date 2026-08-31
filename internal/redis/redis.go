package redis

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func New(addr string , password string , dbIndex int) (*Redis , error){
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: dbIndex,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis: ping failed: %w", err)
	}

	slog.Info("Redis Connection Established successfully")

	return &Redis{client} , nil
}

func Close(r *Redis) error {
	slog.Info("Redis Connection Closed")

	return r.Client.Close()
}