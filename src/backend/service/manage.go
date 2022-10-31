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
}

func (manageSrv *ManageService) DeleteFriend (uid uint64, fid uint64) error {
	err := manageSrv.ManageProvider.DeleteFriend(uid,fid)
	if err != nil{
		global.Logger.Error("delete friend err",zap.Error(err))
		return err
	}
	return err
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
	return gid, err
}

func (manageSrv *ManageService) QuitGroup (uid uint64, gid uint64) error {
	err := manageSrv.ManageProvider.QuitGroup(uid, gid)
	if err != nil{
		global.Logger.Error("quit group err", zap.Error(err))
		return err
	}
	return err
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

	is := manageSrv.ManageProvider.VerifyPermission(uid, gid)
	if is != true{
		return is, errors.New("the user have no permission")
	}

	err := manageSrv.ManageProvider.UpdateGroupInfo(newProfile)
	if err != nil{
		global.Logger.Error("update avatar err", zap.Error(err))
		return is, err
	}
	return is, err
}

func NewManageService() *ManageService {
	return &ManageService{
		ManageProvider: provider.NewManageProvider(),
	}
}