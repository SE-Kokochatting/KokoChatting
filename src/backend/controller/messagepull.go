package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MsgPullController struct {
	baseController
	msgPullService *service.MsgPullService
}

// MsgPullOutline pull message outline
func (pullCtl *MsgPullController) MsgPullOutline(c *gin.Context) {
	uid := pullCtl.getUid(c)
	pullMsgReq := &req.PullOutlineMsgReq{}
	if err := c.BindJSON(pullMsgReq); err != nil {
		global.Logger.Error("pull message request bind err",zap.Error(err))
		pullCtl.WithErr(global.MessagePullBindError, c)
		return
	}

	pullMsgRes := &res.PullOutlineMsgRes{}
	if err := pullCtl.msgPullService.PullOutlineMsg(uid, pullMsgReq, pullMsgRes); err != nil {
		global.Logger.Error("pull outline message error", zap.Error(err))
		pullCtl.WithErr(global.PullOutlineError, c)
		return
	}

	pullCtl.WithData(pullMsgRes, c)
	return
}

func (pullCtl *MsgPullController) MsgPull(c *gin.Context) {
	uid := pullCtl.getUid(c)
	pullMsgReq := &req.PullMsgReq{}
	if err := c.BindJSON(pullMsgReq); err != nil {
		global.Logger.Error("pull message request bind err")
		pullCtl.WithErr(global.MessagePullBindError, c)
		return
	}

	if pullMsgReq.Id == uid {
		global.Logger.Error("can not pull message from yourself, id can not equal to uid")
		pullCtl.WithErr(global.MessageIdError, c)
		return
	}

	pullMsgRes, err := pullCtl.msgPullService.PullMsg(uid, pullMsgReq.LastMessageId, pullMsgReq.Id, pullMsgReq.MsgType)
	if err != nil {
		global.Logger.Error("pull message error")
		pullCtl.WithErr(global.PullMessageError, c)
		return
	}

	pullCtl.WithData(pullMsgRes, c)
	return
}

func (pullCtl *MsgPullController) MsgPullHistory(c *gin.Context) {
	uid := pullCtl.getUid(c)
	pullMsgHsyReq := &req.PullMsgHsyReq{}
	if err := c.BindJSON(pullMsgHsyReq); err != nil {
		global.Logger.Error("pull message request bind err")
		pullCtl.WithErr(global.MessagePullBindError, c)
		return
	}

	if pullMsgHsyReq.Id == uid {
		global.Logger.Error("can not pull message from yourself, id can not equal to uid")
		pullCtl.WithErr(global.MessageIdError, c)
		return
	}

	// 需要分页式拉取消息，并返回给前端
	pullMsgHsyRes, err := pullCtl.msgPullService.PullMsgHistory(uid, pullMsgHsyReq)
	if err != nil {
		global.Logger.Error("pull message error")
		pullCtl.WithErr(global.PullMessageError, c)
		return
	}

	pullCtl.WithData(pullMsgHsyRes, c)
	return
}

func NewMsgPullController() *MsgPullController {
	return &MsgPullController{
		msgPullService: service.NewMsgPullService(),
	}
}
