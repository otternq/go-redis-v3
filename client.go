// Package redis instruments redis interactions with Open Census
package redis

import (
	"context"
	"time"
)

type Client interface {
	Get(ctx context.Context, key string) StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusCmd
}

type HashClient interface {
	HGet(ctx context.Context, key, field string) StringCmd
	HSet(ctx context.Context, key, field, value string) BoolCmd
	HKeys(ctx context.Context, key string) StringSliceCmd
}
