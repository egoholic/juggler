package writer

type LogLevel int

const (
	_     = iota
	PANIC = LogLevel(iota + 1)
	ERROR
	INFO
	DEBUG
	SYNCR
	ASYNCR
	PERF
	MEMUSG
	CPUUSG
)

var LogLevels = [9]LogLevel{PANIC, ERROR, INFO, DEBUG, SYNCR, ASYNCR, PERF, MEMUSG, CPUUSG}

type Writer struct {
	channel chan<- Message
}

func New() *Writer {
	return &Writer{}
}

type Message interface {
	LogLevel() LogLevel
	JSON() string
}

type msgBase struct {
	typ   LogLevel
	ruid  string
	place string
	label string
}
