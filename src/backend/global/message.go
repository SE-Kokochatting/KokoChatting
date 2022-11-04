package global

const(
	SingleMessage = 0 + iota
	GroupMessage
	FriendRequestNotify
	RevertSingleMessageNotify
	RevertGroupMessageNotify
	JoinGroupRequestNotify
	QuitGroupNotify
	JoinGroupNotify
	AddFriendResponseNotify
	DeleteFriendNotify
)

