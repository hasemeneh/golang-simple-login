package constants

import "errors"

var ErrUsernameAlreadyUsed = errors.New("Username unavailable")
var ErrEmailAlreadyRegistered = errors.New("Email is registered")
var ErrWrongPassword = errors.New("Wrong Password")
var ErrInvalidToken = errors.New("Invalid Token")
