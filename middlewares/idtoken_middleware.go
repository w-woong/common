package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/validators"
)

func AuthIDTokenHandler(next http.HandlerFunc, validator validators.IDTokenValidators,
	cookieName, headerName string,
	tokenSourcCookieName string, tokenSourceHeaderName string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idToken := ""
		cookie, err := r.Cookie(cookieName)
		if err == nil {
			idToken = cookie.Value
		}

		if idToken == "" {
			idToken = r.Header.Get(headerName)
			if strings.HasPrefix(idToken, "Bearer") {
				authVals := strings.Split(idToken, " ")
				if len(authVals) != 2 {
					common.HttpError(w, http.StatusUnauthorized)
					logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
					return
				}
				idToken = authVals[1]
			}
		}

		tokenSource := ""
		cookie, err = r.Cookie(tokenSourcCookieName)
		if err == nil {
			tokenSource = cookie.Value
		}
		if tokenSource == "" {
			tokenSource = r.Header.Get(tokenSourceHeaderName)
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
			ctx = context.WithValue(ctx, validators.IDTokenClaimsKey{}, *claims)
			ctx = context.WithValue(ctx, validators.TokenSourceKey{}, tokenSource)
		} else {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
