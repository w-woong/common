package common

import "errors"

var ErrRecordNotFound = errors.New("record not found")
var ErrLoginIDAlreadyExists = errors.New("login id already exists")
var ErrCreateUser = errors.New("failed to create user")

var (
	ErrIDTokenNotFound     = errors.New("id token not found")
	ErrIDTokenInconsistent = errors.New("id token is not consistent")

	ErrInvalidKey      = errors.New("key is invalid")
	ErrInvalidKeyType  = errors.New("key is of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")

	ErrTokenMalformed        = errors.New("token is malformed")
	ErrTokenUnverifiable     = errors.New("token is unverifiable")
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")

	ErrTokenInvalidAudience  = errors.New("token has invalid audience")
	ErrTokenExpired          = errors.New("token is expired")
	ErrTokenUsedBeforeIssued = errors.New("token used before issued")
	ErrTokenInvalidIssuer    = errors.New("token has invalid issuer")
	ErrTokenNotValidYet      = errors.New("token is not valid yet")
	ErrTokenInvalidId        = errors.New("token has invalid id")
	ErrTokenInvalidClaims    = errors.New("token has invalid claims")

	ErrUnexpectedTokenClaims = errors.New("token has unexpected claims")
)

var (
	StatusTryRefreshIDToken = 1000
)
