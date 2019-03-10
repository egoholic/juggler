package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type PanicMessage struct {
	msgBase
	err error
}

func Panic(place, label string, err error) *PanicMessage {
	base := msgBase{typ: PANIC, ruid: uuid.New(), place: place, label: label}
	return &PanicMessage{base, err}
}

func (*PanicMessage) LogType() LogType {
	return PANIC
}

func (msg *PanicMessage) JSON() string {
	return ""
}
