package provider

import (
	"KokoChatting/global"
	"KokoChatting/model/dataobject"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type ManageProvider struct {
	mysqlProvider
}

func (managePro *ManageProvider) DeleteFriend (uid uint64,fid uint64) error {
	var friendRelationEntity = &dataobject.FriendRelation{
		User1: uid,
		User2: fid,
	}
	friendRelationEntity.Preprocess()

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Where("user1 = ? AND user2 = ?",friendRelationEntity.User1,friendRelationEntity.User2).Find(friendRelationEntity).Error
	if err != nil && err == gorm.ErrRecordNotFound{
		global.Logger.Error("the user you want to delete is not your friend.", zap.Error(err))
		return  err
	}
	if err != nil && err != gorm.ErrRecordNotFound{
		global.Logger.Error("database find error.", zap.Error(err))
		return  err
	}

	err = dbClient.Delete(friendRelationEntity).Error
	if err != nil{
		global.Logger.Error("database delete error",zap.Error(err))
	}
	return nil
}

func (managePro *ManageProvider) BlockFriend (uid uint64, fid uint64) error{
	var blockRelationEntity = &dataobject.BlockRelation{
		User: uid,
		Blocker: fid,
	}

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Create(blockRelationEntity).Error
	if err != nil{
		global.Logger.Error("block friend error",zap.Error(err))
	}
	return nil
}

func NewManageProvider() *ManageProvider {
	return &ManageProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}