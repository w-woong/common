package utils

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

// AuthBearer retrieves value from Authorization header
func AuthBearer(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	authVal := strings.Split(authHeader, " ")
	if len(authVal) != 2 {
		return ""
	}

	if authVal[0] != "Bearer" {
		return ""
	}

	return authVal[1]
}

func SetNoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate;")
	w.Header().Set("pragma", "no-cache")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func RequestDump(r *http.Request) ([]byte, error) {
	return httputil.DumpRequest(r, true)
}
