package writer

type PerfData struct{}
type PerfMessage struct {
	msgBase
	data *AsyncRequestData
}

func (*PerfMessage) LogLevel() LogLevel {
	return PERF
}
func (msg *PerfMessage) JSON() string {
	return ""
}
func Perf() {

}
