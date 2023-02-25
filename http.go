package common

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-wonk/si"
)

func HttpOK(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	return si.EncodeJson(w, &HttpBodyOK)
}

func HttpError(w http.ResponseWriter, status int) {
	HttpErrorWithMessage(w, http.StatusText(status), status)
}

func HttpErrorWithMessage(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	si.EncodeJson(w, NewHttpBody(message, status))
}
func HttpErrorWithBody(w http.ResponseWriter, status int, body *HttpBody) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	si.EncodeJson(w, body)
}

var HttpBodyOK = HttpBody{Status: http.StatusOK}
var HttpBodyRecordNotFound = HttpBody{Count: 0, Status: http.StatusOK, Message: ErrRecordNotFound.Error()}

type HttpBody struct {
	Status    int         `json:"status,omitempty"`
	Message   string      `json:"message,omitempty"`
	Count     int         `json:"count,omitempty"`
	Document  interface{} `json:"document,omitempty"`
	Documents interface{} `json:"documents,omitempty"`
}

func NewHttpBody(message string, status int) *HttpBody {
	return &HttpBody{
		Status:  status,
		Message: message,
	}
}

func (m *HttpBody) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *HttpBody) EncodeTo(w io.Writer) error {
	return si.EncodeJson(w, m)
}

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
)

type OAuth2Error struct {
	ErrorType        OAuth2ErrorType `json:"error,omitempty"`
	ErrorDebug       string          `json:"error_debug,omitempty"`
	ErrorDescription string          `json:"error_description,omitempty"`
	ErrorHint        string          `json:"error_hint,omitempty"`
	StatusCode       *int            `json:"status_code,omitempty"`

	TryRefresh        *bool `json:"try_refresh,omitempty"`
	TryReauthenticate *bool `json:"try_reauthenticate,omitempty"`
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

func OAuth2ErrorInvalidRequest(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(InvalidRequestErrorType, errorDescription, statusCode)
}

func OAuth2ErrorInvalidClient(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(InvalidClientErrorType, errorDescription, statusCode)
}
func OAuth2ErrorInvalidGrant(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(InvalidGrantErrorType, errorDescription, statusCode)
}
func OAuth2ErrorInvalidScope(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(InvalidScopeErrorType, errorDescription, statusCode)
}
func OAuth2ErrorUnauthorizedClient(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(UnauthorizedClientErrorType, errorDescription, statusCode)
}
func OAuth2ErrorUnsupportedGrantType(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error(UnsupportedGrantTypeErrorType, errorDescription, statusCode)
}
func OAuth2ErrorTryRefresh(errorDescription string) *OAuth2Error {
	val := true
	sc := http.StatusUnauthorized
	return &OAuth2Error{
		ErrorType:        InvalidRequestErrorType,
		ErrorDescription: errorDescription,
		// ErrorHint:        errorHint,
		StatusCode: &sc,
		TryRefresh: &val,
	}
}
func OAuth2ErrorTryReauthenticate(errorDescription string) *OAuth2Error {
	val := true
	sc := http.StatusUnauthorized
	return &OAuth2Error{
		ErrorType:        InvalidRequestErrorType,
		ErrorDescription: errorDescription,
		// ErrorHint:        errorHint,
		StatusCode:        &sc,
		TryReauthenticate: &val,
	}
}

func OAuth2ErrorInvalidClaims() *OAuth2Error {
	return NewOAuth2Error(InvalidRequestErrorType, ErrTokenInvalidClaims.Error(), http.StatusUnauthorized)
}

// OAuth2ErrorRequestDenied when user denied request
func OAuth2ErrorRequestDenied(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("request_denied", errorDescription, statusCode)
}

func NewOAuth2Error(errorType OAuth2ErrorType, errorDescription string, statusCode int) *OAuth2Error {
	return &OAuth2Error{
		ErrorType:        errorType,
		ErrorDescription: errorDescription,
		StatusCode:       &statusCode,
	}
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
