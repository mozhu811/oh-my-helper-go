package errors

import "errors"

var (
	ErrorCookieExpired = errors.New("cookie已失效")
)
