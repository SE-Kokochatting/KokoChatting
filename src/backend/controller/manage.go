package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteFriendController struct{
	baseController
	deleteFriendService *service.DeleteFriendService
}

func (delFriendCtl *DeleteFriendController) DeleteFriend (c *gin.Context) {
	delFriendReq := &req.DeleteFriendReq{}
	err := c.BindJSON(delFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
	}

	//用户id应由token解析出来，这里先写死用于测试功能
	var uid uint64
	uid = 100005

	err = delFriendCtl.deleteFriendService.DeleteFriend(uid, delFriendReq.Fid)
	if err != nil{
		delFriendCtl.WithErr(global.DeleteFriendError, c)
		return
	}

	delFriendRes := &res.DelFriendRes{}
	delFriendCtl.WithData(delFriendRes, c)
}

func NewDeleteFriendController() *DeleteFriendController {
	return &DeleteFriendController{
		deleteFriendService: service.NewDeleteFriendService(),
	}
}