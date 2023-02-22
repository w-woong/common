package adapter

import (
	"net/http"
	"time"
)

type TokenCookie struct {
	expireAfter   time.Duration
	idTokenCookie string
	mode          http.SameSite
	domain        string
}

func NewTokenCookie(expireAfter time.Duration, idTokenCookie string) *TokenCookie {
	return &TokenCookie{
		expireAfter:   expireAfter,
		idTokenCookie: idTokenCookie,
		mode:          http.SameSiteLaxMode,
		domain:        "",
	}
}

func NewCookie(expireAfter time.Duration, mode http.SameSite,
	idTokenCookie string, domain string) *TokenCookie {

	return &TokenCookie{
		expireAfter:   expireAfter,
		idTokenCookie: idTokenCookie,
		mode:          mode,
		domain:        domain,
	}
}

func set(w http.ResponseWriter, sameSiteMode http.SameSite, name, value string, expireAfter time.Duration, maxAge int, domain string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		SameSite: sameSiteMode,
		Path:     "/",
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

func (a *TokenCookie) GetIDToken(r *http.Request) string {
	return get(r, a.idTokenCookie)
}

func (a *TokenCookie) SetIDToken(w http.ResponseWriter, idToken string) {
	set(w, a.mode, a.idTokenCookie, idToken, a.expireAfter, 0, a.domain)
}
