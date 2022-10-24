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