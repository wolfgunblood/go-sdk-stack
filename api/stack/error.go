package stack

import (
	"errors"
)

var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidToken = errors.New("invalid token")
)
