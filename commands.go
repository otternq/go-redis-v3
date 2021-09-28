package redis

import (
	"fmt"

	pkgredis "gopkg.in/redis.v3"
)

type observabilityCmd interface {
	Err() error
	fmt.Stringer
}

var _ Cmd = &pkgredis.Cmd{}

type Cmd interface {
	Err() error
	Result() (interface{}, error)
	String() string
}

type StatusCmd interface {
	Err() error
	Result() (string, error)
	String() string
	Val() string
}

type StringCmd interface {
	Bytes() ([]byte, error)
	Err() error
	Float64() (float64, error)
	Int64() (int64, error)
	Result() (string, error)
	Scan(val interface{}) error
	String() string
	Uint64() (uint64, error)
	Val() string
}

type StringSliceCmd interface {
	Err() error
	Result() ([]string, error)
	String() string
	Val() []string
}

type IntCmd interface {
	Val() int64
	String() string
	Err() error
	Result() (int64, error)
}

type BoolCmd interface {
	Err() error
	Result() (bool, error)
	String() string
	Val() bool
}
