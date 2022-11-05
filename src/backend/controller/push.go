package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PushController struct {
	baseController
	wssrv  *service.WsService
	msgsrv *service.MessageService
}

func (controller *PushController) SendMsg(c *gin.Context) {
	uid := controller.getUid(c)
	_req := new(req.SendMsgReq)
	if err := c.BindJSON(_req); err != nil {
		global.Logger.Error("request format err", zap.Error(err))
		controller.WithErr(global.RequestFormatError, c)
		return
	}
	// store msg
	msgid, err := controller.msgsrv.StoreMessage(uid, _req.Receiver, _req.MessageContent, _req.MessageType)
	if err != nil {
		global.Logger.Error("store msg err", zap.Error(err))
		controller.WithErr(err, c)
		return
	}

	// wrap msg
	wrapMsg, err := controller.msgsrv.WrapCommonMessage(uid, _req.Receiver, _req.MessageContent, _req.MessageType,msgid)
	if err != nil {
		global.Logger.Error("wrap msg error", zap.Error(err))
		controller.WithErr(err, c)
		return
	}



	// push msg
	err = controller.wssrv.SendMessage(wrapMsg)
	if err != nil {
		global.Logger.Error("send msg error", zap.Error(err))
		controller.WithErr(err, c)
		return
	}
	_res := &res.SendMessageRes{}
	_res.Data.Msgid = msgid
	controller.WithData(_res, c)
}

func (controller *PushController) RevertMessage(c *gin.Context) {
	_req := new(req.RevertMsgReq)
	if err := c.BindJSON(_req); err != nil {
		global.Logger.Error("request format error", zap.Error(err))
		controller.WithErr(global.RequestFormatError, c)
		return
	}
	uid := controller.getUid(c)
	err := controller.msgsrv.RevertMessage(uid, _req.MsgId)
	if err != nil {
		global.Logger.Error("revert message error", zap.Error(err))
		controller.WithErr(err, c)
		return
	}
	controller.WithData(nil, c)
}

func NewPushController() *PushController {
	return &PushController{
		baseController: baseController{},
		wssrv:          new(service.WsService),
		msgsrv:         service.NewMessageService(),
	}
}
