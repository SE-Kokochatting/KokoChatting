package dataobject

type storeUserProfile struct {
	Uid uint64 `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}