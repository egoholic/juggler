package lmsg

type LogType int

const (
	_     = iota
	PANIC = LogType(iota + 1)
	ERROR
	INFO
	DEBUG
	SYNCR
	ASYNCR
	PERF
	MEMUSG
	CPUUSG
)

var LOG_LEVELS = [9]LogType{PANIC, ERROR, INFO, DEBUG, SYNCR, ASYNCR, PERF, MEMUSG, CPUUSG}

type Message interface {
	LogType() LogType
	JSON() string
}

type Meta struct {
	typ   LogType
	ruid  string
	place string
	label string
}

func NewMeta(typ LogType, ruid, place, label string) *Meta {
	return &Meta{typ: typ, ruid: ruid, place: place, label: label}
}
