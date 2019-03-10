package lmsg

import (
	"time"

	"github.com/egoholic/juggler/lgr/lmsg/uuid"
)

type PerfMessage struct {
	msgBase
	durations *map[string]time.Duration
}

func Perf(place, label string, durations *map[string]time.Duration) *PerfMessage {
	base := msgBase{typ: PERF, ruid: uuid.New(), place: place, label: label}
	return &PerfMessage{base, durations}
}

func (*PerfMessage) LogType() LogType {
	return PERF
}

func (msg *PerfMessage) JSON() string {
	return ""
}
