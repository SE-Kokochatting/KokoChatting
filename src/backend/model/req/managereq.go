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
