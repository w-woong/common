package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-wonk/si"
	"github.com/go-wonk/si/sicore"
	"github.com/go-wonk/si/sihttp"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/validators"
)

// AuthHMACHandler to verify the request
func AuthHMACHandler(next http.HandlerFunc, hmacHeader string, hmacKey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBytes []byte
		var err error

		if r.Body != nil {
			reqBytes, err = si.ReadAll(r.Body)
			if err != nil {
				HttpError(w, http.StatusBadRequest)
				logger.Error(http.StatusText(http.StatusBadRequest), logger.UrlField(r.URL.String()))
				return
			}
		}

		var msg []byte
		if strings.Compare(r.Method, http.MethodGet) == 0 || strings.Compare(r.Method, http.MethodHead) == 0 {
			msg = []byte(r.URL.Path)
		} else {
			msg = reqBytes
		}

		hmacHexStr, err := sicore.HmacSha256HexEncoded(string(hmacKey), msg)
		if err != nil {
			HttpError(w, http.StatusInternalServerError)
			logger.Error(http.StatusText(http.StatusInternalServerError), logger.UrlField(r.URL.String()))
			return
		}

		if hmacHexStr != r.Header.Get(hmacHeader) {
			HttpError(w, http.StatusUnauthorized)
			logger.Warn(fmt.Sprintf("generated hmac value %v is invalid(expected: %v)", r.Header.Get(hmacHeader), hmacHexStr),
				logger.UrlField(r.URL.String()), logger.ReqBodyField(reqBytes))
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		next.ServeHTTP(w, r)
	}
}

func AuthBearerHandler(next http.HandlerFunc, bearerKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		authVal := strings.Split(authHeader, " ")
		if len(authVal) != 2 {
			HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		if authVal[1] != bearerKey {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		next.ServeHTTP(w, r)
	}
}

// JWTAuthMultipartHandler to verify the request
func JWTAuthMultipartHandler(next http.HandlerFunc, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := handleJWTAuth(w, r, secret)
		if err != nil {
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// AuthJWTHandler to verify the request
func AuthJWTHandler(next http.HandlerFunc, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := handleJWTAuth(w, r, secret)
		if err != nil {
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func handleJWTAuth(w http.ResponseWriter, r *http.Request, secretKey string) error {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		HttpError(w, http.StatusUnauthorized)
		logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
		return errors.New("not authorized")
	}
	authHeaderSplit := strings.Split(authHeader, " ")
	if len(authHeaderSplit) != 2 {
		HttpError(w, http.StatusUnauthorized)
		logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
		return errors.New("not authorized")
	}
	token := authHeaderSplit[1]
	claim, refresh, err := ValidateAccessToken(token, secretKey)

	if err != nil {
		if refresh {
			refHeader := r.Header.Get("Authorization-Refresh")
			if refHeader == "" {
				// JWT: 액세스토큰 만료되어, Refresh 토큰을 이용하여 액세스토큰을 재발급 받도록 응답 보냄
				resMsg := fmt.Sprintf(`{"status_msg":"Not Authorized", "status": %v, "refresh": true, "refresh_token_header": "%v"}`, http.StatusUnauthorized, "Authorization-Refresh")
				http.Error(w, resMsg, http.StatusUnauthorized)
				logger.Warn(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return errors.New("not authorized")
			}
			rfSub, err := ValidateRefreshToken(refHeader, secretKey, claim.UsrID, claim.Sub)
			if err != nil {
				http.Error(w, "Not Authorized", http.StatusUnauthorized)
				logger.Warn(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return errors.New("not authorized")
			}
			newAccessToken, err := GenerateAccessToken(secretKey, claim.UsrID, claim.AccLvl, rfSub, time.Minute*10)
			if err != nil {
				http.Error(w, "Not Authorized", http.StatusUnauthorized)
				logger.Warn(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return errors.New("not authorized")
			}
			// w.Header().Set("Retry-With-Refresh", "1")

			retry := fmt.Sprintf(`{"status_msg":"Not Authorized", "status": %v, "refreshed": true, "new_access_token": "%v"}`, http.StatusUnauthorized, newAccessToken)
			http.Error(w, retry, http.StatusUnauthorized)
			logger.Warn(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return errors.New("not authorized")
		}
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
		return errors.New("not authorized")
	}

	if claim != nil {
		w.Header().Set("usrID", claim.UsrID)
	}
	return nil
}
func WithHeaderHmac256(key string, secret []byte) sihttp.RequestOptionFunc {
	return sihttp.RequestOptionFunc(func(req *http.Request) error {
		header := req.Header
		if _, ok := header[key]; ok {
			// skip
			return nil
		}

		contentType := header.Get("Content-Type")
		if strings.Contains(contentType, "multipart/form-data") {
			return nil
		}

		if req.Method == http.MethodGet || req.Method == http.MethodHead {
			hashed, err := sicore.HmacSha256HexEncoded(string(secret), []byte(req.URL.Path))
			if err != nil {
				return err
			}
			header[key] = []string{hashed}
			return nil
		}

		if req.GetBody == nil {
			return nil
		}

		r, err := req.GetBody()
		if err != nil {
			return err
		}

		hashed, err := sicore.HmacSha256HexEncodedWithReader(string(secret), r)
		if err != nil {
			return err
		}
		header[key] = []string{hashed}

		return nil
	})
}

func AuthIDTokenHandler(next http.HandlerFunc, validator validators.IDTokenValidators,
	cookieName, headerName string, tokenSourcCookieName string, tokenSourceHeaderName string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		idToken := cookie.Value
		if idToken == "" {
			idToken = r.Header.Get(headerName)
			if strings.HasPrefix(idToken, "Bearer") {
				authVals := strings.Split(idToken, " ")
				if len(authVals) != 2 {
					HttpError(w, http.StatusUnauthorized)
					logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
					return
				}
				idToken = authVals[1]
			}
		}

		cookie, err = r.Cookie(tokenSourcCookieName)
		if err != nil {
			HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}
		tokenSource := cookie.Value
		if tokenSource == "" {
			tokenSource = r.Header.Get(tokenSourceHeaderName)
		}

		if v, ok := validator[tokenSource]; ok {
			_, _, err = v.Validate(idToken)
			if err != nil {
				HttpError(w, http.StatusUnauthorized)
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}
