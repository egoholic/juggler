package lmsg

type CPUUsgData struct{}
type CPUUsgMessage struct {
	msgBase
	data *CPUUsgData
}

func (*CPUUsgMessage) LogType() LogType {
	return CPUUSG
}
func (msg *CPUUsgMessage) JSON() string {
	return ""
}
func CPUUsg() {

}
