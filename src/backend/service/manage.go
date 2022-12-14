package service

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"KokoChatting/provider"
	"errors"
	"go.uber.org/zap"
)

type ManageService struct{
	ManageProvider *provider.ManageProvider
	MessageProvider *provider.MessageProvider
	*MessageService
}

func (manageSrv *ManageService) DeleteFriend (uid uint64, fid uint64) error {
	err := manageSrv.ManageProvider.DeleteFriend(uid,fid)
	if err != nil{
		global.Logger.Error("delete friend err",zap.Error(err))
		return err
	}
	msg := "you have been deleted"
	err = manageSrv.PushUnStoredSystemMessage(uid, fid, msg, global.DeleteFriendNotify)
	if err != nil{
		global.Logger.Error("delete friend notify error", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) AddFriend (uid uint64, fid uint64) error {
	err := manageSrv.ManageProvider.AddFriend(uid, fid)
	if err != nil{
		global.Logger.Error("add friend err", zap.Error(err))
		return err
	}
	msg := "friend request is accepted"
	err = manageSrv.PushUnStoredSystemMessage(uid, fid, msg, global.AddFriendResponseNotify)
	if err != nil{
		global.Logger.Error("add friend response notify error", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) GetFromIdByMsgId (mid uint64) (uint64, int, error) {
	uid, t, err := manageSrv.MessageProvider.GetFromIdAndTypeByMsgId(mid)
	if err != nil{
		global.Logger.Error("get from id err", zap.Error(err))
		return 0, 0, err
	}
	return uid, t, err
}

func (manageSrv *ManageService) BlockFriend (uid uint64, fid uint64) error {
	err := manageSrv.ManageProvider.DeleteFriend(uid,fid)
	if err != nil{
		global.Logger.Error("delete friend err",zap.Error(err))
		return err
	}

	err = manageSrv.ManageProvider.BlockFriend(uid,fid)
	if err != nil{
		global.Logger.Error("block friend err",zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) CreatGroup (name string, uid uint64, aid []uint64, mid []uint64) ( uint64, error) {
	gid, err := manageSrv.ManageProvider.CreateGroup(name)
	if err != nil{
		global.Logger.Error("creat group err",zap.Error(err))
		return 0, err
	}

	err = manageSrv.ManageProvider.CreateMember(gid, uid, false, true)
	if err != nil{
		global.Logger.Error("creat member err",zap.Error(err))
		return 0, err
	}

	for i := range aid{
		err = manageSrv.ManageProvider.CreateMember(gid, aid[i], true, false)
		if err != nil{
			global.Logger.Error("creat member err",zap.Error(err))
			return 0, err
		}
	}

	for j := range mid{
		err = manageSrv.ManageProvider.CreateMember(gid, mid[j], false, false)
		if err != nil{
			global.Logger.Error("creat member err",zap.Error(err))
			return 0, err
		}
	}
	msg := "you have been in this group now"
	err = manageSrv.PushUnStoredSystemMessage(uid, gid, msg, global.JoinGroupNotify)
	if err != nil{
		global.Logger.Error("push join group notify error",zap.Error(err))
		return 0, err
	}
	return gid, err
}

func (manageSrv *ManageService) QuitGroup (uid uint64, gid uint64) error {
	msg := "this user quits the group"
	err := manageSrv.PushUnStoredSystemMessage(uid, gid, msg, global.QuitGroupNotify)
	if err != nil{
		global.Logger.Error("push quit group notify error",zap.Error(err))
		return err
	}
	err = manageSrv.ManageProvider.QuitGroup(uid, gid)
	if err != nil{
		global.Logger.Error("quit group err", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) RemoveMember (admin uint64, uid uint64, gid uint64) (bool, error) {

	can, err := manageSrv.ManageProvider.VerifyPermission(admin, gid)
	if can != true{
		global.Logger.Error("has no permission", zap.Error(err))
		return false, errors.New("the user has no permission")
	}

	isAdmin, err := manageSrv.ManageProvider.IsAdmin(admin, gid)
	isManage, err := manageSrv.ManageProvider.VerifyPermission(uid, gid)
	if isAdmin == true && isManage == true {
		global.Logger.Error("has no permission", zap.Error(err))
		return false, errors.New("the user has no permission")
	}

	err = manageSrv.ManageProvider.QuitGroup(uid, gid)
	if err != nil{
		global.Logger.Error("remove from group err", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (manageSrv *ManageService) GetFriendList (uid uint64) ([]uint64, error) {
	friend, err := manageSrv.ManageProvider.GetFriendList(uid)
	if err != nil{
		global.Logger.Error("get friend err", zap.Error(err))
		return friend, err
	}
	return friend, err
}

func (manageSrv *ManageService) SetGroupAvatar (uid uint64, gid uint64, avatarUrl string) (bool, error) {
	newProfile := &dataobject.GroupProfile{
		Gid: gid,
		AvatarUrl: avatarUrl,
	}

	is,err := manageSrv.ManageProvider.VerifyPermission(uid, gid)
	if is != true{
		return is, err
	}

	err = manageSrv.ManageProvider.UpdateGroupInfo(newProfile)
	if err != nil{
		global.Logger.Error("update avatar err", zap.Error(err))
		return is, err
	}
	return is, err
}

func (manageSrv *ManageService) GetGroupList (uid uint64) ([]uint64, error) {
	group, err := manageSrv.ManageProvider.GetGroupList(uid)
	if err != nil{
		global.Logger.Error("get group err", zap.Error(err))
		return group, err
	}
	return group, err
}

func (manageSrv *ManageService) GetGroupInfo (groupProfile *dataobject.GroupProfile) error {
	err := manageSrv.ManageProvider.GetGroupInfo(groupProfile)
	if err != nil{
		global.Logger.Error("get group information err", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) TransferHost (host uint64, gid uint64, uid uint64) error {
	is, err := manageSrv.ManageProvider.IsHost(host, gid)
	if is !=  true{
		global.Logger.Error("the user is not host", zap.Error(err))
		return err
	}

	groupMember := &dataobject.GroupMember{
		Uid: uid,
		Gid: gid,
		IsAdmin: false,
		IsHost: true,
	}

	err = manageSrv.ManageProvider.ChangeMemberPermission(groupMember)
	if err != nil{
		global.Logger.Error("change permission err", zap.Error(err))
	}
	return err
}

func (manageSrv *ManageService) TransferAdmin (host uint64, gid uint64, uid uint64) error {
	is, err := manageSrv.ManageProvider.IsHost(host, gid)
	if is !=  true{
		global.Logger.Error("the user is not host", zap.Error(err))
		return err
	}

	groupMember := &dataobject.GroupMember{
		Uid: uid,
		Gid: gid,
		IsAdmin: true,
		IsHost: false,
	}

	err = manageSrv.ManageProvider.ChangeMemberPermission(groupMember)
	if err != nil{
		global.Logger.Error("change permission err", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) TransferMember (host uint64, gid uint64, uid uint64) error {
	is, err := manageSrv.ManageProvider.IsHost(host, gid)
	if is !=  true{
		global.Logger.Error("the user is not host", zap.Error(err))
		return err
	}

	groupMember := &dataobject.GroupMember{
		Uid: uid,
		Gid: gid,
		IsAdmin: false,
		IsHost: false,
	}

	err = manageSrv.ManageProvider.ChangeMemberPermission(groupMember)
	if err != nil{
		global.Logger.Error("change permission err", zap.Error(err))
		return err
	}
	return err
}

func (manageSrv *ManageService) IsMember (gid uint64, uid uint64) (bool, error) {
	is, err := manageSrv.ManageProvider.IsMember(uid, gid)
	if is != true{
		global.Logger.Error("the user is not member", zap.Error(err))
		return false, err
	}

	return true, err
}

func NewManageService() *ManageService {
	return &ManageService{
		ManageProvider: provider.NewManageProvider(),
		MessageProvider: provider.NewMessageProvider(),
		MessageService: NewMessageService(),
	}
}