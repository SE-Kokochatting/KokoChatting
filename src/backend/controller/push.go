package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PushController struct{
	baseController
	wssrv *service.WsService
	msgsrv *service.MessageService
}

func (controller *PushController) SendMsg(c *gin.Context){
	uid := controller.getUid(c)
	req := new(req.SendMsgReq)
	if err := c.BindJSON(req);err != nil{
		global.Logger.Error("request format err",zap.Error(err))
		controller.WithErr(global.RequestFormatError,c)
		return
	}

	// wrap msg
	wrapMsg,err := controller.msgsrv.WrapCommonMessage(uid,req.Reciever,req.MessageContent,req.MessageType)
	if err != nil{
		global.Logger.Error("wrap msg error",zap.Error(err))
		controller.WithErr(err,c)
		return
	}

	// store msg
	msgid,err := controller.msgsrv.StoreMessage(uid,req.Reciever,req.MessageContent,req.MessageType)
	if err != nil{
		global.Logger.Error("store msg err",zap.Error(err))
		controller.WithErr(err,c)
		return
	}

	// push msg
	err = controller.wssrv.SendMessage(wrapMsg)
	if err != nil{
		global.Logger.Error("send msg error",zap.Error(err))
		controller.WithErr(err,c)
		return
	}
	res := &res.SendMessageRes{}
	res.Data.Msgid = msgid
	controller.WithData(res,c)
}

func NewPushController()*PushController{
	return &PushController{
		baseController: baseController{},
		wssrv: new(service.WsService),
		msgsrv: service.NewMessageService(),
	}
}