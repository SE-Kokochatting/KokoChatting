package global

import "errors"

type Error struct {
	Status int
	Err error
}

func (err Error) Error() string {
	return err.Err.Error()
}

var (
	ConfigPathError = NewError(errors.New("config path do not exist in config file"), 1000)
	RegisterError = NewError(errors.New("user register err"), 1001)
	DeleteFriendError = NewError(errors.New("friend delete err"),2001)
	PasswordError = NewError(errors.New("password is err"), 3001)
)

func NewError(err error, status int) error {
	return Error{
		Status: status,
		Err: err,
	}
}