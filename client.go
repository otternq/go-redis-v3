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
