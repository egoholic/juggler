package writer

type SyncRequestData struct{}
type SyncRequestMessage struct {
	msgBase
	data *SyncRequestData
}

func (*SyncRequestMessage) LogLevel() LogLevel {
	return SYNCR
}
func (msg *SyncRequestMessage) JSON() string {
	return ""
}
func Syncr() {

}
