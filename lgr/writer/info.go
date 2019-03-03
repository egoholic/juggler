package writer

type InfoMessage struct {
	msgBase
	msg string
}

func (*InfoMessage) LogLevel() LogLevel {
	return INFO
}
func (msg *InfoMessage) JSON() string {
	return ""
}
func Info() {

}
