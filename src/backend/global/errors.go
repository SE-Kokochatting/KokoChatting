package global

import "errors"

var (
	ConfigPathError = errors.New("config path do not exist in config file")
)