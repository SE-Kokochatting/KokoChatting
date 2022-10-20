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
	RegisterError = NewError(errors.New("user regiter err"), 1001)
)

func NewError(err error, status int) error {
	return Error{
		Status: status,
		Err: err,
	}
}