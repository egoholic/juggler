package lmsg

type PanicMessage struct {
	msgBase
	err error
}

func (*PanicMessage) LogType() LogType {
	return PANIC
}
func (msg *PanicMessage) JSON() string {
	return ""
}
func Panic() {

}
