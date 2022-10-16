package res

type UserRes struct {
	BaseResponse
	Data struct {
		Uid uint64 `json:"uid"`
	} `json:"data"`
}