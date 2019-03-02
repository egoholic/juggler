package writer

type MType int

const (
	_     = iota
	PANIC = MType(iota + 1)
	ERROR
	INFO
	DEBUG
	SYNCR
	ASYNCR
	PERF
	MEM
	CPUUSG
)

type Message interface {
	MType() MType
}

type msgBase struct {
	typ   MType
	ruid  string
	place string
	label string
}

type Writer struct {
	channel chan<- Message
}

func New() *Writer {
	return &Writer{}
}

type DebugMessage struct {
	msgBase
	vals map[string]string
}

func Debug() {

}

type InfoMessage struct {
	msgBase
	msg string
}

func Info() {

}

type ErrorMessage struct {
	msgBase
	err error
}

func Error() {

}

type PanicMessage struct {
	msgBase
	err error
}

func Panic() {

}

type SyncRequestData struct{}
type SyncRequestMessage struct {
	msgBase
	data *SyncRequestData
}

func Syncr() {

}

type AsyncRequestData struct{}
type AsyncRequestMessage struct {
	msgBase
	data *AsyncRequestData
}

func Asyncr() {

}

type PerfData struct{}
type PerfMessage struct {
	msgBase
	data *AsyncRequestData
}

func Perf() {

}

type MemData struct{}
type MemMessage struct {
	msgBase
	data *MemData
}

func Mem() {

}

type CPUUsgData struct{}
type CPUUsgMessage struct {
	msgBase
	data *CPUUsgData
}

func CPUUsg() {

}
