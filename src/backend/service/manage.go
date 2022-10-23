package service

import (
	"KokoChatting/global"
	"KokoChatting/provider"
	"go.uber.org/zap"
)

type ManageService struct{
	ManageProvider *provider.ManageProvider
}

func(manageSrv *ManageService) DeleteFriend(uid uint64,fid uint64) error {
	err := manageSrv.ManageProvider.DeleteFriend(uid,fid)
	if err != nil{
		global.Logger.Error("delete friend err",zap.Error(err))
		return err
	}
	return err
}

func(manageSrv *ManageService) BlockFriend(uid uint64,fid uint64) error {
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

func NewManageService() *ManageService {
	return &ManageService{
		ManageProvider: provider.NewManageProvider(),
	}
}