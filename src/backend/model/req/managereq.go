package req

type DeleteFriendReq struct {
	Fid uint64 `json:"fid"`
}

type BlockFriendReq struct {
	Fid uint64 `json:"fid"`
}

type CreatGroupReq struct {
	Name string `json:"name"`
	Administrator []uint64 `json:"administrator"`
	Member []uint64 `json:"member"`
}

type QuitGroupReq struct {
	Gid uint64 `json:"gid"`
}

type GroupSetAvatarReq struct {
	Gid uint64 `json:"gid"`
	AvatarUrl string `gorm:"column:avatarUrl" json:"avatarUrl"`
}

type TransferHostReq struct {
	Uid uint64 `json:"uid"`
	Gid uint64 `json:"gid"`
}

type ChangePermissionReq struct {
	Uid uint64 `json:"uid"`
	Gid uint64 `json:"gid"`
}

type AgreeFriendReq struct {
	Id uint64 `json:"id"`
}