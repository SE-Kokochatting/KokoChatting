package dataobject

// 本文件中的结构体与数据库中的结构体一一对应

type UserProfile struct {
	Uid      uint64 `json:"uid"`
	Password string `json:"password"`
	Name     string `json:"name"`
	// 由于gorm会自动解析成avatar_url，因此加上gorm自定义转换解析
	AvatarUrl string `gorm:"column:avatarUrl" json:"avatarUrl"`
}