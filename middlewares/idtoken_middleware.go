package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/port"
)

func GetIDToken(r *http.Request, key string) (string, error) {
	idToken := ""
	cookie, err := r.Cookie(key)
	if err == nil {
		idToken = cookie.Value
	}

	if idToken == "" {
		idToken = r.Header.Get(key)
		if strings.HasPrefix(idToken, "Bearer") {
			authVals := strings.Split(idToken, " ")
			if len(authVals) != 2 {
				return "", common.ErrIDTokenNotFound
			}
			idToken = authVals[1]
		}
	}

	if idToken == "" {
		return "", common.ErrIDTokenNotFound
	}

	return idToken, nil
}

func GetTokenSource(r *http.Request, key string) (string, error) {
	tokenSource := ""
	cookie, err := r.Cookie(key)
	if err == nil {
		tokenSource = cookie.Value
	}
	if tokenSource == "" {
		tokenSource = r.Header.Get(key)
	}

	if tokenSource == "" {
		return "", common.ErrTokenSourceNotFound
	}

	return tokenSource, nil
}
func GetTokenIdentifier(r *http.Request, key string) (string, error) {
	val := ""
	cookie, err := r.Cookie(key)
	if err == nil {
		val = cookie.Value
	}
	if val == "" {
		val = r.Header.Get(key)
	}

	if val == "" {
		return "", common.ErrTokenIdentifierNotFound
	}

	return val, nil
}

func AuthIDTokenHandler(next http.HandlerFunc, validator port.IDTokenValidators,
	tokenIdentifierKey string, idTokenKey string, tokenSourcKey string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idToken, err := GetIDToken(r, idTokenKey)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		tokenSource, err := GetTokenSource(r, tokenSourcKey)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		ctx := r.Context()
		if v, ok := validator[tokenSource]; ok {
			_, claims, err := v.Validate(idToken)
			if err != nil {
				if errors.Is(err, common.ErrTokenExpired) {
					common.HttpErrorWithBody(w, http.StatusUnauthorized,
						common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
					logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
					return
				}
				common.HttpError(w, http.StatusUnauthorized)
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
			ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
			ctx = context.WithValue(ctx, dto.TokenSourceKey{}, tokenSource)
		} else {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AuthIDTokenUserAccountHandler(next http.HandlerFunc, validator port.IDTokenValidators,
	userSvc port.UserSvc,
	tokenIdentifierKey string, idTokenKey string, tokenSourcKey string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idToken, err := GetIDToken(r, idTokenKey)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		tokenSource, err := GetTokenSource(r, tokenSourcKey)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		tokenIdentifier, err := GetTokenIdentifier(r, tokenIdentifierKey)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		ctx := r.Context()
		if v, ok := validator[tokenSource]; ok {
			_, claims, err := v.Validate(idToken)
			if err != nil {
				if errors.Is(err, common.ErrTokenExpired) {
					common.HttpErrorWithBody(w, http.StatusUnauthorized,
						common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
					logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
					return
				}
				common.HttpError(w, http.StatusUnauthorized)
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
			userAccount, err := userSvc.FindByLoginID(ctx, tokenSource, tokenIdentifier, idToken)
			if err != nil {
				common.HttpError(w, http.StatusUnauthorized)
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
			ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
			ctx = context.WithValue(ctx, dto.TokenSourceKey{}, tokenSource)
			ctx = context.WithValue(ctx, dto.UserAccountKey{}, userAccount)
		} else {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
