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
		return err
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
		return err
	}
	return nil
}

func (managePro *ManageProvider) CreateGroup (name string) (uint64, error) {
	var groupEntity = &dataobject.GroupProfile{
		Name: name,
	}

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Create(groupEntity).Error
	if err != nil{
		global.Logger.Error("creat group error",zap.Error(err))
		return 0, err
	}

	return groupEntity.Gid, nil
}

func (managePro *ManageProvider) CreateMember (gid uint64, uid uint64, isAdmin bool, isHost bool) error {
	var creatMemberEntity = &dataobject.GroupMember{
		Gid: gid,
		Uid: uid,
		IsAdmin: isAdmin,
		IsHost: isHost,
	}

	dbClient := managePro.mysqlProvider.mysqlDb
	
	err := dbClient.Create(creatMemberEntity).Error
	if err != nil{
		global.Logger.Error("creat member error",zap.Error(err))
		return err
	}
	return nil
}

func (managePro *ManageProvider) QuitGroup (uid uint64, gid uint64) error {

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Where("uid = ? AND gid = ?", uid, gid).Delete(dataobject.GroupMember{}).Error
	if err != nil{
		global.Logger.Error("quit group error",zap.Error(err))
		return err
	}
	return nil
}

func (managePro *ManageProvider) GetFriendList (uid uint64) ([]uint64, error) {
	var friendRelationEntity = &[]dataobject.FriendRelation{}
	var friend []uint64

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Where("user1 = ? OR user2 = ?", uid, uid).Find(friendRelationEntity).Error
	if err != nil{
		global.Logger.Error("get friend error",zap.Error(err))
		return friend, err
	}

	for _,relation := range *friendRelationEntity{
		if relation.User1 == uid{
			friend = append(friend,relation.User2)
		}
		if relation.User2 == uid{
			friend = append(friend,relation.User1)
		}
	}

	return friend, nil
}

func (managePro *ManageProvider) GetGroupList (uid uint64) ([]uint64, error) {
	var MemberEntity = &[]dataobject.GroupMember{}
	group := make([]uint64, 0)

	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Where("uid = ?", uid).Find(MemberEntity).Error
	if err != nil{
		global.Logger.Error("get group error",zap.Error(err))
	}
	for _,member := range *MemberEntity{
		group = append(group, member.Gid)
	}

	return group, nil
}

func (managePro *ManageProvider) UpdateGroupInfo (newProfile *dataobject.GroupProfile) error {
	dbClient := managePro.mysqlProvider.mysqlDb

	groupProfile := &dataobject.GroupProfile{}
	err := dbClient.Where("gid = ?", newProfile.Gid).First(groupProfile).Error
	if err != nil{
		global.Logger.Error("update group error: group is not exist", zap.Error(err))
		return err
	}

	err = dbClient.Model(&dataobject.GroupProfile{}).Where("gid = ?", newProfile.Gid).Updates(dataobject.GroupProfile{
		Name: newProfile.Name,
		AvatarUrl: newProfile.AvatarUrl,
	}).Error
	if err != nil {
		global.Logger.Error("update group error: can not update", zap.Error(err))
		return err
	}

	return nil
}
// VerifyPermission 返回是否是管理员或群主，是则返回true，不是则返回false
func (managePro *ManageProvider) VerifyPermission (uid uint64, gid uint64) bool {
	dbClient := managePro.mysqlProvider.mysqlDb

	memberProfile := &dataobject.GroupMember{
		Gid: gid,
		Uid: uid,
	}
	err := dbClient.Where("uid = ? and gid = ?", memberProfile.Uid, memberProfile.Gid).Find(memberProfile).Error
	if err != nil{
		global.Logger.Error("the user is not in this group", zap.Error(err))
		return false
	}

	if memberProfile.IsHost == false && memberProfile.IsAdmin == false {
		return false
	}
	return true
}

func (managePro *ManageProvider) GetGroupInfo (groupProfile *dataobject.GroupProfile) error {
	dbClient := managePro.mysqlProvider.mysqlDb

	err := dbClient.Where("gid = ?", groupProfile.Gid).First(groupProfile).Error
	if err != nil{
		global.Logger.Error("get group information error", zap.Error(err))
		return err
	}
	return err
}

func (managePro *ManageProvider) GetUserIdOfGroup (gid uint64) ([]uint64, error){
	dbClient := managePro.mysqlProvider.mysqlDb
	type user struct {
		Uid uint64
	}
	uid := make([]uint64, 0)
	userList := make([]user, 0)

	err := dbClient.Table("group_members").Select("uid").Where("gid = ?", gid).Find(&userList).Error
	if err != nil{
		global.Logger.Error("find member error", zap.Error(err))
		return uid, err
	}
	for _,u := range userList {
		uid = append(uid, u.Uid)
	}

	return uid, nil
}

func (managePro *ManageProvider) IsInGroup (uid uint64, gid uint64) (bool, error) {
	dbClient := managePro.mysqlProvider.mysqlDb

	memberProfile := &dataobject.GroupMember{
		Gid: gid,
		Uid: uid,
	}
	err := dbClient.Where("uid = ? and gid = ?", memberProfile.Uid, memberProfile.Gid).Find(memberProfile).Error
	if err != nil{
		global.Logger.Error("the user is not in this group", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (managePro *ManageProvider) IsInBlock (user uint64, blocker uint64) (bool, error) {
	dbClient := managePro.mysqlProvider.mysqlDb

	blockRelation := &dataobject.BlockRelation{
		User: user,
		Blocker: blocker,
	}
	err := dbClient.Where("user = ? and blocker = ?", blockRelation.User, blockRelation.Blocker).Find(blockRelation).Error
	if err != nil{
		global.Logger.Error("the blocker is not blocked", zap.Error(err))
		return false, err
	}

	return true, nil
}

func NewManageProvider() *ManageProvider {
	return &ManageProvider{
		mysqlProvider: *NewMysqlProvider(),
	}
}