package res

type UserRes struct {
	BaseResponse
	Data struct {
		Uid uint64 `json:"uid"`
		// Token string `json:"token"`
	} `json:"data"`
}