package redis

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"gopkg.in/redis.v3"
)

const (
	statusError = "ERROR"
	statusOK    = "OK"
)

// The following tags are aooplied to stats recorded by this package
var (
	// GoRedisName is the name of the redis instance.
	GoRedisInstanceName, _ = tag.NewKey("go_redis_instance_name")

	// GoRedisMethod is the client method called.
	GoRedisMethod, _ = tag.NewKey("go_redis_method")

	// GoRedisStatus identifies the command status
	GoRedisStatus, _ = tag.NewKey("go_redis_status")

	DefaultTags = []tag.Key{
		GoRedisMethod,
		GoRedisStatus,
	}
)

// The following measures are supported for use in custom views.
var (
	MeasureLatencyMs     = stats.Int64("go.redis/latency", "The latency of calls in milliseconds", stats.UnitMilliseconds)
	MeasureResponseBytes = stats.Int64("go.redis/received_bytes", "The number of bytes returned from a command", stats.UnitBytes)
)

// Default distributions used by views in this package
var (
	DefaultSizeDistribution = view.Distribution(
		1024,
		2048,
		4096,
		16384,
		65536,
		262144,
		1048576,
		4194304,
		16777216,
		67108864,
		268435456,
		1073741824,
		4294967296,
	)
	DefaultMillisecondsDistribution = view.Distribution(
		0.0,
		0.001,
		0.005,
		0.01,
		0.05,
		0.1,
		0.5,
		1.0,
		1.5,
		2.0,
		2.5,
		5.0,
		10.0,
		25.0,
		50.0,
		100.0,
		200.0,
		400.0,
		600.0,
		800.0,
		1000.0,
		1500.0,
		2000.0,
		2500.0,
		5000.0,
		10000.0,
		20000.0,
		40000.0,
		100000.0,
		200000.0,
		500000.0,
	)
)

// Package cache provides some convenience views.
// You still need to register these views for data to actually be collected.
// You can use the RegisterAllViews function for this.
var (
	GoRedisLatencyView = &view.View{
		Name:        "go.redis/client/latency",
		Description: "The distribution of latency of various calls in milliseconds",
		Measure:     MeasureLatencyMs,
		Aggregation: DefaultMillisecondsDistribution,
		TagKeys:     DefaultTags,
	}

	GoRedisCallsView = &view.View{
		Name:        "go.redis/client/calls",
		Description: "The number of various calls of methods",
		Measure:     MeasureLatencyMs,
		Aggregation: view.Count(),
		TagKeys:     DefaultTags,
	}

	GoRedisBytesView = &view.View{
		Name:        "go.redis/client/received_bytes",
		Description: "Total bytes received from Redis",
		Measure:     MeasureResponseBytes,
		Aggregation: DefaultSizeDistribution,
		TagKeys:     DefaultTags,
	}

	DefaultViews = []*view.View{GoRedisLatencyView, GoRedisCallsView, GoRedisBytesView}
)

// RegisterAllViews registers all the cache views to enable collection of stats
func RegisterAllViews() error {
	return view.Register(DefaultViews...)
}

func recordCall(ctx context.Context, method string, instanceName string) func(cmd observabilityCmd) {
	var startTime = time.Now()

	return func(cmd observabilityCmd) {
		var (
			timeSpentMs = time.Since(startTime).Milliseconds()
			tags        = []tag.Mutator{
				tag.Insert(GoRedisInstanceName, instanceName),
				tag.Insert(GoRedisMethod, method),
			}
		)

		if cmd.Err() != nil && cmd.Err() != redis.Nil {
			tags = append(tags, tag.Insert(GoRedisStatus, statusError))
		} else {
			tags = append(tags, tag.Insert(GoRedisStatus, statusOK))
		}

		_ = stats.RecordWithTags(ctx, tags, MeasureLatencyMs.M(timeSpentMs))
		_ = stats.RecordWithTags(ctx, tags, MeasureResponseBytes.M(int64(len([]byte(cmd.String())))))
	}
}
