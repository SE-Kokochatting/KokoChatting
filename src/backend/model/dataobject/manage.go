package dataobject
// FriendRelation 代表着好友关系表
type FriendRelation struct {
	Id uint64 `json:"id"`
	User1 uint64 `json:"user1"`
	User2 uint64 `json:"user2"`
}
// BlockRelation 代表着拉黑的关系表
type BlockRelation struct {
	Id uint64 `json:"id"`
	User uint64 `json:"user"`
	Blocker uint64 `json:"blocker"`
}
// Preprocess 是所有插入、或查询好友关系表时都需要使用的，其功能是使好友列表中user1的id始终小于user2的id
func (friendRelation *FriendRelation)Preprocess (){
	u1 := friendRelation.User1
	u2 := friendRelation.User2
	if friendRelation.User1 > friendRelation.User2 {
		u1 = friendRelation.User2
		u2 = friendRelation.User1
	}
	friendRelation.User1 = u1
	friendRelation.User2 = u2
}