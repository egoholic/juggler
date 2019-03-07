package lmsg

type InfoMessage struct {
	msgBase
	msg string
}

func (*InfoMessage) LogType() LogType {
	return INFO
}
func (msg *InfoMessage) JSON() string {
	return ""
}
func Info() {

}
