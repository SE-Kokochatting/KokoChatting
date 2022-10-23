package res

type UserRegisterRes struct {
	Data struct {
		Uid uint64 `json:"uid"`
	} `json:"data"`
}

type UserLoginRes struct {
	Data struct {
		Token string `json:"token"`
	}
}