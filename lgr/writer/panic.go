package writer

type PanicMessage struct {
	msgBase
	err error
}

func (*PanicMessage) LogLevel() LogLevel {
	return PANIC
}
func (msg *PanicMessage) JSON() string {
	return ""
}
func Panic() {

}
