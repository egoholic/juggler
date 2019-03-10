package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type DebugMessage struct {
	msgBase
	vals *map[string]string
}

func Debug(place, label string, vals *map[string]string) *DebugMessage {
	base := msgBase{typ: DEBUG, ruid: uuid.New(), place: place, label: label}
	return &DebugMessage{base, vals}
}

func (*DebugMessage) LogType() LogType {
	return DEBUG
}

func (msg *DebugMessage) JSON() string {
	return ""
}
