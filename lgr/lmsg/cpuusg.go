package lmsg

type CPUUsgMessage struct {
	meta       *Meta
	percentage int
}

func CPUUsg(meta *Meta, percentage int) *CPUUsgMessage {
	return &CPUUsgMessage{meta, percentage}
}

func (*CPUUsgMessage) LogType() LogType {
	return CPUUSG
}

func (msg *CPUUsgMessage) JSON() string {
	return ""
}
