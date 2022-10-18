package dataobject

// 本文件中的结构体与数据库中的结构体一一对应

type UserProfile struct {
	Uid uint64 `json:"uid"`
	Password string `json:"password"`
	Name string `json:"name"`
}