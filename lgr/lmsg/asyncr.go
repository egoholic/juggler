package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type asyncRequestData struct {
	sender      string
	receiver    string
	description string
}

type AsyncRequestMessage struct {
	msgBase
	asyncRequestData
}

func Asyncr(place, label, sender, receiver, description string) *AsyncRequestMessage {
	base := msgBase{typ: ASYNCR, ruid: uuid.New(), place: place, label: label}
	reqData := asyncRequestData{sender: sender, receiver: receiver, description: description}
	return &AsyncRequestMessage{base, reqData}
}

func (*AsyncRequestMessage) LogType() LogType {
	return ASYNCR
}

func (msg *AsyncRequestMessage) JSON() string {
	return ""
}
