package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/allegro/bigcache/v3"
)

type JwksGetter interface {
	Get() ([]byte, error)
}

type JwksCache struct {
	store   *bigcache.BigCache
	jwksUrl string
}

func NewJwksCache(jwksUrl string) (*JwksCache, error) {

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
	return &JwksCache{
		store:   cache,
		jwksUrl: jwksUrl,
	}, nil
}

func (u *JwksCache) Get() ([]byte, error) {
	jwksJson, err := u.store.Get("jwks")
	if err != nil {
		jwksJson, err = getJwks(u.jwksUrl)
		if err != nil {
			return nil, err
		}
		if err = u.store.Set("jwks", jwksJson); err != nil {
			return nil, err
		}
	}
	return jwksJson, err
}

type JwksHttp struct {
	jwksUrl string
}

func NewJwksHttp(jwksUrl string) (*JwksHttp, error) {
	return &JwksHttp{
		jwksUrl: jwksUrl,
	}, nil
}

func (u *JwksHttp) Get() ([]byte, error) {
	return getJwks(u.jwksUrl)
}

func getJwks(url string) (json.RawMessage, error) {
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

	m, err := GetOpenIDConfig(openIDConfUrl)
	if err != nil {
		return "", err
	}
	jwksUrl, ok := m["jwks_uri"]
	if !ok {
		return "", errors.New("not found")
	}
	return jwksUrl.(string), nil
}

func GetOpenIDConfig(openIDConfUrl string) (map[string]interface{}, error) {
	res, err := http.Get(openIDConfUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}
