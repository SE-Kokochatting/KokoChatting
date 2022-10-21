package dataobject

type FriendRelation struct {
	Id uint64 `json:"id"`
	User1 uint64 `json:"user1"`
	User2 uint64 `json:"user2"`
}

func (friendRelation *FriendRelation)Preprocess (){
	var u1,u2 uint64
	if friendRelation.User1 < friendRelation.User2 {
		u1 = friendRelation.User2
		u2 = friendRelation.User1
	}
	friendRelation.User1 = u1
	friendRelation.User2 = u2
}