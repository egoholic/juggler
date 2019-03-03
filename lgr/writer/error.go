package writer

type ErrorMessage struct {
	msgBase
	err error
}

func (*ErrorMessage) LogLevel() LogLevel {
	return ERROR
}
func (msg *ErrorMessage) JSON() string {
	return ""
}
func Error() {

}
