package req

type UserRegisterReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Uid uint64 `json:"uid"`
	Password string `json:"password"`
}

type UserInfoReq struct {
	Uid uint64 `json:"uid"`
}

type UserSetAvatarReq struct {
	AvatarUrl string `gorm:"column:avatarUrl" json:"avatarUrl"`
}