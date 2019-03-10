package lmsg

import "github.com/egoholic/juggler/lgr/lmsg/uuid"

type memData struct {
	objCount  *map[string]int8
	memBefore int64
	memAfter  int64
}

type MemMessage struct {
	msgBase
	memData
}

func MemUsg(place, label string, objCount *map[string]int8, memBefore, memAfter int64) *MemMessage {
	base := msgBase{typ: MEMUSG, ruid: uuid.New(), place: place, label: label}
	memData := memData{objCount: objCount, memBefore: memBefore, memAfter: memAfter}
	return &MemMessage{base, memData}
}

func (*MemMessage) LogType() LogType {
	return MEMUSG
}
func (msg *MemMessage) JSON() string {
	return ""
}
