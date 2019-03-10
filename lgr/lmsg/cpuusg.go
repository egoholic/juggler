package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type CPUUsgMessage struct {
	msgBase
	percentage int
}

func CPUUsg(place, label string, percentage int) *CPUUsgMessage {
	base := msgBase{typ: CPUUSG, ruid: uuid.New(), place: place, label: label}
	return &CPUUsgMessage{base, percentage}
}

func (*CPUUsgMessage) LogType() LogType {
	return CPUUSG
}

func (msg *CPUUsgMessage) JSON() string {
	return ""
}
