package common

import "errors"

var ErrRecordNotFound = errors.New("record not found")
var ErrLoginIDAlreadyExists = errors.New("login id already exists")
var ErrCreateUser = errors.New("failed to create user")
