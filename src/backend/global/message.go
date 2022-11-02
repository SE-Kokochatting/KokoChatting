package global

const(
	SingleMessage = 0 + iota
	GroupMessage
	FriendRequestNotify
	RevertSingleMessageNotify
	RevertGroupMessageNotify
	HasReadSingleNotify
	HasReadGroupNotify
	JoinGroupRequestNotify
	QuitGroupNotify
	JoinGroupNotify
)

