package lmsg

type ErrorMessage struct {
	meta *Meta
	err  error
}

func Error(meta *Meta, err error) *ErrorMessage {
	return &ErrorMessage{meta, err}
}

func (*ErrorMessage) LogType() LogType {
	return ERROR
}

func (msg *ErrorMessage) JSON() string {
	return ""
}
