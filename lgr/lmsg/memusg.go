package lmsg

type memData struct {
	objCount  *map[string]int8
	memBefore int64
	memAfter  int64
}

type MemMessage struct {
	meta *Meta
	memData
}

func MemUsg(meta *Meta, objCount *map[string]int8, memBefore, memAfter int64) *MemMessage {
	memData := memData{objCount: objCount, memBefore: memBefore, memAfter: memAfter}
	return &MemMessage{meta, memData}
}

func (*MemMessage) LogType() LogType {
	return MEMUSG
}
func (msg *MemMessage) JSON() string {
	return ""
}
