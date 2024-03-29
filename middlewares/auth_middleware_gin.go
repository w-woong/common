package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/port"
	"github.com/w-woong/common/utils"
)

func GetIDTokenGin(cookie port.Cookie) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.Get(r)
		}

		c.Set(dto.IDTokenCtxKey, idToken)
		c.Next()
	}
}

// GetIDTokenJwtAndClaims validates and refresh id_token. It first retrieves id_token from request, r.
// Then, it parses the token to get the claims. It ignores ErrTokenExpired.
func GetIDTokenJwtAndClaimsGin(cookie port.Cookie, parser port.IDTokenParser) gin.HandlerFunc {

	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.Get(r)
		}

		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if !errors.Is(err, dto.ErrTokenExpired) {
				oerr := dto.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
				c.Abort()
				return
			}
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := dto.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		c.Set(dto.IDTokenClaimsCtxKey, *claims)
		c.Set(dto.JwtCtxKey, *jwtToken)
		c.Next()
	}
}

func AuthIDTokenGin(cookie port.Cookie, parser port.IDTokenParser) gin.HandlerFunc {

	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.Get(r)
		}
		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if errors.Is(err, dto.ErrTokenExpired) {
				oerr := dto.OAuth2ErrorTryRefresh(err.Error())
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
				c.Abort()

				// common.HttpErrorWithBody(w, http.StatusUnauthorized,
				// 	common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
				// logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}

			oerr := dto.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := dto.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		c.Set(dto.IDTokenClaimsCtxKey, *claims)
		c.Next()
	}
}

func AuthIDTokenIgnoreErrGin(cookie port.Cookie, parser port.IDTokenParser) gin.HandlerFunc {

	return func(c *gin.Context) {
		r := c.Request
		// w := c.Writer

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.Get(r)
		}

		jwtToken, _ := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if jwtToken == nil {
			c.Set(dto.IDTokenClaimsCtxKey, dto.IDTokenClaims{})
			c.Next()
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok || claims == nil {
			c.Set(dto.IDTokenClaimsCtxKey, dto.IDTokenClaims{})
			c.Next()
			return
		}

		c.Set(dto.IDTokenClaimsCtxKey, *claims)
		c.Next()
	}
}

func AuthIDTokenUserAccountSvcGin(cookie port.Cookie, parser port.IDTokenParser,
	userSvc port.UserSvc) gin.HandlerFunc {

	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.Get(r)
		}
		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if errors.Is(err, dto.ErrTokenExpired) {
				oerr := dto.OAuth2ErrorTryRefresh(err.Error())
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
				c.Abort()
				return
			}
			oerr := dto.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := dto.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		userAccount, err := userSvc.FindByIDToken(c, idToken)
		if err != nil {
			oerr := dto.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}
		c.Set(dto.IDTokenClaimsCtxKey, *claims)
		c.Set(dto.UserAccountCtxKey, userAccount)
		c.Next()
	}
}
