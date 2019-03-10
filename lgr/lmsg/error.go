package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type ErrorMessage struct {
	msgBase
	err error
}

func Error(place, label string, err error) *ErrorMessage {
	base := msgBase{typ: ERROR, ruid: uuid.New(), place: place, label: label}
	return &ErrorMessage{base, err}
}

func (*ErrorMessage) LogType() LogType {
	return ERROR
}

func (msg *ErrorMessage) JSON() string {
	return ""
}
