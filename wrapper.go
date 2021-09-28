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

var (
	_ Client     = &Wrapper{}
	_ EvalClient = &Wrapper{}
	_ HashClient = &Wrapper{}
)

type Wrapper struct {
	client       *pkgredis.Client
	instanceName string
}

func (w *Wrapper) Eval(ctx context.Context, script string, keys, args []string) (cmd Cmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.eval", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()

	cmd = w.client.Eval(script, keys, args)
	return
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

func (w *Wrapper) Incr(ctx context.Context, key string) (cmd IntCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.incr", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Incr(key)
	return

}

func (w *Wrapper) Decr(ctx context.Context, key string) (cmd IntCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.decr", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Decr(key)
	return
}

func (w *Wrapper) HGet(ctx context.Context, key, field string) (cmd StringCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.hget", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HGet(key, field)
	return
}
func (w *Wrapper) HSet(ctx context.Context, key, field, value string) (cmd BoolCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.hset", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HSet(key, field, value)
	return
}
func (w *Wrapper) HKeys(ctx context.Context, key string) (cmd StringSliceCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.hkeys", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.HKeys(key)
	return
}

func (w *Wrapper) Ping(ctx context.Context) (cmd StatusCmd) {
	var recordCallFunc = recordCall(ctx, "go.redis.ping", w.instanceName)
	defer func() {
		recordCallFunc(cmd)
	}()
	cmd = w.client.Ping()
	return cmd
}
