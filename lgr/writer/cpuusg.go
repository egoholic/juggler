package writer

type CPUUsgData struct{}
type CPUUsgMessage struct {
	msgBase
	data *CPUUsgData
}

func (*CPUUsgMessage) LogLevel() LogLevel {
	return CPUUSG
}
func (msg *CPUUsgMessage) JSON() string {
	return ""
}
func CPUUsg() {

}
