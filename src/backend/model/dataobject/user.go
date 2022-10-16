package dataobject

type UserProfile struct {
	Uid uint64 `json:"uid"`
	Password string `json:"password"`
	Name string `json:"name"`
}