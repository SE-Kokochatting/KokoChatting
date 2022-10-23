package req

type DeleteFriendReq struct {
	Fid uint64 `json:"fid"`
}

type BlockFriendReq struct {
	Fid uint64 `json:"fid"`
}