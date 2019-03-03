package writer

type AsyncRequestData struct{}
type AsyncRequestMessage struct {
	msgBase
	data *AsyncRequestData
}

func (*AsyncRequestMessage) LogLevel() LogLevel {
	return ASYNCR
}
func (msg *AsyncRequestMessage) JSON() string {
	return ""
}
func Asyncr() {

}
