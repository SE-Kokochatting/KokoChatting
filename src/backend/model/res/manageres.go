package res

type DelFriendRes struct {
	Data struct {} `json:"data"`
}

type BlockFriendRes struct {
	Data struct {} `json:"data"`
}

type CreatGroupRes struct {
	Data struct{
		Gid uint64 `json:"gid"`
	}`json:"data"`
}

type QuitGroupRes struct {
	Data struct {} `json:"data"`
}

type FriendListRes struct {
	Data struct{
		Friend []User `json:"friend"`
	}`json:"data"`
}

type User struct {
	Uid uint64 `json:"uid"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type GroupSetAvatarRes struct {
	Data struct {} `json:"data"`
}

type GroupListRes struct {
	Data struct{
		Group [] GroupInfo `json:"group"`
	} `json:"data"`
}

type GroupInfo struct {
	Gid uint64 `json:"gid"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type TransferHostRes struct {
	Data struct {} `json:"data"`
}

type ChangePermissionRes struct {
	Data struct {} `json:"data"`
}

type AgreeFriendRes struct {
	Data struct {} `json:"data"`
}

type RefuseFriendRes struct {
	Data struct {} `json:"data"`
}

type RemoveMemberRes struct {
	Data struct {} `json:"data"`
}