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

type msgBase struct {
	typ   LogType
	ruid  string
	place string
	label string
}
