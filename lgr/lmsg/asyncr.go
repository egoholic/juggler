package lmsg

type AsyncRequestData struct{}
type AsyncRequestMessage struct {
	msgBase
	data *AsyncRequestData
}

func (*AsyncRequestMessage) LogType() LogType {
	return ASYNCR
}

func (msg *AsyncRequestMessage) JSON() string {
	return ""
}

func Asyncr() {

}
