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

type OAuth2Error struct {
	ErrorTitle       string `json:"error"`
	ErrorDebug       string `json:"error_debug,omitempty"`
	ErrorDescription string `json:"error_description"`
	ErrorHint        string `json:"error_hint,omitempty"`
	StatusCode       int    `json:"status_code"`

	TryRefresh        *bool `json:"try_refresh,omitempty"`
	TryReauthenticate *bool `json:"try_reauthenticate,omitempty"`
}

func OAuth2ErrorInvalidRequest(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("invalid_request", errorDescription, "", statusCode)
}

func OAuth2ErrorInvalidClient(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("invalid_client", errorDescription, "", statusCode)
}
func OAuth2ErrorInvalidGrant(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("invalid_grant", errorDescription, "", statusCode)
}
func OAuth2ErrorInvalidScope(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("invalid_scope", errorDescription, "", statusCode)
}
func OAuth2ErrorUnauthorizedClient(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("unauthorized_client", errorDescription, "", statusCode)
}
func OAuth2ErrorUnsupportedGrantType(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("unsupported_grant_type", errorDescription, "", statusCode)
}
func OAuth2ErrorTryRefresh(errorDescription string) *OAuth2Error {
	val := true
	return &OAuth2Error{
		ErrorTitle:       "invalid_request",
		ErrorDescription: errorDescription,
		// ErrorHint:        errorHint,
		StatusCode: http.StatusUnauthorized,
		TryRefresh: &val,
	}
}
func OAuth2ErrorTryReauthenticate(errorDescription string) *OAuth2Error {
	val := true
	return &OAuth2Error{
		ErrorTitle:       "invalid_request",
		ErrorDescription: errorDescription,
		// ErrorHint:        errorHint,
		StatusCode:        http.StatusUnauthorized,
		TryReauthenticate: &val,
	}
}

// OAuth2ErrorRequestDenied when user denied request
func OAuth2ErrorRequestDenied(errorDescription string, statusCode int) *OAuth2Error {
	return NewOAuth2Error("request_denied", errorDescription, "", statusCode)
}

func NewOAuth2Error(errorTitle string, errorDescription string, errorHint string, statusCode int) *OAuth2Error {
	return &OAuth2Error{
		ErrorTitle:       errorTitle,
		ErrorDescription: errorDescription,
		ErrorHint:        errorHint,
		StatusCode:       statusCode,
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
