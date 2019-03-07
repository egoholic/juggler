package lmsg

type DebugMessage struct {
	msgBase
	vals map[string]string
}

func (*DebugMessage) LogType() LogType {
	return DEBUG
}
func (msg *DebugMessage) JSON() string {
	return ""
}

func Debug() {

}
