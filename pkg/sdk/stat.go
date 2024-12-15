package sdk

import (
	"sync/atomic"
)

type stat struct {
	lastIdUpdate atomic.Int64 // time.Time
}

func newStat() *stat {
	return &stat{}
}
