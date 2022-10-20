package dataobject

// 本文件中的结构体与数据库中的结构体一一对应

type UserProfile struct {
	Uid      uint64 `json:"uid"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type StoreFriendProfile struct {
	Id uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	User1 uint64 `gorm:"column:user1"`
	User2 uint64 `gorm:"column:user2"`
}