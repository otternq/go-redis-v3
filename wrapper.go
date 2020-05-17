package redis

import (
	"context"
	"time"

	pkgredis "gopkg.in/redis.v3"
)

func Wrap(client *pkgredis.Client, instanceName string) *Wrapper {
	return &Wrapper{
		client:       client,
		instanceName: instanceName,
	}
}

var _ Client = &Wrapper{}

type Wrapper struct {
	client       *pkgredis.Client
	instanceName string
}

func (w *Wrapper) Get(ctx context.Context, key string) (cmd StringCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.get", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Get(key)
	return
}

func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (cmd StatusCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.set", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Set(key, value, expiration)
	return
}
