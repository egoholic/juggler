package lgr

import (
	"github.com/egoholic/juggler/lgr/lmsg"
)

type LogWriter interface {
	WriteLog(lmsg.Message) error
}

type Logger struct {
	ruid   string
	writer LogWriter
}

func New(ruid string, writer LogWriter) *Logger {
	return &Logger{ruid, writer}
}

func (l *Logger) Asyncr(place, label, sender, receiver, description string) {
	meta := lmsg.NewMeta(lmsg.ASYNCR, l.ruid, place, label)
	msg := lmsg.Asyncr(meta, sender, receiver, description)

	l.writer.WriteLog(msg)
}
