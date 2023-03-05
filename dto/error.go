package dto

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-wonk/si/v2"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrLoginIDAlreadyExists = errors.New("login id already exists")
var ErrCreateUser = errors.New("failed to create user")

var (
	ErrTokenSourceNotFound     = errors.New("token source not found")
	ErrTokenIdentifierNotFound = errors.New("token identifier not found")
	ErrIDTokenNotFound         = errors.New("id token not found")
	ErrIDTokenInconsistent     = errors.New("id token is not consistent")

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
	StatusTryRefreshIDToken  = 1000
	StatusInvalidTokenClaims = 1001
)

type OAuth2ErrorType string

const (
	InvalidRequestErrorType       OAuth2ErrorType = "invalid_request"
	InvalidClientErrorType        OAuth2ErrorType = "invalid_client"
	InvalidGrantErrorType         OAuth2ErrorType = "invalid_grant"
	InvalidScopeErrorType         OAuth2ErrorType = "invalid_scope"
	UnauthorizedClientErrorType   OAuth2ErrorType = "unauthorized_client"
	UnsupportedGrantTypeErrorType OAuth2ErrorType = "unsupported_grant_type"

	AccessDeniedErrorType    OAuth2ErrorType = "access_denied"
	UnsupportedResponseType  OAuth2ErrorType = "unsupported_response_type"
	ServerErrorType          OAuth2ErrorType = "server_error"
	TempUnavailableErrorType OAuth2ErrorType = "temporarily_unavailable"
	RequestDeniedErrorType   OAuth2ErrorType = "request_denied"
)

func OAuth2ErrorInvalidRequest(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(InvalidRequestErrorType, errorDescription, statusCode)
}

func OAuth2ErrorInvalidClient(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(InvalidClientErrorType, errorDescription, http.StatusUnauthorized)
}
func OAuth2ErrorInvalidGrant(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(InvalidGrantErrorType, errorDescription, http.StatusUnauthorized)
}
func OAuth2ErrorInvalidScope(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(InvalidScopeErrorType, errorDescription, http.StatusUnauthorized)
}
func OAuth2ErrorUnauthorizedClient(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(UnauthorizedClientErrorType, errorDescription, http.StatusUnauthorized)
}
func OAuth2ErrorUnsupportedGrantType(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(UnsupportedGrantTypeErrorType, errorDescription, http.StatusBadRequest)
}
func OAuth2ErrorAccessDenied(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(AccessDeniedErrorType, errorDescription, http.StatusUnauthorized)
}
func OAuth2ErrorUnsupportedResponseType(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(UnsupportedResponseType, errorDescription, http.StatusBadRequest)
}
func OAuth2ErrorServerError(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(ServerErrorType, errorDescription, http.StatusInternalServerError)
}
func OAuth2ErrorTempUnavailable(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(TempUnavailableErrorType, errorDescription, http.StatusInternalServerError)
}
func OAuth2ErrorTryRefresh(errorDescription string) *OAuth2Error {
	oerr := NewOAuth2Error(InvalidRequestErrorType, errorDescription, http.StatusUnauthorized)
	oerr.SetTryRefresh(true)
	return oerr
}
func OAuth2ErrorTryReauthenticate(errorDescription string) *OAuth2Error {
	oerr := NewOAuth2Error(InvalidRequestErrorType, errorDescription, http.StatusUnauthorized)
	oerr.SetTryReauthenticate(true)
	return oerr
}

func OAuth2ErrorInvalidClaims() *OAuth2Error {
	return NewOAuth2Error(InvalidRequestErrorType, ErrTokenInvalidClaims.Error(), http.StatusUnauthorized)
}

// OAuth2ErrorRequestDenied
func OAuth2ErrorRequestDenied(errorDescription string) *OAuth2Error {
	return NewOAuth2Error(RequestDeniedErrorType, errorDescription, http.StatusUnauthorized)
}

type OAuth2Error struct {
	ErrorType        OAuth2ErrorType `json:"error,omitempty"`
	ErrorDebug       string          `json:"error_debug,omitempty"`
	ErrorDescription string          `json:"error_description,omitempty"`
	ErrorHint        string          `json:"error_hint,omitempty"`
	StatusCode       *int            `json:"status_code,omitempty"`

	TryRefresh        *bool `json:"try_refresh,omitempty"`
	TryReauthenticate *bool `json:"try_reauthenticate,omitempty"`
}

func NewOAuth2Error(errorType OAuth2ErrorType, errorDescription string, statusCode int) *OAuth2Error {
	return &OAuth2Error{
		ErrorType:        errorType,
		ErrorDescription: errorDescription,
		StatusCode:       &statusCode,
	}
}

func (o OAuth2Error) GetStatusCode() int {
	if o.StatusCode == nil {
		return 0
	}
	return *o.StatusCode
}
func (o OAuth2Error) GetTryRefresh() bool {
	if o.TryRefresh == nil {
		return false
	}
	return *o.TryRefresh
}
func (o OAuth2Error) GetTryReauthenticate() bool {
	if o.TryReauthenticate == nil {
		return false
	}
	return *o.TryReauthenticate
}
func (o *OAuth2Error) SetTryRefresh(val bool) {
	o.TryRefresh = &val
}
func (o *OAuth2Error) SetTryReauthenticate(val bool) {
	o.TryReauthenticate = &val
}

func (m *OAuth2Error) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *OAuth2Error) EncodeTo(w io.Writer) error {
	return si.EncodeJson(w, m)
}

func (m *OAuth2Error) Error() string {
	return m.String()
}
