package dataobject

type StoreUserProfile struct {
	Uid uint64 `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

type StoreFriendProfile struct {
	Id uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	User1 uint64 `gorm:"column:user1"`
	User2 uint64 `gorm:"column:user2"`
}