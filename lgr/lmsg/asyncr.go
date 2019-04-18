package lmsg

type asyncRequestData struct {
	sender      string
	receiver    string
	description string
}

type AsyncRequestMessage struct {
	meta *Meta
	asyncRequestData
}

func Asyncr(meta *Meta, sender, receiver, description string) *AsyncRequestMessage {
	reqData := asyncRequestData{sender: sender, receiver: receiver, description: description}
	return &AsyncRequestMessage{meta, reqData}
}

func (*AsyncRequestMessage) LogType() LogType {
	return ASYNCR
}

func (msg *AsyncRequestMessage) JSON() string {
	return ""
}
