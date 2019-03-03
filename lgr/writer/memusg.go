package writer

type MemData struct{}
type MemMessage struct {
	msgBase
	data *MemData
}

func (*MemMessage) LogLevel() LogLevel {
	return MEMUSG
}
func (msg *MemMessage) JSON() string {
	return ""
}
func MemUsg() {

}
