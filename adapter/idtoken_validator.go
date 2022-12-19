package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/allegro/bigcache/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
)

type jwksIDTokenValidator struct {
	cache              *bigcache.BigCache
	jwksUrl            string
	tokenSourceKey     string
	tokenIdentifierKey string
	idTokenKey         string
}

func NewJwksIDTokenValidator(jwksUrl string, tokenSourceKey, tokenIdentifierKey, idTokenKey string) (*jwksIDTokenValidator, error) {

	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: false,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 1024,
	}
	// cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &jwksIDTokenValidator{
		cache:              cache,
		jwksUrl:            jwksUrl,
		tokenSourceKey:     tokenSourceKey,
		tokenIdentifierKey: tokenIdentifierKey,
		idTokenKey:         idTokenKey,
	}, nil
}

func (u *jwksIDTokenValidator) TokenSourceKey() string {
	return u.tokenSourceKey
}

func (u *jwksIDTokenValidator) TokenIdentifierKey() string {
	return u.tokenIdentifierKey
}

func (u *jwksIDTokenValidator) IDTokenKey() string {
	return u.idTokenKey
}

func (u *jwksIDTokenValidator) Validate(idToken string) (*jwt.Token, *dto.IDTokenClaims, error) {
	if idToken == "" {
		return nil, nil, common.ErrIDTokenNotFound
	}

	jwksJson, err := u.cache.Get("jwks")
	if err != nil {
		jwksJson, err = GetJwks(u.jwksUrl)
		if err != nil {
			return nil, nil, err
		}
		if err = u.cache.Set("jwks", jwksJson); err != nil {
			return nil, nil, err
		}
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.NewJSON(jwksJson)
	if err != nil {
		return nil, nil, err
	}

	// Parse the JWT.
	// token, err := jwt.Parse(idToken, jwks.Keyfunc)
	token, err := jwt.ParseWithClaims(idToken, &dto.IDTokenClaims{}, jwks.Keyfunc)
	if token.Valid {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, err
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, nil, common.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, nil, common.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		// return nil, auth.ErrTokenNotValidYet
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	}

	return nil, nil, err
}

func GetJwks(url string) (json.RawMessage, error) {
	// TODO: cache jwks

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jwksJSON json.RawMessage = b

	return jwksJSON, nil
}

func GetJwksUrl(openIDConfUrl string) (string, error) {

	res, err := http.Get(openIDConfUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	if err = json.Unmarshal(resb, &m); err != nil {
		return "", err
	}
	jwksUrl, ok := m["jwks_uri"]
	if !ok {
		return "", errors.New("not found")
	}
	return jwksUrl.(string), nil
}
