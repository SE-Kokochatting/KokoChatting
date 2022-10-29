package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/model/req"
	"KokoChatting/model/res"
	"KokoChatting/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageController struct{
	baseController
	ManageService *service.ManageService
	UserService *service.UserService
}

func (manageCtl *ManageController) DeleteFriend (c *gin.Context) {
	delFriendReq := &req.DeleteFriendReq{}
	err := c.BindJSON(delFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		return
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

func (manageCtl *ManageController) BlockFriend (c *gin.Context) {
	blockFriendReq := &req.BlockFriendReq{}
	err := c.BindJSON(blockFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		return
	}
	//从报文首部中获取uid
	uid := manageCtl.getUid(c)

	err = manageCtl.ManageService.BlockFriend(uid, blockFriendReq.Fid)
	if err != nil{
		manageCtl.WithErr(global.BlockFriendError, c)
		return
	}

	blockFriendRes := &res.BlockFriendRes{}
	manageCtl.WithData(blockFriendRes, c)
}

func (manageCtl *ManageController) CreatGroup (c *gin.Context) {
	creatGroupReq := &req.CreatGroupReq{}
	err := c.BindJSON(creatGroupReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		return
	}
	//从报文首部中获取uid
	uid := manageCtl.getUid(c)

	gid, err1 := manageCtl.ManageService.CreatGroup(creatGroupReq.Name, uid, creatGroupReq.Administrator, creatGroupReq.Member)
	if err1 != nil{
		manageCtl.WithErr(global.CreatGroupError, c)
		return
	}

	creatGroupRes := &res.CreatGroupRes{}
	creatGroupRes.Data.Gid = gid
	manageCtl.WithData(creatGroupRes, c)
}

func (manageCtl *ManageController) QuitGroup (c *gin.Context) {
	quitGroupQeq := &req.QuitGroupReq{}
	err := c.BindJSON(quitGroupQeq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		return
	}
	//从报文首部中获取uid
	uid := manageCtl.getUid(c)

	err = manageCtl.ManageService.QuitGroup(uid, quitGroupQeq.Gid)
	if err != nil{
		manageCtl.WithErr(global.QuitGroupError, c)
		return
	}

	quitGroupRes := &res.QuitGroupRes{}
	manageCtl.WithData(quitGroupRes, c)
}

func (manageCtl *ManageController) GetFriendListInfo (c *gin.Context) {
	uid := manageCtl.getUid(c)

	friend, err := manageCtl.ManageService.GetFriendList(uid)
	if err != nil{
		manageCtl.WithErr(global.GetFriendListError, c)
		return
	}

	friendListRes := &res.FriendListRes{}

	for i := range friend{
		userProfile := &dataobject.UserProfile{
			Uid: friend[i],
		}

		err := manageCtl.UserService.GetUserInfo(friend[i], userProfile)
		if err != nil {
			global.Logger.Error("get user info error", zap.Error(err))
			manageCtl.WithErr(global.GetFriendInfoError, c)
			return
		}
		friendListRes.Data.Friend = append(friendListRes.Data.Friend, res.User{
			Uid: userProfile.Uid,
			Name: userProfile.Name,
			AvatarUrl: userProfile.AvatarUrl,
		})
		//friendListRes.Data.Friend[i].Uid = userProfile.Uid
		//friendListRes.Data.Friend[i].Name = userProfile.Name
		//friendListRes.Data.Friend[i].AvatarUrl = userProfile.AvatarUrl
	}

	manageCtl.WithData(friendListRes, c)
}

func NewManageController() *ManageController {
	return &ManageController{
		ManageService: service.NewManageService(),
		UserService: service.NewUserService(),
	}
}