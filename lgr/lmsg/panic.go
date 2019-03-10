package lmsg

type PanicMessage struct {
	meta *Meta
	err  error
}

func Panic(meta *Meta, err error) *PanicMessage {
	return &PanicMessage{meta, err}
}

func (*PanicMessage) LogType() LogType {
	return PANIC
}

func (msg *PanicMessage) JSON() string {
	return ""
}
