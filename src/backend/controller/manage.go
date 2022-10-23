package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageController struct{
	baseController
	ManageService *service.ManageService
}

func (manageCtl *ManageController) DeleteFriend (c *gin.Context) {
	delFriendReq := &req.DeleteFriendReq{}
	err := c.BindJSON(delFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
	}
	//从报文首部中获取uid
	uid := manageCtl.getUid(c)

	err = manageCtl.ManageService.DeleteFriend(uid, delFriendReq.Fid)
	if err != nil{
		manageCtl.WithErr(global.DeleteFriendError, c)
		return
	}

	delFriendRes := &res.DelFriendRes{}
	manageCtl.WithData(delFriendRes, c)
}

func NewManageController() *ManageController {
	return &ManageController{
		ManageService: service.NewManageService(),
	}
}