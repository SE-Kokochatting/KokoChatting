package global

import "errors"

type Error struct {
	Status int
	Err error
}

func (err Error) Error() string {
	return err.Err.Error()
}
// 错误码按照模块的不同第一位不同，user是1，manage是2等等，然后模块内错误码自增
var (
	ConfigPathError = NewError(errors.New("config path do not exist in config file"), 1000)

	RegisterError = NewError(errors.New("user register err"), 1001)
	PasswordError = NewError(errors.New("password is err"), 1002)
	GetInfoError = NewError(errors.New("get userinfo err"), 1003)
  
	DeleteFriendError = NewError(errors.New("friend delete err"),2001)
	BlockFriendError = NewError(errors.New("friend block err"),2002)
	CreatGroupError = NewError(errors.New("creat group err"),2003)
	QuitGroupError = NewError(errors.New("quit group err"),2004)
  
  	MessageServerBusy = NewError(errors.New("message server busy,please try again later"),4000)
	MessageInternalError = NewError(errors.New("message server internal unknown error"),4001)

)

func NewError(err error, status int) error {
	return Error{
		Status: status,
		Err: err,
	}
}