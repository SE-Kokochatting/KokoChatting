package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
)

type MsgPullController struct {
	baseController
	msgPullService *service.MsgPullService
}

// MsgPullOutline pull message outline
func (pullCtl *MsgPullController) MsgPullOutline(c *gin.Context) {
	uid := pullCtl.getUid(c)
	pullMsgReq := &req.PullMsgReq{}
	if err := c.BindJSON(pullMsgReq); err != nil {
		global.Logger.Error("pull message request bind err")
		pullCtl.WithErr(global.MessagePullBindError, c)
		return
	}

	pullMsgRes := &res.PullOutlineMsgRes{}
	pullCtl.msgPullService.PullOutlineMsg(uid, pullMsgReq, pullMsgRes)

	pullCtl.WithData(pullMsgRes, c)
	return
}

func NewMsgPullController() *MsgPullController{
	return &MsgPullController{
		msgPullService: service.NewMsgPullService(),
	}
}