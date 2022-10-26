package res

type UserRegisterRes struct {
	Data struct {
		Uid uint64 `json:"uid"`
	} `json:"data"`
}

type UserLoginRes struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type UserInfoRes struct {
	Data struct {
		Uid uint64 `json:"uid"`
		Name string `json:"name"`
		AvatarUrl string `json:"avatarUrl"`
	} `json:"data"`
}

type UserSetAvatarRes struct {
	Data struct {} `json:"data"`
}

