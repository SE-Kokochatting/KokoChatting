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