package adapter

import (
	"net/http"
	"time"
)

type SecureCookie struct {
	expireAfter   time.Duration
	idTokenCookie string
	mode          http.SameSite
	domain        string
	path          string
}

func NewSecureCookie(expireAfter time.Duration, mode http.SameSite,
	idTokenCookie string, domain string, path string) *SecureCookie {

	return &SecureCookie{
		expireAfter:   expireAfter,
		idTokenCookie: idTokenCookie,
		mode:          mode,
		domain:        domain,
		path:          path,
	}
}

func (a *SecureCookie) Get(r *http.Request) string {
	return get(r, a.idTokenCookie)
}

func (a *SecureCookie) Set(w http.ResponseWriter, idToken string) {
	set(w, a.mode, a.idTokenCookie, idToken, a.expireAfter, 0, a.domain, a.path)
}

func set(w http.ResponseWriter, sameSiteMode http.SameSite, name, value string, expireAfter time.Duration, maxAge int, domain string, path string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		SameSite: sameSiteMode,
		Path:     path,
		Expires:  time.Now().Add(expireAfter),
		MaxAge:   maxAge,
		Secure:   true,
		Domain:   domain,
	}
	http.SetCookie(w, &cookie)
}

func get(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}
