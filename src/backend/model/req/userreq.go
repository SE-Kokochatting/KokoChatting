package req

type UserRegisterReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Uid uint64 `json:"uid"`
	Password string `json:"password"`
}