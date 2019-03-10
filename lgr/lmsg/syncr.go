package lmsg

import (
	"time"
)

type syncRequestData struct {
	sender      string
	receiver    string
	description string
	latency     time.Duration
}

type SyncRequestMessage struct {
	meta *Meta
	syncRequestData
}

func Syncr(meta *Meta, sender, receiver, description string, latency time.Duration) *SyncRequestMessage {
	data := syncRequestData{sender: sender, receiver: receiver, description: description, latency: latency}
	return &SyncRequestMessage{meta, data}
}

func (*SyncRequestMessage) LogType() LogType {
	return SYNCR
}

func (msg *SyncRequestMessage) JSON() string {
	return ""
}
