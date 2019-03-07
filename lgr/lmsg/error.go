package lmsg

type ErrorMessage struct {
	msgBase
	err error
}

func (*ErrorMessage) LogType() LogType {
	return ERROR
}
func (msg *ErrorMessage) JSON() string {
	return ""
}
func Error() {

}
