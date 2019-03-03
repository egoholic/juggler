package writer

type DebugMessage struct {
	msgBase
	vals map[string]string
}

func (*DebugMessage) LogLevel() LogLevel {
	return DEBUG
}
func (msg *DebugMessage) JSON() string {
	return ""
}

func Debug() {

}
