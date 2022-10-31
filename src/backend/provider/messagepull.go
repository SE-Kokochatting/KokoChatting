package provider

type MsgPullProvider struct {
	mysqlProvider
}

func (m *MsgPullProvider) PullOutlineMsg() {

}

func NewMsgPullProvider() *MsgPullProvider{
	return &MsgPullProvider{

	}
}