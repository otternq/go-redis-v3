package redis

import "fmt"

type Cmd interface {
	Err() error
	fmt.Stringer
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
