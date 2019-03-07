package lmsg

type MemData struct{}
type MemMessage struct {
	msgBase
	data *MemData
}

func (*MemMessage) LogType() LogType {
	return MEMUSG
}
func (msg *MemMessage) JSON() string {
	return ""
}
func MemUsg() {

}
