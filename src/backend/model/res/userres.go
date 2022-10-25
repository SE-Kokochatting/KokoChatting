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

type UserInfo struct {
	Data struct {
		Uid uint64 `json:"uid"`
		Name string `json:"name"`
		AvatarUrl string `json:"avatarUrl"`
	}
}

