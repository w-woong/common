package adapter

import (
	"net/http"
	"time"
)

type TokenCookie struct {
	expireAfter   time.Duration
	idTokenCookie string
}

func NewTokenCookie(expireAfter time.Duration, idTokenCookie string) *TokenCookie {

	return &TokenCookie{
		expireAfter:   expireAfter,
		idTokenCookie: idTokenCookie,
	}
}

func set(w http.ResponseWriter, sameSiteMode http.SameSite, name, value string, expireAfter time.Duration, maxAge int) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		SameSite: sameSiteMode,
		Path:     "/",
		Expires:  time.Now().Add(expireAfter),
		MaxAge:   maxAge,
		Secure:   true,
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
	set(w, http.SameSiteLaxMode, a.idTokenCookie, idToken, a.expireAfter, 0)
}
