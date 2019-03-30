package metro

import "errors"

var (
	errInvTkn  = errors.New("invalid token")
	errExists  = errors.New("resource already exists")
	errNExists = errors.New("resource does NOT exist")
)
