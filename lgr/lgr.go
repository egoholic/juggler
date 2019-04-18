package lgr

import (
	"time"

	"github.com/egoholic/juggler/lgr/lmsg"
)

type LogsWriter interface {
	WriteLog(string) error
}

type Logger struct {
	ruid   string
	writer LogsWriter
}

func New(ruid string, writer LogsWriter) *Logger {
	return &Logger{ruid, writer}
}

func (l *Logger) Asyncr(place, label, sender, receiver, description string) {
	meta := lmsg.NewMeta(lmsg.ASYNCR, l.ruid, place, label)
	msg := lmsg.Asyncr(meta, sender, receiver, description)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) CPUUsg(place, label string, percentage int) {
	meta := lmsg.NewMeta(lmsg.CPUUSG, l.ruid, place, label)
	msg := lmsg.CPUUsg(meta, percentage)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Debug(place, label string, vals *map[string]string) {
	meta := lmsg.NewMeta(lmsg.DEBUG, l.ruid, place, label)
	msg := lmsg.Debug(meta, vals)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Error(place, label string, err error) {
	meta := lmsg.NewMeta(lmsg.ERROR, l.ruid, place, label)
	msg := lmsg.Error(meta, err)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Info(place, label string, message string) {
	meta := lmsg.NewMeta(lmsg.INFO, l.ruid, place, label)
	msg := lmsg.Info(meta, message)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) MemUsg(place, label string, objCount *map[string]int8, memBefore, memAfter int64) {
	meta := lmsg.NewMeta(lmsg.INFO, l.ruid, place, label)
	msg := lmsg.MemUsg(meta, objCount, memBefore, memAfter)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Panic(place, label string, err error) {
	meta := lmsg.NewMeta(lmsg.INFO, l.ruid, place, label)
	msg := lmsg.Panic(meta, err)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Perf(place, label string, durations *map[string]time.Duration) {
	meta := lmsg.NewMeta(lmsg.INFO, l.ruid, place, label)
	msg := lmsg.Perf(meta, durations)
	l.writer.WriteLog(msg.JSON())
}

func (l *Logger) Syncr(place, label string, sender, receiver, description string, latency time.Duration) {
	meta := lmsg.NewMeta(lmsg.INFO, l.ruid, place, label)
	msg := lmsg.Syncr(meta, sender, receiver, description, latency)
	l.writer.WriteLog(msg.JSON())
}
