package lmsg

import (
	"time"
)

type PerfMessage struct {
	meta      *Meta
	durations *map[string]time.Duration
}

func Perf(meta *Meta, durations *map[string]time.Duration) *PerfMessage {
	return &PerfMessage{meta, durations}
}

func (*PerfMessage) LogType() LogType {
	return PERF
}

func (msg *PerfMessage) JSON() string {
	return ""
}
