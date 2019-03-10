package lmsg

import (
	"time"

	"github.com/egoholic/juggler/lgr/lmsg/uuid"
)

type syncRequestData struct {
	sender      string
	receiver    string
	description string
	latency     time.Duration
}

type SyncRequestMessage struct {
	msgBase
	syncRequestData
}

func Syncr(place, label, sender, receiver, description string, latency time.Duration) *SyncRequestMessage {
	base := msgBase{typ: SYNCR, ruid: uuid.New(), place: place, label: label}
	data := syncRequestData{sender: sender, receiver: receiver, description: description, latency: latency}

	return &SyncRequestMessage{base, data}
}

func (*SyncRequestMessage) LogType() LogType {
	return SYNCR
}

func (msg *SyncRequestMessage) JSON() string {
	return ""
}
