package service

import (
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/provider"
)

type MsgPullService struct {
	msgPullPro *provider.MsgPullProvider
}

func (msgPullSrv *MsgPullService) PullOutlineMsg(uid uint64, pullReq *req.PullMsgReq, pullOutlineRes *res.PullOutlineMsgRes) {

	msgPullSrv.msgPullPro.PullOutlineMsg()

}

func NewMsgPullService() *MsgPullService{
	return &MsgPullService{
		msgPullPro: provider.NewMsgPullProvider(),
	}
}