package filter

import "github.com/egoholic/juggler/lgr/lmsg"

var ALL_LLS = lmsg.LOG_LEVELS
var MINIMAL = []lmsg.LogType{lmsg.PANIC, lmsg.ERROR, lmsg.SYNCR, lmsg.ASYNCR}

type Filter struct {
	ltypes []lmsg.LogType
}

func New(lls []lmsg.LogType) *Filter {
	return &Filter{lls}
}

func (f *Filter) Check(mlt lmsg.LogType) bool {
	for _, lt := range f.ltypes {
		if lt == mlt {
			return true
		}
	}

	return false
}
