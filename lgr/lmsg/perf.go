package lmsg

type PerfData struct{}
type PerfMessage struct {
	msgBase
	data *AsyncRequestData
}

func (*PerfMessage) LogType() LogType {
	return PERF
}
func (msg *PerfMessage) JSON() string {
	return ""
}
func Perf() {

}
