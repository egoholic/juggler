package lmsg

type InfoMessage struct {
	meta *Meta
	msg  string
}

func Info(meta *Meta, msg string) *InfoMessage {
	return &InfoMessage{meta, msg}
}

func (*InfoMessage) LogType() LogType {
	return INFO
}

func (msg *InfoMessage) JSON() string {
	return ""
}
