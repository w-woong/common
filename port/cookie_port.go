package port

import "net/http"

type Cookie interface {
	Get(r *http.Request) string
	Set(w http.ResponseWriter, value string)
}
