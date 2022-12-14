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
	MessageService *service.MessageService
}

func (manageCtl *ManageController) DeleteFriend (c *gin.Context) {
	delFriendReq := &req.DeleteFriendReq{}
	err := c.BindJSON(delFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
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

func (manageCtl *ManageController) AgreeFriendRequest (c *gin.Context) {
	agreeFriendReq := &req.AgreeFriendReq{}
	err := c.BindJSON(agreeFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}
	uid := manageCtl.getUid(c)

	fid, t, err := manageCtl.ManageService.GetFromIdByMsgId(agreeFriendReq.Id)
	if t != global.FriendRequestNotify{
		global.Logger.Error("message type err", zap.Error(err))
		manageCtl.WithErr(global.MessageTypeError, c)
		return
	}
	if err != nil{
		global.Logger.Error("get from id err", zap.Error(err))
		manageCtl.WithErr(global.DatabaseQueryError, c)
		return
	}

	err = manageCtl.ManageService.AddFriend(uid, fid)
	if err != nil{
		global.Logger.Error("add friend err", zap.Error(err))
		manageCtl.WithErr(global.AgreeFriendError, c)
		return
	}

	err = manageCtl.MessageService.DeleteMessage(agreeFriendReq.Id)
	if err != nil{
		global.Logger.Error("delete message err", zap.Error(err))
		manageCtl.WithErr(global.MessageDeleteError, c)
		return
	}

	agreeFriendRes := &res.AgreeFriendRes{}
	manageCtl.WithData(agreeFriendRes, c)
}

func (manageCtl *ManageController) RefuseFriendRequest (c *gin.Context) {
	refuseFriendReq := &req.RefuseFriendReq{}
	err := c.BindJSON(refuseFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}

	err = manageCtl.MessageService.DeleteMessage(refuseFriendReq.Id)
	if err != nil{
		global.Logger.Error("delete message err", zap.Error(err))
		manageCtl.WithErr(global.MessageDeleteError, c)
		return
	}

	refuseFriendRes := &res.RefuseFriendRes{}
	manageCtl.WithData(refuseFriendRes, c)
}

func (manageCtl *ManageController) BlockFriend (c *gin.Context) {
	blockFriendReq := &req.BlockFriendReq{}
	err := c.BindJSON(blockFriendReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
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
		manageCtl.WithErr(global.RequestFormatError, c)
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
		manageCtl.WithErr(global.RequestFormatError, c)
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

func (manageCtl *ManageController) RemoveMember (c *gin.Context) {
	removeMemberReq := &req.RemoveMemberReq{}
	err := c.BindJSON(removeMemberReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}

	admin := manageCtl.getUid(c)

	is, err := manageCtl.ManageService.RemoveMember(admin, removeMemberReq.Uid, removeMemberReq.Gid)
	if is != true{
		global.Logger.Error("the user has no permission", zap.Error(err))
		manageCtl.WithErr(global.PermissionError, c)
		return
	}

	if err != nil{
		global.Logger.Error("remove member error", zap.Error(err))
		manageCtl.WithErr(global.RemoveMemberError, c)
		return
	}
	removeMemberRes := &res.RemoveMemberRes{}

	manageCtl.WithData(removeMemberRes, c)
}

func (manageCtl *ManageController) GetFriendListInfo (c *gin.Context) {
	uid := manageCtl.getUid(c)

	friend, err := manageCtl.ManageService.GetFriendList(uid)
	if err != nil{
		manageCtl.WithErr(global.GetFriendListError, c)
		return
	}

	friendListRes := &res.FriendListRes{}

	for i := range friend {
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

func (manageCtl *ManageController) GetGroupListInfo (c *gin.Context) {
	uid := manageCtl.getUid(c)

	group, err := manageCtl.ManageService.GetGroupList(uid)
	if err != nil{
		manageCtl.WithErr(global.GetGroupError, c)
		return
	}

	groupListRes := &res.GroupListRes{}

	for i := range group{
		groupProfile := &dataobject.GroupProfile{
			Gid: group[i],
		}

		err := manageCtl.ManageService.GetGroupInfo(groupProfile)
		if err != nil {
			global.Logger.Error("get group info error", zap.Error(err))
			manageCtl.WithErr(global.GetGroupInfoError, c)
			return
		}
		groupListRes.Data.Group = append(groupListRes.Data.Group, res.GroupInfo{
			Gid: groupProfile.Gid,
			Name: groupProfile.Name,
			AvatarUrl: groupProfile.AvatarUrl,
		})
	}

	manageCtl.WithData(groupListRes, c)
}

func (manageCtl *ManageController) SetGroupAvatar (c *gin.Context) {
	uid := manageCtl.getUid(c)

	groupSetAvatarReq := &req.GroupSetAvatarReq{}
	if err := c.BindJSON(groupSetAvatarReq); err != nil {
		global.Logger.Error("set group avatar bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}

	is, err := manageCtl.ManageService.SetGroupAvatar(uid, groupSetAvatarReq.Gid, groupSetAvatarReq.AvatarUrl)
	if is != true{
		global.Logger.Error("have no permission", zap.Error(err))
		manageCtl.WithErr(global.PermissionError, c)
		return
	}
	if err != nil{
		global.Logger.Error("set group avatar error", zap.Error(err))
		manageCtl.WithErr(global.SetGroupAvatarError, c)
		return
	}

	groupSetAvatarRes := &res.GroupSetAvatarRes{}

	manageCtl.WithData(groupSetAvatarRes, c)
}

func (manageCtl *ManageController) TransferHost (c *gin.Context) {
	transferHostReq := &req.TransferHostReq{}
	err := c.BindJSON(transferHostReq)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}
	host := manageCtl.getUid(c)

	err = manageCtl.ManageService.TransferHost(host, transferHostReq.Gid, transferHostReq.Uid)
	if err != nil{
		global.Logger.Error("transfer host err", zap.Error(err))
		manageCtl.WithErr(global.TransferHostError, c)
		return
	}

	err = manageCtl.ManageService.TransferMember(host, transferHostReq.Gid, host)
	if err != nil{
		global.Logger.Error("transfer member err", zap.Error(err))
		manageCtl.WithErr(global.TransferMemError, c)
		return
	}

	transferHostRes := &res.TransferHostRes{}

	manageCtl.WithData(transferHostRes, c)
}

func (manageCtl *ManageController) ChangePermission (c *gin.Context) {
	changePermission := &req.ChangePermissionReq{}
	err := c.BindJSON(changePermission)
	if err != nil{
		global.Logger.Error("bind json error", zap.Error(err))
		manageCtl.WithErr(global.RequestFormatError, c)
		return
	}
	host := manageCtl.getUid(c)

	is, err := manageCtl.ManageService.IsMember(changePermission.Gid, changePermission.Uid)
	if is == true{
		if err := manageCtl.ManageService.TransferAdmin(host, changePermission.Gid, changePermission.Uid); err != nil{
			global.Logger.Error("transfer admin err", zap.Error(err))
			manageCtl.WithErr(global.TransferAdminError, c)
			return
		}
	}else{
		if err := manageCtl.ManageService.TransferMember(host, changePermission.Gid, changePermission.Uid); err != nil {
			global.Logger.Error("transfer member err", zap.Error(err))
			manageCtl.WithErr(global.TransferMemError, c)
			return
		}
	}

	changePermissionRes := &res.ChangePermissionRes{}

	manageCtl.WithData(changePermissionRes, c)
}

func NewManageController() *ManageController {
	return &ManageController{
		ManageService: service.NewManageService(),
		UserService: service.NewUserService(),
		MessageService: service.NewMessageService(),
	}
}