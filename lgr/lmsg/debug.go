package lmsg

type DebugMessage struct {
	meta *Meta
	vals *map[string]string
}

func Debug(meta *Meta, vals *map[string]string) *DebugMessage {
	return &DebugMessage{meta, vals}
}

func (*DebugMessage) LogType() LogType {
	return DEBUG
}

func (msg *DebugMessage) JSON() string {
	return ""
}
