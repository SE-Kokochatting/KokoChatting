package global

import "errors"

type Error struct {
	Status int
	Err    error
}

func (err Error) Error() string {
	return err.Err.Error()
}

// 错误码按照模块的不同第一位不同，user是1，manage是2等等，然后模块内错误码自增
var (
	ConfigPathError    = NewError(errors.New("config path do not exist in config file"), 1000)
	RegisterError      = NewError(errors.New("user register err"), 1001)
	LoginError         = NewError(errors.New("user login err"), 1002)
	PasswordError      = NewError(errors.New("password is err"), 1003)
	GetInfoError       = NewError(errors.New("get userinfo err"), 1004)
	AvatarError        = NewError(errors.New("update avatar err"), 1005)
	JwtParseError      = NewError(errors.New("jwt parse err, please check your token"), 1006)
	JwtExpiredError    = NewError(errors.New("jwt expired err"), 1007)
	IncorrectToken     = NewError(errors.New("token incorrect err"), 1008)
	BindError          = NewError(errors.New("bind error, please check the request body"), 1009)
	DatabaseQueryError = NewError(errors.New("database query error"), 1010)

	DeleteFriendError   = NewError(errors.New("friend delete err"), 2001)
	BlockFriendError    = NewError(errors.New("friend block err"), 2002)
	CreatGroupError     = NewError(errors.New("create group err"), 2003)
	QuitGroupError      = NewError(errors.New("quit group err"), 2004)
	GetFriendListError  = NewError(errors.New("get friend list err"), 2005)
	GetFriendInfoError  = NewError(errors.New("get friend info err"), 2006)
	SetGroupAvatarError = NewError(errors.New("set group avatar error"), 2007)
	PermissionError     = NewError(errors.New("permission error"), 2008)
	GetGroupError       = NewError(errors.New("get group list err"), 2009)
	GetGroupInfoError   = NewError(errors.New("get group info err"), 2010)
	TransferHostError   = NewError(errors.New("transfer host err"), 2011)
	TransferAdminError  = NewError(errors.New("transfer administrator err"), 2012)
	TransferMemError    = NewError(errors.New("transfer member err"), 2013)

	RequestFormatError = NewError(errors.New("request body request error"), 3000)

	MessageServerBusy        = NewError(errors.New("message server busy,please try again later"), 4000)
	MessageInternalError     = NewError(errors.New("message server internal unknown error"), 4001)
	StoreMessageError        = NewError(errors.New("internal error : store msg error"), 4002)
	MessageTypeError         = NewError(errors.New("message type error"), 4003)
	UpgradeProtocolError     = NewError(errors.New("upgrade protocol error,check your http header"), 4004)
	WsJsonMarshalError       = NewError(errors.New("internal error: json marshal error"), 4005)
	QueryBlockRelationError  = NewError(errors.New("internal error: query block relation error"), 4006)
	HasBeenBlocked           = NewError(errors.New("the user has been blocked"), 4007)
	QueryIsInGroup           = NewError(errors.New("query user is in group error"), 4008)
	MessageSenderError       = NewError(errors.New("current user is not in receiver group"), 4009)
	RevertMessageError       = NewError(errors.New("msg whose id equals to 'msgid' is not sent by current user"), 4010)
	RevertedMessageTypeError = NewError(errors.New("only group msg and single msg can be reverted"), 4011)
	MsgHasBeenRevertedError  = NewError(errors.New("the msg has been reverted"), 4012)
	MessagePullBindError     = NewError(errors.New("pull message bind error"), 4013)
	PullOutlineError         = NewError(errors.New("pull outline error"), 4014)
	PullMessageError         = NewError(errors.New("pull message error"), 4015)
)

func NewError(err error, status int) error {
	return Error{
		Status: status,
		Err:    err,
	}
}
