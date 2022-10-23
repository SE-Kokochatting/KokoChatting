package service

import (
	"KokoChatting/global"
	"KokoChatting/provider"
	"go.uber.org/zap"
)

type DeleteFriendService struct{
	deleteFriendProvider *provider.DeleteFriendProvider
}

func(delFriendSrv *DeleteFriendService) DeleteFriend(uid uint64,fid uint64) error {
	err := delFriendSrv.deleteFriendProvider.DeleteFriend(uid,fid)
	if err != nil{
		global.Logger.Error("delete friend err",zap.Error(err))
		return err
	}
	return err
}

func NewDeleteFriendService() *DeleteFriendService {
	return &DeleteFriendService{
		deleteFriendProvider: provider.NewDeleteFriendProvider(),
	}
}