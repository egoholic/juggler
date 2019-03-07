package lmsg

type SyncRequestData struct{}
type SyncRequestMessage struct {
	msgBase
	data *SyncRequestData
}

func (*SyncRequestMessage) LogType() LogType {
	return SYNCR
}
func (msg *SyncRequestMessage) JSON() string {
	return ""
}
func Syncr() {

}
