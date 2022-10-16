package internal

import (
	"errors"
)

var (
	AddConnError     = errors.New("conn nums is full")
	UserOfflineError = errors.New("user have been offline")
)
