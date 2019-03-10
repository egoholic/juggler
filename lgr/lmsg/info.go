package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type InfoMessage struct {
	msgBase
	msg string
}

func Info(place, label, msg string) *InfoMessage {
	base := msgBase{typ: INFO, ruid: uuid.New(), place: place, label: label}
	return &InfoMessage{base, msg}
}

func (*InfoMessage) LogType() LogType {
	return INFO
}

func (msg *InfoMessage) JSON() string {
	return ""
}
