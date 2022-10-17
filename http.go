package common

import (
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

var HttpBodyOK = HttpBody{Status: http.StatusOK}

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

func (m *HttpBody) EncodeTo(w io.Writer) error {
	return si.EncodeJson(w, m)
}
